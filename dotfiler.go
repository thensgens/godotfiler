package main

import (
	"fmt"
	"io/ioutil"
	"path"
)

type Dotfiler struct {
	option     FlagResult
	exceptions []string
	files      []string
}

func (self *Dotfiler) execute() {
	if self.option.verbose {
		self.logStart()
	}
	for _, file := range self.files {
		// these calls could be dispatched as goroutines later on
		self.processElement(file)
	}
	if self.option.verbose {
		self.logStop()
	}
}

func (self *Dotfiler) processElement(file string) {
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
}

func (self *Dotfiler) logStart() {
	fmt.Printf("Started task [%s]\n", *self.option.name)
}

func (self *Dotfiler) logStop() {
	fmt.Printf("Stopped task [%s]\n", *self.option.name)
}
