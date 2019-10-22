package provider

import (
    "fmt"
    "io"
    "io/ioutil"
    "os"
    "path/filepath"

    "github.com/pkg/sftp"
    "golang.org/x/crypto/ssh"
)

type SftpProvider struct {
    SshClientConfig *ssh.ClientConfig
    SshUrl          string
    BasePath        string
}

func NewSftpProvider(basePath string, host string, port string, user string, sshKeyPath string) (*SftpProvider, error) {
    sshUrl := fmt.Sprintf("%s:%s", host, port)
    keyBuf, err := ioutil.ReadFile(sshKeyPath)
    if err != nil {
        return nil, err
    }
    key, err := ssh.ParsePrivateKey(keyBuf)
    sshClientConfig := &ssh.ClientConfig{
        User: user,
        Auth: []ssh.AuthMethod{
            ssh.PublicKeys(key),
        },
        HostKeyCallback: ssh.InsecureIgnoreHostKey(),
    }
    return &SftpProvider{
        BasePath:        basePath,
        SshUrl:          sshUrl,
        SshClientConfig: sshClientConfig,
    }, nil
}

func (sp *SftpProvider) SendFile(sourcePath string, destPathSuffix string) error {
    client, err := ssh.Dial("tcp", sp.SshUrl, sp.SshClientConfig)
    defer func() {
        _ = client.Close()
    }()
    if err != nil {
        return err
    }

    sftpCli, err := sftp.NewClient(client)
    defer func() {
        _ = sftpCli.Close()
    }()
    if err != nil {
        return err
    }

    // Open the source file to copy
    in, err := os.Open(sourcePath)
    defer func() {
        _ = in.Close()
    }()
    if err != nil {
        return err
    }

    fullDestPath := filepath.Join(sp.BasePath, destPathSuffix)

    // Create the destination folders
    if err := sftpCli.MkdirAll(fullDestPath); err != nil {
        return err
    }

    _, fileName := filepath.Split(sourcePath)

    // Open the destination file
    out, err := sftpCli.OpenFile(filepath.Join(fullDestPath, fileName), os.O_CREATE|os.O_WRONLY)
    defer func() {
        _ = out.Close()
    }()
    if err != nil {
        return err
    }

    _, err = io.Copy(out, in)
    if err != nil {
        return err
    }
    return nil
}
