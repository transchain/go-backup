package provider

import (
    "io"
    "os"
    "path/filepath"
)

type LocalProvider struct {
    BasePath string
}

func NewLocalProvider(basePath string) *LocalProvider {
    return &LocalProvider{BasePath: basePath}
}

func (lp *LocalProvider) SendFile(sourcePath string, destPathSuffix string) error {
    // Open the source file to copy
    in, err := os.Open(sourcePath)
    defer func() {
        _ = in.Close()
    }()
    if err != nil {
        return err
    }

    fullDestPath := filepath.Join(lp.BasePath, destPathSuffix)

    // Create the destination folders
    if err := os.MkdirAll(fullDestPath, 0770); err != nil {
        return err
    }

    _, fileName := filepath.Split(sourcePath)

    // Open the destination file
    out, err := os.OpenFile(filepath.Join(fullDestPath, fileName), os.O_CREATE|os.O_WRONLY, 0660)
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
