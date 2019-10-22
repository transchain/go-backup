package provider

import (
    "github.com/transchain/go-backup/types"
)

type Provider interface {
    SendFile(sourcePath string, destPathSuffix string) error
}

func GetProviders(destCfg *types.DestinationsConfig) map[string]Provider {
    providers := map[string]Provider{}
    for name, cfg := range destCfg.Local {
        providers["local."+name] = NewLocalProvider(cfg.BasePath)
    }
    for name, cfg := range destCfg.Sftp {
        prov, err := NewSftpProvider(cfg.BasePath, cfg.Host, cfg.Port, cfg.User, cfg.SshKeyPath)
        if err == nil {
            providers["sftp."+name] = prov
        }
    }
    return providers
}
