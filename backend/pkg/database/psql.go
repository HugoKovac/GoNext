package database

import (
	"GoNext/base/ent"
	"GoNext/base/ent/migrate"
	"GoNext/base/pkg/config"
	"context"
	"fmt"
	"log"

	"entgo.io/ent/dialect"
	_ "github.com/lib/pq"
)

// NewEntClient creates a new Ent client connected to PostgreSQL
func NewEntClient(config config.DbConfig) *ent.Client {
    dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
        config.Host, config.Port, config.User, config.Password, config.DbName)
    
    client, err := ent.Open(dialect.Postgres, dsn)
    if err != nil {
        log.Fatalf("failed opening connection to postgres: %v", err)
    }
    
    // Run the migrations
    ctx := context.Background()
    if err := client.Schema.Create(
        ctx,
        migrate.WithDropIndex(true),
        migrate.WithDropColumn(true),
    ); err != nil {
        log.Fatalf("failed creating schema resources: %v", err)
    }
    
    return client
}