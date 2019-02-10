package main

// This is an example code copied from https://github.com/graph-gophers/graphql-go
// curl -XPOST -d '{"query": "{ hello }"}' localhost:8080/query
import (
	"log"
	"net/http"

	graphql "github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
)

type query struct{}

func (*query) Hello() string { return "Hello, world!" }

func main() {
	s := `
                schema {
                        query: Query
                }
                type Query {
                        hello: String!
                }
        `
	schema := graphql.MustParseSchema(s, &query{})
	http.Handle("/query", &relay.Handler{Schema: schema})
	log.Fatal(http.ListenAndServe(":8080", nil))
}
