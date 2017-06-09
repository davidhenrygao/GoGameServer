package config

import (
	"fmt"
	"path"
)

//CfgInf is a interface define some operations to get or set the configure value.
type CfgInf interface {
	Value(key string) (string, error)
	Int(key string) (int, error)
	Bool(key string) (bool, error)
	Set(key, val string)
}

type loadfunc func(string) (CfgInf, error)

var cfgMap = make(map[string]loadfunc)
var defaultCfgType = "ini"

//Cfg is the global configuration object
//initiated by LoadCfgFile or LoadCfgFileByType.
var Cfg CfgInf

func registerCfgLoadfunc(name string, f loadfunc) {
	cfgMap[name] = f
}

func cfgLoadfunc(name string) (loadfunc, bool) {
	if cfgMap[name] == nil {
		return nil, false
	}
	return cfgMap[name], true
}

func ifCfgTypeExist(t string) bool {
	return cfgMap[t] != nil
}

//DefaultCfgType returns the default configure file type.
func DefaultCfgType() string {
	return defaultCfgType
}

//SetDefaultCfgType set the default configure file type to the specific type.
func SetDefaultCfgType(t string) error {
	var err error
	if ifCfgTypeExist(t) {
		defaultCfgType = t
	} else {
		err = fmt.Errorf("unknown configure file type %s", t)
	}
	return err
}

//LoadCfgFile loads the specific configure file.
//It will use the file extension as configre file type if supported,
//otherwise use the default configure file type.
//return error if fails.
func LoadCfgFile(file string) error {
	var t string
	ext := path.Ext(file)
	if ext == "" {
		t = defaultCfgType
	} else {
		t = ext[1:]
		if !ifCfgTypeExist(t) {
			t = defaultCfgType
		}
	}
	return LoadCfgFileByType(file, t)
}

//LoadCfgFileByType loads the specific configure file by specific configure
//file type.
//Success return true, or return false while failed.
func LoadCfgFileByType(file string, t string) error {
	f, ok := cfgLoadfunc(t)
	if !ok {
		return fmt.Errorf("unknown configure file type %s", t)
	}
	var err error
	Cfg, err = f(file)
	return err
}
