package module

import (
	"gopkg.in/yaml.v2"
	"github.com/meilisearch/meilisearch-go"
	"io/ioutil"
)

type ConfigInfo struct{
	Port string `yaml:"port"`
	Ip string `yaml:"ip"`
}

var conf  ConfigInfo
var client = connect(conf.Readconfig())

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
	res , _ :=client.Index(index).Search(keywork, &meilisearch.SearchRequest{
		AttributesToRetrieve:[]string{"id","name","describe"},
	})
	return res.Hits
}

func Insert(myjson map[string]interface{}, index string) string{
	res , err := client.Index(index).AddDocuments(myjson)
	if err != nil{
		panic(err)
	}
	return string(res.Status)
}

func Delete(id string, index string) string{
	res ,err := client.Index(index).DeleteDocument(id)
	if err !=nil{
		panic(err)
	}
	
	return string(res.Status)
}