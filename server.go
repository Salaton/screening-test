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
	"github.com/joho/godotenv"
)

const defaultPort = "8080"

type environmentVar struct {
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	// Load our .env file
	err := godotenv.Load()
	if err != nil {
		log.Println(".env file could not be found in this directory.")
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

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}

// InitDB function to start the db connections process
func InitDB() (db.DBClient, error) {
	graph.DB = &db.PostgresClient{}
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	host := os.Getenv("DB_HOST")
	dsn := "user=" + user + " password=" + password + " host=" + host + " dbname=" + dbname
	return graph.DB, graph.DB.Open(dsn)
}
