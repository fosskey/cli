package cmd

import (
	"errors"
	"fmt"

	"github.com/fosskey/cli/internal/util"
	"github.com/fosskey/vault"
	"github.com/spf13/cobra"
)

// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use:   "update NAME",
	Short: "Update a secret",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) == 1 {
			return nil
		}
		return fmt.Errorf("\"%s %s\" requires 1 argument", rootCmd.Name(), cmd.Name())
	},
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]
		masterkey := util.Password("Enter master key: ")

		// Verify master key
		verified, err := vault.Verify(masterkey)
		if !verified {
			err = errors.New("incorrect master key")
		}
		cobra.CheckErr(err)

		// If name doesn't exists
		if _, err := vault.Fetch(masterkey, name); err != nil {
			if err.Error() == "NotFound" {
				err = fmt.Errorf("%q doesn't exists", name)
			}
			cobra.CheckErr(err)
		}

		secret := util.Password("Enter new secret: ")
		err = vault.Update(masterkey, name, secret)
		cobra.CheckErr(err)
		fmt.Printf("%s is now updated in the vault\n", name)
	},
	DisableFlagsInUseLine: true,
}

func init() {
	rootCmd.AddCommand(updateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// updateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// updateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
