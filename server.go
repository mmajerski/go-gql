package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-pg/pg/v10"
	"github.com/joho/godotenv"
	"github.com/userq11/meetmeup/graph"
	"github.com/userq11/meetmeup/graph/generated"
	"github.com/userq11/meetmeup/postgres"
)

func goDotEnvVariable(key string) string {

	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}

const defaultPort = "8081"

func main() {
	DB := postgres.New(&pg.Options{
		User:     goDotEnvVariable("DB_USER"),
		Password: goDotEnvVariable("DB_PASSWORD"),
		Database: goDotEnvVariable("DB_DATABASE"),
	})

	defer DB.Close()

	DB.AddQueryHook(postgres.DBLogger{})

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	c := generated.Config{Resolvers: &graph.Resolver{MeetupsRepo: postgres.MeetupsRepo{DB: DB}, UsersRepo: postgres.UsersRepo{DB: DB}}}

	queryHandler := handler.NewDefaultServer(generated.NewExecutableSchema(c))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", graph.DataloaderMiddleware(DB, queryHandler))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
