package module

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"github.com/meilisearch/meilisearch-go"
	"io/ioutil"
)

type ConfigInfo struct{
	Port string `yaml:"port"`
	Ip string `yaml:"ip"`
}


func (c *ConfigInfo) Readconfig() *ConfigInfo {
	yamlConfig, err := ioutil.ReadFile("config/meilisearch.yaml")
	if err !=nil{
		panic(err)
	}
	err = yaml.Unmarshal(yamlConfig, c)
	if err != nil{
		panic(err)
	}
	return c
}

func connect(conf *ConfigInfo) *meilisearch.Client{
	client := meilisearch.NewClient(meilisearch.ClientConfig{
		Host : "http://" + conf.Ip + ":" + conf.Port,
	  })
	return client
}

func Search(keywork string,index string) []interface{}{
	var conf  ConfigInfo
	client := connect(conf.Readconfig())
	fmt.Println(index,keywork)
	res , _ :=client.Index(index).Search(keywork, &meilisearch.SearchRequest{})
	return res.Hits
}