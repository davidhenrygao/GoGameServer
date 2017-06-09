package main

import (
	"fmt"
	"github.com/davidhenrygao/GoGameServer/utils/config"
	"reflect"
)

//Required configurations when server start.
type initCfg struct {
	Id          int
	Ip          string
	Port        int
	LogDir      string
	LogFileName string
	LogLevel    string
	Logstderr   bool
}

type loadField struct {
	fName   string
	fType   string
	fDefult interface{} //Can't be omitted if is nil.
}

//Must match the fields in the struct initCfg.
var cfgMap = map[string]loadField{
	"Id":          {"id", "int", nil},
	"Ip":          {"ip", "string", "127.0.0.1"},
	"Port":        {"port", "int", nil},
	"LogDir":      {"log,dir", "string", "./log/"},
	"LogFileName": {"log,filename", "string", "server"},
	"LogLevel":    {"log,level", "string", "error"},
	"Logstderr":   {"log,stderr", "bool", false},
}

func loadConfig(file string) {
	err := config.LoadCfgFile(cfgFile)
	if err != nil {
		fmt.Printf("Load configure file error: %+v\n", err)
		return
	}
	cfg := config.Cfg
	var iCfg initCfg
	var val interface{}
	for name, field := range cfgMap {
		switch field.fType {
		case "int":
			val, err = cfg.Int(field.fName)
		case "bool":
			val, err = cfg.Bool(field.fName)
		case "string":
			val, err = cfg.Value(field.fName)
		default:
			panic(fmt.Sprintf("Unknown cfgMap loadField fType: %s.", field.fType))
		}
		if err != nil {
			fmt.Printf("%+v\n", err)
			if field.fDefult == nil {
				panic(fmt.Sprintf("%+v\n", err))
			} else {
				val = field.fDefult
			}
		}
		v := reflect.ValueOf(&iCfg).Elem().FieldByName(name)
		switch field.fType {
		case "int":
			i := val.(int)
			v.SetInt(int64(i))
		case "bool":
			b := val.(bool)
			v.SetBool(b)
		case "string":
			s := val.(string)
			v.SetString(s)
		default:
		}
	}
	fmt.Printf("iCfg = %+v\n", iCfg)
}
