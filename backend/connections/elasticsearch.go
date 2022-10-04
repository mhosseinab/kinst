package connections

import (
	"github.com/olivere/elastic/v7"
	"tools"
)

func GetElasticsearch() (*elastic.Client, error) {
	client, err := elastic.NewClient(
		elastic.SetURL(tools.GetEnv("ES_HOST", "http://127.0.0.1:9200/")),
	)

	return client, err
}
