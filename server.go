package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-pg/pg/v10"
	"github.com/joho/godotenv"
	"github.com/rs/cors"
	"github.com/userq11/meetmeup/graph"
	"github.com/userq11/meetmeup/graph/domain"
	"github.com/userq11/meetmeup/graph/generated"
	customMiddleware "github.com/userq11/meetmeup/middleware"
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

	userRepo := postgres.UsersRepo{DB: DB}

	router := chi.NewRouter()

	router.Use(cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:8081"},
		AllowCredentials: true,
		Debug:            true,
	}).Handler)

	router.Use(middleware.RequestID)
	router.Use(middleware.Logger)
	router.Use(customMiddleware.AuthMiddleware(userRepo))

	d := domain.NewDomain(userRepo, postgres.MeetupsRepo{DB: DB})

	c := generated.Config{Resolvers: &graph.Resolver{Domain: d}}

	queryHandler := handler.NewDefaultServer(generated.NewExecutableSchema(c))

	router.Handle("/", playground.Handler("GraphQL playground", "/query"))
	router.Handle("/query", graph.DataloaderMiddleware(DB, queryHandler))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
