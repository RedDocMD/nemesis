package cmd

import (
	"os"
	"path"

	"github.com/RedDocMD/nemesis/event"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var rootCmd = &cobra.Command{
	Use:   "nme [command]",
	Short: "Nemesis is a simple reminder system, that works locally",
	Long: `Nemesis can remember dates and times of various events
and give reminders at fixed intervals. It comes in
two parts - a CLI (which is what you are using) to
create, delete, and modify events and a daemon to
send notifications (via D-Bus).`,
}

var events []event.Event

func Events() []event.Event {
	return events
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	initConfig()

	dbPath := viper.GetString("dbPath")
	localEvents, err := event.GetEvents(dbPath)
	cobra.CheckErr(err)
	events = append(events, localEvents...)

	rootCmd.AddCommand(addCmd)
	rootCmd.AddCommand(listCmd)
	rootCmd.AddCommand(completeCmd)
	rootCmd.AddCommand(editCmd)
}

func initConfig() {
	home, err := os.UserHomeDir()
	cobra.CheckErr(err)

	viper.AddConfigPath(path.Join(home, ".config", "nemesis"))
	viper.AddConfigPath(path.Join(home, ".nemesis"))
	viper.SetConfigType("yaml")
	viper.SetConfigName("config")

	viper.SetDefault("dbPath", path.Join(home, ".nemesisDB.json"))

	viper.ReadInConfig()
}
