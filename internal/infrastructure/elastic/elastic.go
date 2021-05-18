package elastic

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/aszanky/goprojectfull/internal/helper"
	"github.com/elastic/go-elasticsearch/v7"
	"github.com/elastic/go-elasticsearch/v7/esapi"
	"github.com/opentracing/opentracing-go"
)

// FlexiDoc can hold a diverse kind of data payload
type FlexibleDoc struct {
	ID    string      `json:"id"`
	Index string      `json:"index"`
	Data  interface{} `json:"data"`
}

type ElasticSearch interface {
	GetIndexName(vertical, name string) string
	Insert(ctx context.Context, doc FlexibleDoc) (err error)
	UpdateByQuery(ctx context.Context, index []string, query string) (err error)
	Update(ctx context.Context, query string, doc FlexibleDoc) (err error)
}

type elasticSearch struct {
	host    string
	elastic *elasticsearch.Client
}

func New(
	host string,
) (ElasticSearch, error) {
	cfg := elasticsearch.Config{
		Addresses: []string{host},
	}

	client, err := elasticsearch.NewClient(cfg)
	if err != nil {
		log.Fatal(err)
	}

	return &elasticSearch{
		host:    host,
		elastic: client,
	}, err
}

// GetIndexName to get elastic index
func (e *elasticSearch) GetIndexName(vertical, name string) string {
	return vertical + "-" + name
}

// Insert remarks
func (e *elasticSearch) Insert(ctx context.Context, doc FlexibleDoc) (err error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "ElasticSearch_Insert")
	defer span.Finish()

	data, err := helper.StructJSON(doc.Data)
	if err != nil {
		return err
	}

	req := esapi.IndexRequest{
		Index:   doc.Index,
		Body:    strings.NewReader(data),
		Refresh: "true",
	}

	// Perform the request with the client.
	res, err := req.Do(ctx, e.elastic)
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}
	defer res.Body.Close()

	fmt.Printf("response %v", res)

	return err
}

func (e *elasticSearch) UpdateByQuery(ctx context.Context, index []string, query string) (err error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "ElasticSearch_UpdateByQuery")
	defer span.Finish()

	queryReader := strings.NewReader(query)

	res, err := e.elastic.UpdateByQuery(
		index,
		e.elastic.UpdateByQuery.WithContext(ctx),
		e.elastic.UpdateByQuery.WithConflicts("proceed"),
		e.elastic.UpdateByQuery.WithBody(queryReader))
	if err != nil {
		return err
	}
	defer res.Body.Close()

	fmt.Printf("response update by query %v", res)

	return
}

//UPDATE data or UPSERT Update if exist, create if does not exist
func (e *elasticSearch) Update(ctx context.Context, query string, doc FlexibleDoc) (err error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "ElasticSearch_UpdateByQuery")
	defer span.Finish()

	res, err := e.elastic.Update(
		doc.Index,
		doc.ID,
		strings.NewReader(query),
	)

	fmt.Printf("response update %v", res)

	return err
}
