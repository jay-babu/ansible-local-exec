/*
Copyright © 2022 Jay Patel 36803168+jay-babu@users.noreply.github.com
*/
package cmd

import (
	"log"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

type source struct {
	file string
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "ansible-local-exec",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		file, _ := cmd.Flags().GetString("source")
		s := source{
			file: file,
		}

		local := exec.Command("ansible-playbook", "--connection", "local", "--inventory", "127.0.0.1,", "-K", s.file)
		local.Stdout = os.Stdout
		local.Stderr = os.Stderr
		err := local.Run()
		if err != nil {
			log.Fatal(err)
		}
	},
}

/*
*
Execute adds all child commands to the root command and sets flags appropriately.
This is called by main.main(). It only needs to happen once to the rootCmd.
*/
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

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.ansible-local-exec.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().StringP("source", "s", "", "File to run")
	rootCmd.MarkFlagRequired("source")
}
