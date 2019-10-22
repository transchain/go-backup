package commands

import (
    "github.com/spf13/cobra"
    "github.com/spf13/viper"

    "github.com/transchain/go-backup/backup"
    "github.com/transchain/go-backup/mail"
    "github.com/transchain/go-backup/provider"
    "github.com/transchain/go-backup/types"
)

func GetRunCmd(ctx *types.Context) *cobra.Command {
    return &cobra.Command{
        Use:   "run",
        Short: "Run the backup tool",
        RunE:  GetRunRunEFn(ctx),
    }
}

func GetRunRunEFn(ctx *types.Context) func(*cobra.Command, []string) error {
    return func(cmd *cobra.Command, args []string) error {

        // Load the specified config
        configPath := viper.GetString(ConfigFlag)
        err := ctx.LoadConfig(configPath)
        if err != nil {
            return err
        }

        providers := provider.GetProviders(ctx.Config.Destinations)
        backuper := backup.NewBackuper(providers, ctx.Config.Projects)

        projectsCustomErrors := backuper.Run()

        if len(projectsCustomErrors) > 0 {
            return mail.Notify(ctx.Config.Mail, projectsCustomErrors)
        }

        return nil
    }
}
