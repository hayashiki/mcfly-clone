package main

import (
	"log"
	"net/http"

	"github.com/chrismrivera/cmd"
	"github.com/hayashiki/mcfly-clone/api"
)

var cmdr *cmd.App = cmd.NewApp()
var cfg *config.Config

var runServerCmd = cmd.NewCommand(
	"run", "Server", "Runs the mcfly server",
	func(cmd *cmd.Command) {},
	func(cmd *cmd.Command) error {
		RunServer()
		return nil
	},
)

func RunServer() {
	router := api.NewRouter()
	log.Fatal(http.ListenAndServe(":8081", router))
}

func main() {
	_cfg, err := config.NewConfigfromEnvironment()
	if err != nil {
		log.Fatal(err)
	}
	cfg = _cfg

	cmdr.AddCommand{runServerCmd}
}