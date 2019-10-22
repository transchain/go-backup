package types

import (
    "path/filepath"
    "strings"

    "github.com/spf13/viper"
)

type Context struct {
    Config *Config
}

func NewContext() *Context {
    return &Context{
        Config: NewConfig(),
    }
}

func (c *Context) LoadConfig(path string) error {
    v := viper.New()
    configPath, configFile := filepath.Split(path)
    v.SetConfigName(strings.TrimSuffix(configFile, filepath.Ext(configFile)))
    v.AddConfigPath(configPath)
    v.SetConfigType("yaml")
    err := v.ReadInConfig()
    if err != nil {
        return err
    }
    return v.Unmarshal(c.Config)
}
