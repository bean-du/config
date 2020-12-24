package conf

import (
	"encoding/json"
	"io/ioutil"
)

var Conf = &Config{}

type Config struct {
	Port string `json:"port"`

	Postgre struct {
		Databases string `json:"databases"`
		Username  string `json:"username"`
		Password  string `json:"password"`
		Host      string `json:"host"`
		Port      int    `json:"port"`
	} `json:"postgre"`

	Etcd struct {
		Endpoint []string `json:"endpoint"`
		Timeout  int64    `json:"timeout"`
		Username string   `json:"username"`
		Password string   `json:"password"`
	} `json:"etcd"`
}

func Init(filename string) []func() {
	c, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	if err := json.Unmarshal(c, Conf); err != nil {
		panic(err)
	}

	cancel := InitDB()
	return []func(){cancel}
}
