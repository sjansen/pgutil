package main

import (
	"fmt"
	"os"

	"github.com/sjansen/pgutil/internal/cli"
	"github.com/sjansen/pgutil/internal/commands"
	"github.com/sjansen/pgutil/internal/db"
)

var build string // set by goreleaser

func main() {
	if build == "" {
		build = version
	}
	parser := cli.RegisterCommands(build)

	cmd, err := parser.Parse(os.Args[1:])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	err = cmd.Run(os.Stdout, os.Stderr, &commands.Dependencies{
		DB: func(opts map[string]string) (commands.DB, error) {
			return db.New(opts)
		},
	})
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
