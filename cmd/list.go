package cmd

import (
	"errors"
	"fmt"

	"github.com/fosskey/cli/internal/util"
	"github.com/fosskey/cli/internal/vault"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "ls",
	Short: "List all names",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			return nil
		}
		return fmt.Errorf("\"%s %s\" doesn't accept any argument", rootCmd.Name(), cmd.Name())
	},
	Run: func(cmd *cobra.Command, args []string) {
		masterkey := util.Password("Enter master key: ")
		names, err := vault.List(masterkey)
		if err != nil && err.Error() == "AuthFailed" {
			err = errors.New("incorrect master key")
		}
		cobra.CheckErr(err)
		if len(names) > 0 {
			fmt.Println("Vault")
			lastname := names[len(names)-1]
			names = names[:len(names)-1]
			for _, name := range names {
				fmt.Printf("├──%s\n", name)
			}
			fmt.Printf("└──%s\n", lastname)
		} else {
			fmt.Println("Vault is empty")
		}
	},
	DisableFlagsInUseLine: true,
}

func init() {
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
