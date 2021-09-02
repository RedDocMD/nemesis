package main

import (
	"github.com/RedDocMD/nemesis/cmd"
	"github.com/RedDocMD/nemesis/event"
	"github.com/spf13/viper"
)

func main() {
	err := cmd.Execute()
	if err == nil {
		dbPath := viper.GetString("dbPath")
		event.DumpEvents(dbPath, cmd.Events())
	}
}
