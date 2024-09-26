package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/mrrizkin/boot/system"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}

func run() error {
	migrate := flag.Bool("migrate", false, "migrate database")
	seeder := flag.Bool("seeder", false, "seeding database")
	flag.Parse()

	switch {
	case *migrate:
		return system.MigrateDB()
	case *seeder:
		return system.SeedDB()
	case len(os.Args) > 1 && os.Args[1] == "run":
		return system.Run()
	default:
		return fmt.Errorf("no valid command provided. Use 'run' to start the server")
	}
}
