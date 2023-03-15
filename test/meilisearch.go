package test

import (
	// "encoding/json"
	"fmt"
	// "gopkg.in/yaml.v2"
	"github.com/meilisearch/meilisearch-go"
)
 


func Meilisearch() *meilisearch.SearchResponse {
	client := meilisearch.NewClient(meilisearch.ClientConfig{
	  Host: "http://47.99.75.203:7700",
	})
	res , _ :=client.Index("movies").Search("botman", &meilisearch.SearchRequest{})
	// data , _ :=json.Marshal(res.Hits)
	for _,i:=range res.Hits{
		// fmt.Println(i)
		ok,_:=i.(map[string]*interface{})
		fmt.Printf("%v %T\n",ok["id"],ok["id"])
	}
	return res
  }
  