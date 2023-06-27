package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/olivere/elastic/v7"
	"zangetsu/internal/domain/entity"
)

type AnimeESRepository struct {
	client *elastic.Client
	index  string
}

type IAnimeESRepository interface {
	Index(anime *entity.AnimeViewModel) error
	Search(query string) ([]*entity.AnimeViewModel, error)
	CreateAnimeIndex() error
}

func NewElasticsearchAnimeRepository(client *elastic.Client, index string) *AnimeESRepository {
	return &AnimeESRepository{
		client: client,
		index:  index,
	}
}
func (r *AnimeESRepository) CreateAnimeIndex() error {
	// Check if the index already exists
	exists, err := r.client.IndexExists(r.index).Do(context.Background())
	if err != nil {
		return err
	}

	// If the index doesn't exist, create it
	if !exists {
		createIndex, err := r.client.CreateIndex(r.index).Do(context.Background())
		if err != nil {
			return err
		}
		if !createIndex.Acknowledged {
			return fmt.Errorf("index creation not acknowledged")
		}

		// Define the mapping for the index
		//	mapping := `{
		//    "settings": {
		//        "analysis": {
		//            "analyzer": {
		//                "synonym_analyzer": {
		//                    "tokenizer": "standard",
		//                    "filter": [
		//                        "lowercase",
		//                        "synonym_filter"
		//                    ]
		//                }
		//            },
		//            "filter": {
		//                "synonym_filter": {
		//                    "type": "synonym",
		//                    "synonyms": [
		//                        "хианта, хинаат => хината",   // Replace foo and bar with baz
		//                        "один, 1",            // Replace one with 1
		//                        "бабниу, бабник"         // Replace cat with kitty
		//                    ]
		//                }
		//            }
		//        }
		//    },
		//    "mappings": {
		//        "properties": {
		//            "titleRus": {
		//                "type": "text",
		//                "analyzer": "synonym_analyzer"
		//            },
		//            "description": {
		//                "type": "text",
		//                "analyzer": "synonym_analyzer"
		//            },
		//            "titleEng": {
		//                "type": "keyword"
		//            },
		//            "releaseDate": {
		//                "type": "date",
		//                "format": "yyyy-MM-dd"
		//            }
		//            // Add more fields and their types as needed
		//        }
		//    }
		//}`

		mapping := `{
        "properties": {
            "titleRus": {
                "type": "text"
            }
            // Add more fields and their types as needed
        }
    }`

		// Put the mapping for the index
		putMapping, err := r.client.PutMapping().Index(r.index).BodyString(mapping).Do(context.Background())
		if err != nil {
			return err
		}
		if !putMapping.Acknowledged {
			return fmt.Errorf("mapping not acknowledged")
		}
	}

	return nil
}

func (r *AnimeESRepository) Index(anime *entity.AnimeViewModel) error {
	dataJSON, err := json.Marshal(anime)
	js := string(dataJSON)

	_, err = r.client.Index().
		Index(r.index).
		//Id(anime.ID).
		BodyJson(js).
		Refresh("true").
		Do(context.Background())
	if err != nil {
		return err
	}

	return nil
}

func (r *AnimeESRepository) Search(query string) ([]*entity.AnimeViewModel, error) {

	fuzzyQuery := elastic.NewMatchQuery("titleRus", query).Fuzziness("2") // Adjust fuzziness value as needed
	//prefixQuery := elastic.NewPrefixQuery("titleRus", query)
	//matchQuery := elastic.NewMatchQuery("titleRus", query).Analyzer("russian_latin_analyzer")

	//Create the search query
	//searchQuery := elastic.NewMultiMatchQuery(query).
	//	Field("titleRus").
	//	Field("titleEng").
	//	Field("description")

	//boolQuery := elastic.NewBoolQuery().
	//	Must(fuzzyQuery)
	//Must(searchQuery)

	// Execute the search request
	searchResult, err := r.client.Search().
		Index(r.index).
		Query(fuzzyQuery).
		Size(10).
		Do(context.Background())

	if err != nil {
		return nil, fmt.Errorf("failed to execute search query: %v", err)
	}

	// Parse and return the search results
	var animeList []*entity.AnimeViewModel
	for _, hit := range searchResult.Hits.Hits {
		var anime entity.AnimeViewModel
		err := json.Unmarshal(hit.Source, &anime)
		if err != nil {
			return nil, fmt.Errorf("failed to parse search result: %v", err)
		}
		animeList = append(animeList, &anime)
	}

	return animeList, nil
}
