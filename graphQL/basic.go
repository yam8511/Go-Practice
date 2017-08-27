package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/graphql-go/graphql"
)

func main1() {
	// Schema
	Hello := &graphql.Field{
		Type: graphql.String,
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			data := []string{"World", "Zuolar"}
			log.Println("Param", p.Args)
			index, exists := p.Args["id"]
			if exists {
				return data[index.(int)], nil
			}
			return data, nil
		},
		Args: map[string]*graphql.ArgumentConfig{
			"id": &graphql.ArgumentConfig{
				Type:        graphql.Int,
				Description: "Just a ID",
			},
		},
	}

	fields := graphql.Fields{
		"hello": Hello,
	}
	rootQuery := graphql.ObjectConfig{Name: "RootQuery", Fields: fields}
	schemaConfig := graphql.SchemaConfig{Query: graphql.NewObject(rootQuery)}
	schema, err := graphql.NewSchema(schemaConfig)
	if err != nil {
		log.Fatalf("failed to create new schema, error: %v", err)
	}

	// Query
	query := `
		{
			hello(id:1)
		}
	`
	params := graphql.Params{Schema: schema, RequestString: query}
	r := graphql.Do(params)
	if r.HasErrors() {
		log.Fatalf("failed to execute graphql operation, errors: %+v", r.Errors)
	}
	rJSON, _ := json.Marshal(r)
	fmt.Printf("%s \n", rJSON) // {“data”:{“hello”:”world”}}
}
