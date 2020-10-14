package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"

	"github.com/99designs/gqlgen/handler"
	auth "github.com/Salaton/screening-test/auth"
	"github.com/Salaton/screening-test/graph"
	"github.com/Salaton/screening-test/graph/generated"
	db "github.com/Salaton/screening-test/postgres"
	"github.com/go-chi/chi"
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
		log.Fatal("Error loading .env file")
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
	user := os.Getenv("user")
	password := os.Getenv("password")
	dbname := os.Getenv("dbname")
	port := os.Getenv("port")
	sslmode := os.Getenv("sslmode")
	TimeZone := os.Getenv("TimeZone")
	dsn := "user=" + user + " password=" + password + " dbname=" + dbname + " port=" + port + " sslmode=" + sslmode + " TimeZone=" + TimeZone
	// dsn := "user=sala password=$krychowiak-254$ dbname=savannahtest port=5432 sslmode=disable TimeZone=Africa/Nairobi"
	return graph.DB, graph.DB.Open(dsn)
}
