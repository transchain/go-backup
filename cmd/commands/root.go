package commands

import (
    "github.com/spf13/cobra"
    "github.com/spf13/viper"
)

func GetRootCmd() *cobra.Command {
    cmd := &cobra.Command{
        Use:   "backup",
        Short: "Backup tool",
    }
    cmd.PersistentFlags().StringP(ConfigFlag, ConfigShorthand, "./config.yml", "Specify the path to the config file")
    return cmd
}

func PrepareBaseCmd(cmd *cobra.Command, fs ...CobraCmdFunction) *cobra.Command {
    fs = append(fs, cmd.PersistentPreRunE)
    fs = append([]CobraCmdFunction{getBindFlagsPersistentPreRunE()}, fs...)
    cmd.PersistentPreRunE = concatCobraCmdFunctions(fs...)
    return cmd
}

type CobraCmdFunction func(cmd *cobra.Command, args []string) error

func concatCobraCmdFunctions(fs ...CobraCmdFunction) CobraCmdFunction {
    return func(cmd *cobra.Command, args []string) error {
        for _, f := range fs {
            if f != nil {
                if err := f(cmd, args); err != nil {
                    return err
                }
            }
        }
        return nil
    }
}

func getBindFlagsPersistentPreRunE() CobraCmdFunction {
    return func(cmd *cobra.Command, args []string) error {
        return viper.BindPFlags(cmd.Flags())
    }
}
