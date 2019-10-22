package types

type ProjectConfig struct {
    Path         string   `yaml:"path"`
    FileNumber   int      `yaml:"fileNumber"`
    CronSpec     string   `yaml:"cronSpec"`
    Task         string   `yaml:"task"`
    BackupFolder string   `yaml:"backupFolder"`
    Destinations []string `yaml:"destinations"`
}

type MailConfig struct {
    Host       string   `yaml:"host"`
    Port       int      `yaml:"port"`
    Username   string   `yaml:"username"`
    Password   string   `yaml:"password"`
    DestEmails []string `yaml:"destEmails"`
}

type SftpConfig struct {
    User       string `yaml:"user"`
    Host       string `yaml:"host"`
    Port       string `yaml:"port"`
    SshKeyPath string `yaml:"sshKeyPath"`
    BasePath   string `yaml:"basePath"`
}

type LocalConfig struct {
    BasePath string `yaml:"basePath"`
}

type DestinationsConfig struct {
    Local map[string]*LocalConfig `yaml:"local"`
    Sftp  map[string]*SftpConfig  `yaml:"sftp"`
}

type Config struct {
    Destinations *DestinationsConfig       `yaml:"destinations"`
    Mail         *MailConfig               `yaml:"mail"`
    Projects     map[string]*ProjectConfig `yaml:"projects"`
}

func NewConfig() *Config {
    return &Config{}
}
