//go:generate go-bindata -pkg assets -o ../assets/tmpl.go ../tmpl/

package main

import (
    // "git.transchain.fr/tools/go-backup/commands"
    "github.com/transchain/go-backup/cmd/commands"
    "github.com/transchain/go-backup/types"
)

func main() {
    ctx := types.NewContext()

    rootCmd := commands.GetRootCmd()
    rootCmd.AddCommand(
        commands.GetRunCmd(ctx),
        commands.GetVersionCmd(),
    )

    cmd := commands.PrepareBaseCmd(rootCmd)
    if err := cmd.Execute(); err != nil {
        panic(err)
    }
}
