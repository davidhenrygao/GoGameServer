package main

import (
	"flag"
	"fmt"
)

var cfgFile string

func init() {
	var defaultPath = ""
	flag.StringVar(&cfgFile, "cfgfile", defaultPath, "configure file path.")
	flag.StringVar(&cfgFile, "f", defaultPath, "configure file path (short-hand).")
}

func main() {
	flag.Parse()
	if cfgFile == "" {
		fmt.Println("Usage: program -f filepath.\n")
		return
	}
	initCfg := loadConfig(cfgFile)
	fmt.Printf("initCfg = %+v\n", initCfg)
	initServer(initCfg)
}
