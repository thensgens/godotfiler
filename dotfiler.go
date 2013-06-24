package main

import (
	"fmt"
	"io/ioutil"
	"path"
    "bufio"
    "os"
)

type Dotfiler struct {
	option     FlagResult
	exceptions []string
	files      []string
}

func (self *Dotfiler) execute() error {
    output := "/home/thens/.vim/bundle"
    var err error

	if self.option.verbose {
		self.logStart()
	}
	for _, file := range self.files {
		// these calls could be dispatched as goroutines later on
	    if err = self.processElement(file); err != nil {
            return err
        }
	}

    files, _ := ioutil.ReadDir(output)
    allFiles := make([]string, 0)
    for _, fileObj := range files {
        if fileObj.IsDir() {
            allFiles = append(allFiles, fileObj.Name())
        }
    }
    if err = self.savePluginNames(allFiles); err != nil {
        return err
    }

	if self.option.verbose {
		self.logStop()
	}
    return nil
}

func (self *Dotfiler) processElement(file string) error {
	input := path.Join(self.option.source, file)
	output := path.Join(self.option.target, file)

	if self.option.verbose {
		fmt.Printf("Copying %s  -->  %s\n", input, output)
	}
	b, err := ioutil.ReadFile(input)
	if err != nil {
		panic(err)
	}

	err = ioutil.WriteFile(output, b, 0644)
	if err != nil {
		panic(err)
	}
    return nil
}

func (self *Dotfiler) savePluginNames(allFiles []string) error {
    output := "/home/thens/Dropbox/linux/vim_plugins"

	if self.option.verbose {
		fmt.Printf("Copying plugin dir names to --> %s", output)
	}

    fo, err := os.Create(output)
    if err != nil { return err }
    // close fo on exit and check for its returned error
    defer func() {
        if err := fo.Close(); err != nil {
            panic(err)
        }
    }()
    w := bufio.NewWriter(fo)

    for _, dirName := range allFiles {
        if _, err = w.WriteString(dirName + "\n"); err != nil {
            return err
        }
    }
    if err = w.Flush(); err != nil { return err }
    return nil
}

func (self *Dotfiler) logStart() {
	fmt.Printf("Started task [%s]\n", *self.option.name)
}

func (self *Dotfiler) logStop() {
	fmt.Printf("Stopped task [%s]\n", *self.option.name)
}
