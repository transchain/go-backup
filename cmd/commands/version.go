package commands

import (
    "fmt"

    "github.com/spf13/cobra"

    "github.com/transchain/go-backup/version"
)

func GetVersionCmd() *cobra.Command {
    return &cobra.Command{
        Use:   "version",
        Short: "Show version info",
        Run:   GetVersionRunFn(),
    }
}

func GetVersionRunFn() func(*cobra.Command, []string) {
    return func(cmd *cobra.Command, args []string) {
        fmt.Println(version.Version)
    }
}
