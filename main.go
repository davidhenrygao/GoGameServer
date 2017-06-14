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
		fmt.Println("Usage: program -cfgfile/-f filepath.\n")
		return
	}
	iCfg := loadConfig(cfgFile)
	fmt.Printf("iCfg = %+v\n", iCfg)
}
