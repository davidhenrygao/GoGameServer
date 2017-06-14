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
	name   string
	defult interface{} //Can't be omitted if is nil.
}

//Must match the fields in the struct initCfg.
var cfgMap = map[string]loadField{
	"Id":          {"id", nil},
	"Ip":          {"ip", "127.0.0.1"},
	"Port":        {"port", nil},
	"LogDir":      {"log,dir", "./log/"},
	"LogFileName": {"log,filename", "server"},
	"LogLevel":    {"log,level", "error"},
	"Logstderr":   {"log,stderr", false},
}

func loadConfig(file string) initCfg {
	err := config.LoadCfgFile(cfgFile)
	if err != nil {
		panic(fmt.Sprintf("Load configure file error: %+v\n", err))
	}
	cfg := config.Cfg
	var iCfg initCfg
	t := reflect.TypeOf(iCfg)
	v := reflect.ValueOf(&iCfg).Elem()
	nf := t.NumField()
	var val interface{}
	for i := 0; i < nf; i++ {
		f := t.Field(i)
		name := f.Name
		loadf, ok := cfgMap[name]
		if !ok {
			panic(fmt.Sprintf("Miss load field(%s) in the cfgMap.", name))
		}
		kind := f.Type.Kind()
		switch kind {
		case reflect.Bool:
			val, err = cfg.Bool(loadf.name)
		case reflect.Int:
			val, err = cfg.Int(loadf.name)
		case reflect.String:
			val, err = cfg.Value(loadf.name)
		default:
			panic(fmt.Sprintf("Unsupported configure type: %s.", kind.String()))
		}
		if err != nil {
			if loadf.defult == nil {
				panic(fmt.Sprintf("%+v\n", err))
			} else {
				val = loadf.defult
			}
		}
		v.FieldByName(name).Set(reflect.ValueOf(val))
	}
	return iCfg
}
