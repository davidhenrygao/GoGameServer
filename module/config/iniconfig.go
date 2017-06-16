package config

import (
	"github.com/Unknwon/goconfig"
	"strings"
)

//Note that there is a rule for the parameter string key:
//format: "section,key" or "key"
//the default section is "DEFAULT"
//the real "key" string can't contain comma symbol
func decode(key string) (string, string) {
	var s, k string
	ss := strings.Split(key, ",")
	if len(ss) == 1 {
		s = goconfig.DEFAULT_SECTION
		k = strings.Trim(ss[0], " ")
	} else {
		s = strings.Trim(ss[0], " ")
		k = strings.Trim(ss[1], " ")
	}
	return s, k
}

type iniCfg struct {
	cf *goconfig.ConfigFile
}

func (r *iniCfg) Value(key string) (string, error) {
	s, k := decode(key)
	return r.cf.GetValue(s, k)
}

func (r *iniCfg) Int(key string) (int, error) {
	s, k := decode(key)
	return r.cf.Int(s, k)
}

func (r *iniCfg) Bool(key string) (bool, error) {
	s, k := decode(key)
	return r.cf.Bool(s, k)
}

func (r *iniCfg) Set(key, val string) {
	s, k := decode(key)
	r.cf.SetValue(s, k, val)
}

func loadIniConfigFile(file string) (CfgInf, error) {
	var cfg iniCfg
	var err error
	cfg.cf, err = goconfig.LoadConfigFile(file)
	return &cfg, err
}

func init() {
	registerCfgLoadfunc("ini", loadIniConfigFile)
}
