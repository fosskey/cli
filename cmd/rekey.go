package cmd

import (
	"errors"
	"fmt"

	"github.com/fosskey/cli/internal/util"
	"github.com/fosskey/cli/internal/vault"
	"github.com/spf13/cobra"
)

// rekeyCmd represents the rekey command
var rekeyCmd = &cobra.Command{
	Use:   "rekey",
	Short: "Change the master key",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			return nil
		}
		return fmt.Errorf("\"%s %s\" doesn't accept any argument", rootCmd.Name(), cmd.Name())
	},
	Run: func(cmd *cobra.Command, args []string) {
		masterkey := util.Password("Enter old master key: ")

		// Verify master key
		verified, err := vault.Verify(masterkey)
		if !verified {
			err = errors.New("incorrect master key")
		}
		cobra.CheckErr(err)

		newkey := util.Password("Enter new master key: ")

		if err = vault.Rekey(masterkey, newkey); err != nil {
			if err.Error() == "VaultEmpty" {
				err = errors.New("vault is empty, rekey requires non-empty vault")
			}
			cobra.CheckErr(err)
		}

		fmt.Printf("Masterkey is now changed\n")
	},
	DisableFlagsInUseLine: true,
}

func init() {
	rootCmd.AddCommand(rekeyCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// rekeyCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// rekeyCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
