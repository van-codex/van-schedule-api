package elastic

import (
	"github.com/elastic/go-elasticsearch/v8"
	"gopkg.in/ini.v1"
	"log"
	"strings"
)

type Elastic struct {
	client *elasticsearch.Client
	index  string
}

func Init(config *ini.Section) *Elastic {
	var err error
	elastic := new(Elastic)
	elastic.index = config.Key("default_index").String()
	cfg := elasticsearch.Config{
		Addresses: strings.Split(config.Key("hosts").String(), ","),
		Username:  config.Key("username").String(),
		Password:  config.Key("password").String(),
		CloudID:   config.Key("cloud_id").String(),
		APIKey:    config.Key("api_key").String(),
	}
	elastic.client, err = elasticsearch.NewClient(cfg)
	if err != nil {
		log.Fatalln(err)
	}
	return elastic
}
