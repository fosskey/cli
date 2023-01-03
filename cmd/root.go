package cmd

import (
	"os"

	"github.com/fosskey/cli/internal/util"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:     "foss",
	Short:   "A free, open-source, secure, and self-custodial keychain",
	Version: "0.0.0",
	CompletionOptions: cobra.CompletionOptions{
		DisableDefaultCmd:   false,
		HiddenDefaultCmd:    true,
		DisableDescriptions: true,
	},
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// Custom usage template
	rootCmd.SetUsageTemplate(util.UsageTemplate())
	rootCmd.SetHelpTemplate(util.HelpTemplate())

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.foss.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
