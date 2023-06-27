package database

import (
	"context"
	"fmt"
	"github.com/olivere/elastic/v7"
)

func InitESDb() (*elastic.Client, error) {
	// Initialize Elasticsearch client
	username := "elastic"
	password := "VfU6SCFaD34fpVxxStV7"
	esClient, err := elastic.NewClient(elastic.SetURL("http://localhost:9200"), elastic.SetBasicAuth(username, password))
	if err != nil {
		return nil, fmt.Errorf("failed to create Elasticsearch client: %w", err)
	}

	// Ping Elasticsearch to check if it's running
	info, code, err := esClient.Ping("http://localhost:9200").Do(context.Background())
	if err != nil {
		return nil, fmt.Errorf("failed to ping Elasticsearch: %w", err)
	}

	if code != 200 {
		return nil, fmt.Errorf("Elasticsearch ping returned non-200 status code: %d", code)
	}

	fmt.Println("Elasticsearch is running")
	fmt.Printf("Cluster: %s\n", info.ClusterName)
	fmt.Printf("Version: %s\n", info.Version.Number)

	return esClient, nil
}
