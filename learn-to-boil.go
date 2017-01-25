package main

import (
	"fmt"
	"github.com/magleff/learn-to-boil/cmd"
	"os"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}
