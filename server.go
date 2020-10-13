package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/handler"
	auth "github.com/Salaton/screening-test/auth"
	"github.com/Salaton/screening-test/graph"
	"github.com/Salaton/screening-test/graph/generated"
	db "github.com/Salaton/screening-test/postgres"
	"github.com/go-chi/chi"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	db, err := InitDB()
	if err != nil {
		log.Fatalf("FAILED TO CONNECT TO DB %v", err.Error())
	}

	router := chi.NewRouter()

	router.Use(auth.Middleware(db))
	srv := handler.GraphQL(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	router.Handle("/", handler.Playground("GraphQL playground", "/query"))
	router.Handle("/query", srv)

	// srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}

// InitDB function to start the db connections process
func InitDB() (db.DBClient, error) {
	graph.DB = &db.PostgresClient{}
	dsn := "user=sala password=$krychowiak-254$ dbname=savannahtest port=5432 sslmode=disable TimeZone=Africa/Nairobi"
	return graph.DB, graph.DB.Open(dsn)
}
