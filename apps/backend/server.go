package main

import (
	"log"
	"net/http"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/SomaTakata/real-habit/apps/backend/graph"
)

const defaultPort = "8080"

func main() {
    port := os.Getenv("PORT")
    if port == "" {
        port = defaultPort
    }

    dsn := "host=" + os.Getenv("POSTGRES_HOST") +
        " user=" + os.Getenv("POSTGRES_USER") +
        " password=" + os.Getenv("POSTGRES_PASSWORD") +
        " dbname=" + os.Getenv("POSTGRES_DB") +
        " port=" + os.Getenv("POSTGRES_PORT") +
        " sslmode=disable"
    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatalf("Failed to connect to database: %v", err)
    }

    // ここでデータベースのマイグレーションや初期化を行う
    // db.AutoMigrate(&YourModel{})

    srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{
        DB: db, // Resolverにデータベース接続を渡す
    }}))

    http.Handle("/", playground.Handler("GraphQL playground", "/query"))
    http.Handle("/query", srv)

    log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
    log.Fatal(http.ListenAndServe(":"+port, nil))
}
