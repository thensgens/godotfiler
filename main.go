package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

var (
	default_source_dir = "/home/thens/Dropbox/linux/dotfiles"
	default_target_dir = "/home/thens"
	files              = []string{".bashrc", ".vimrc", ".tmux.conf"}
)

type FlagResult struct {
	name   *string
	source string
	target string
}

func checkFlags() FlagResult {
	if len(os.Args) <= 1 {
		usage()
	}
	option := flag.String("option", "restore", "Option can either be restore or archive.")
	flag.Parse()
	var result FlagResult
	if *option == "restore" {
		result := FlagResult{name: option, source: default_source_dir, target: default_target_dir}
	} else {
		result := FlagResult{name: option, source: default_target_dir, target: default_source_dir}
	}
	return result
}

func usage() {
	fmt.Println("Godotfiles has to be invoked with the -option flag.")
	fmt.Println("Possible values are: restore / archive")
	log.Fatal("Exiting..")
}

func main() {
	flagOption := checkFlags()
	dotfiler := &Dotfiler{option: flagOption}
	err := dotfiler.execute()
	if err != nil {
		log.Fatalf("Error occured: %s\nExiting..", err)
	}
}
