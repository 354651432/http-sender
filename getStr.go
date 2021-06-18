package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
)

func getStr(fileName string) ([]byte, error) {
	if fileName == "" {
		stat, _ := os.Stdin.Stat()
		if stat.Mode()&os.ModeCharDevice != 0 {
			fmt.Println("usage:")
			flag.PrintDefaults()
			os.Exit(0)
		}
		return ioutil.ReadAll(os.Stdin)
	} else {
		return ioutil.ReadFile(fileName)
	}
}
