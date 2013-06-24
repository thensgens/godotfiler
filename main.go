package main

import (
	"flag"
	"fmt"
	"os"
)

var (
	default_source_dir = "/home/thens/Dropbox/linux/dotfiles"
	default_target_dir = "/home/thens"

	// just for testing purposes
	// default_source_dir = "/home/thens/gocode/src/github.com/thensgens/godotfiler/test_files/source"
	// default_target_dir = "/home/thens/gocode/src/github.com/thensgens/godotfiler/test_files/target"

	// static configuration for now
	files = []string{".bashrc", ".vimrc", ".tmux.conf"}
)

type FlagResult struct {
	name    *string
	source  string
	target  string
	verbose bool
}

func checkFlags() FlagResult {
	if len(os.Args) <= 1 {
		usage()
	}
	flag.Usage = func() {
		fmt.Println("Godotfiles has to be invoked with the --option flag.")
		fmt.Println("Possible values are: restore / archive. Exiting..")
		os.Exit(-1)
	}
	option := flag.String("option", "restore", "Option can either be restore or archive.")
	verbose_long := flag.Bool("verbose", false, "Verbose command flag.")
	verbose_short := flag.Bool("v", false, "Verbose command flag.")
	flag.Parse()
	verbose := *verbose_long || *verbose_short

	var result FlagResult
	if *option == "restore" {
		result = FlagResult{name: option, source: default_source_dir, target: default_target_dir, verbose: verbose}
	} else {
		result = FlagResult{name: option, source: default_target_dir, target: default_source_dir, verbose: verbose}
	}
	return result
}

func usage() {
	fmt.Println("Godotfiler has to be invoked with the --option flag.")
	fmt.Println("Possible values are: restore / archive. Exiting..")
	os.Exit(-1)
}

func main() {
	// sanity cli flag checks
	flagOption := checkFlags()
	dotfiler := &Dotfiler{option: flagOption, files: files}
    if err := dotfiler.execute(); err != nil {
        fmt.Println(err)
    }
}
