/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/jstone28/quant/pkg/engine"

	"github.com/spf13/cobra"
)

// chatCmd represents the chat command
var chatCmd = &cobra.Command{
	Use:   "chat",
	Short: "open a chat session with quant, GPT-3 and Google/the Internet",
	Long: `Open a prompt with GPT-3. The response is assessed for correctness
and then sent to Google/the Internet for additional information.

Users have the ability to provide feedback to quant to improve the accuracy of
the responses.

Additional in-chat commands are available to [exit] the chat session, get [help], and learn
more [about] quant.

quant is a perpetual work in progress 🚧.
`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: load credentials from viper and the user's config file
		engine.Run()
	},
}

func init() {
	rootCmd.AddCommand(chatCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// chatCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// chatCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
