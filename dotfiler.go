package main

import (
	"fmt"
)

type Dotfiler struct {
	option     FlagResult
	exceptions []string
}

func (self *Dotfiler) execute() error {
	self.logStart()
	return nil
}

func (self *Dotfiler) logStart() {
	fmt.Printf("Starting task [%s]\n", *self.option.name)
}
