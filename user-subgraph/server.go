package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"
	"user-subgraph/graph"

	"asap.local/sys-headers/middleware"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/lru"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi/v5"
	_ "github.com/lib/pq"
	"github.com/vektah/gqlparser/v2/ast"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	db, err := sql.Open("postgres", "dbname=postgres user=postgres password=pg123 sslmode=disable")
	if err != nil {
		log.Fatal("Failed to open database connection: ", err)
	}

	resolver := &graph.Resolver{
		Datastore: db,
	}

	srv := handler.New(graph.NewExecutableSchema(graph.Config{Resolvers: resolver}))

	srv.AddTransport(transport.Options{})
	srv.AddTransport(transport.GET{})
	srv.AddTransport(transport.POST{})

	srv.SetQueryCache(lru.New[*ast.QueryDocument](1000))

	srv.Use(extension.Introspection{})
	srv.Use(extension.AutomaticPersistedQuery{
		Cache: lru.New[string](100),
	})

	mux := chi.NewMux()
	mux.Use(middleware.HTTPSystemHeaders)
	mux.Handle("/", playground.Handler("GraphQL playground", "/query"))
	mux.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, mux))
}
