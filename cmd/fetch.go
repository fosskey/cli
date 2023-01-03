package cmd

import (
	"errors"
	"fmt"

	"github.com/fosskey/cli/internal/util"
	"github.com/fosskey/vault"
	"github.com/spf13/cobra"
)

// fetchCmd represents the fetch command
var fetchCmd = &cobra.Command{
	Use:   "fetch NAME",
	Short: "Fetch a secret",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) == 1 {
			return nil
		}
		return fmt.Errorf("\"%s %s\" requires 1 argument", rootCmd.Name(), cmd.Name())
	},
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]
		masterkey := util.Password("Enter master key: ")
		secret, err := vault.Fetch(masterkey, name)
		if err != nil {
			switch err.Error() {
			case "AuthFailed":
				err = errors.New("incorrect master key")
			case "NotFound":
				err = fmt.Errorf("could not find %q", name)
			}
		}
		cobra.CheckErr(err)
		fmt.Println(secret)
	},
	DisableFlagsInUseLine: true,
}

func init() {
	rootCmd.AddCommand(fetchCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// fetchCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// fetchCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
