package cmd

import "github.com/spf13/cobra"

var rootCmd = &cobra.Command{
	Use:   "nme [command]",
	Short: "Nemesis is a simple reminder system, that works locally",
	Long: `Nemesis can remember dates and times of various events
and give reminders at fixed intervals. It comes in
two parts - a CLI (which is what you are using) to
create, delete, and modify events and a daemon to
send notifications (via D-Bus).`,
}

func Execute() error {
	return rootCmd.Execute()
}
