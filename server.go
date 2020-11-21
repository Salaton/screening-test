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

	// srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}

// InitDB function to start the db connections process
func InitDB() (db.DBClient, error) {
	graph.DB = &db.PostgresClient{}
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	// port := os.Getenv("DB_PORT")
	dbname := os.Getenv("DB_NAME")
	host := os.Getenv("DB_HOST")
	// sslmode := os.Getenv("SSLMODE")
	// TimeZone := os.Getenv("TimeZone")
	dsn := "user=" + user + " password=" + password + " host=" + host + " dbname=" + dbname
	// dsn := "user=sala password=$krychowiak-254$ dbname=savannahtest port=5432 sslmode=disable TimeZone=Africa/Nairobi"
	return graph.DB, graph.DB.Open(dsn)
}
