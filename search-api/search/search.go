package search

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	schemas "app/schemas"
	"time"

	"github.com/olivere/elastic"
	"github.com/teris-io/shortid"
)

const (
	elasticIndexName = "documents"
	elasticTypeName  = "document"
)

var (
	elasticClient *elastic.Client
)

func Start() {
	var err error
	// Create Elastic client and wait for Elasticsearch to be ready
	for {
		elasticClient, err = elastic.NewClient(
			elastic.SetURL("http://elasticsearch:9200"),
			elastic.SetSniff(false),
		)
		if err != nil {
			log.Println(err)
			// Retry every 3 seconds
			time.Sleep(3 * time.Second)
		} else {
			break
		}
	}
}

func CreateDocuments(c context.Context, docs []schemas.DocumentRequest) error {
	// Insert documents in bulk
	bulk := elasticClient.
		Bulk().
		Index(elasticIndexName).
		Type(elasticTypeName)
	for _, d := range docs {
		doc := schemas.Document{
			ID:        shortid.MustGenerate(),
			Title:     d.Title,
			CreatedAt: time.Now().UTC(),
			Content:   d.Content,
		}
		bulk.Add(elastic.NewBulkIndexRequest().Id(doc.ID).Doc(doc))
	}
	if _, err := bulk.Do(c); err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func SearchDocuments(c context.Context, query schemas.DocumentQuery) (schemas.SearchResponse, error) {
	// Perform search
	esQuery := elastic.NewMultiMatchQuery(query.Query, "title", "content").
		Fuzziness("2").
		MinimumShouldMatch("2")

	result, err := elasticClient.Search().
		Index(elasticIndexName).
		Query(esQuery).
		From(query.Skip).Size(query.Take).
		Do(c)
	
	res := schemas.SearchResponse{}
	if err != nil {
		log.Println(err)
		return res, err
	}

	res.Time = fmt.Sprintf("%d", result.TookInMillis)
	res.Hits = fmt.Sprintf("%d", result.Hits.TotalHits)

	// Transform search results before returning them
	docs := make([]schemas.DocumentResponse, 0)
	for _, hit := range result.Hits.Hits {
		var doc schemas.DocumentResponse
		if hit.Source != nil {
			err := json.Unmarshal(*hit.Source, &doc)
			if err != nil {
				log.Println(err)
				continue
			}
			docs = append(docs, doc)
		}
	}

	res.Documents = docs
	
	return res, nil
}