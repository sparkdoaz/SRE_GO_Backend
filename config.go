// config.go
package main

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"github.com/redis/go-redis/v9"
)

func DBAndRedisInit() {
	db = initPostgres()
	client = initRedis()
}

func initRedis() *redis.Client {

	redisHost := os.Getenv("REDIS_HOST")
	if redisHost == "" {
		sugarLogger.Info("REDIS_HOST environment variable is not set.")
		return nil
	} else {
		sugarLogger.Info("REDIS_HOST: ", redisHost)
	}

	client := redis.NewClient(&redis.Options{
		Addr: redisHost + ":6379",
	})

	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		sugarLogger.Info("Failed to connect to Redis: %v\n", err)
		return nil
	}

	return client
}

func initPostgres() *sql.DB {

	postgresHost := os.Getenv("POSTGRES_HOST")
	if postgresHost == "" {
		sugarLogger.Info("POSTGRES_HOST environment variable is not set.")
		return nil
	} else {
		sugarLogger.Info("POSTGRES_HOST: ", postgresHost)
	}

	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	dbname := os.Getenv("POSTGRES_DB")

	//sslmode=verify-full
	connStr := fmt.Sprintf("user=%s dbname=%s password=%s host=%s port=5432 sslmode=disable", user, password, dbname, postgresHost)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		sugarLogger.Info("Failed to open PostgreSQL connection: %v\n", err)
		return nil
	}

	return db
}

func HealthCheckHandler(db *sql.DB, client *redis.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		status := "All OK"
		if db == nil && client == nil {
			status = "Neither DB nor Redis is connected"
		} else if db == nil {
			status = "DB is not connected"
		} else if client == nil {
			status = "Redis is not connected"
		}

		postgresHost := os.Getenv("POSTGRES_HOST")
		redisHost := os.Getenv("REDIS_HOST")
		user := os.Getenv("POSTGRES_USER")
		password := os.Getenv("POSTGRES_PASSWORD")
		dbname := os.Getenv("POSTGRES_DB")

		c.JSON(http.StatusOK, gin.H{
			"status":   status,
			"postgres": postgresHost,
			"redis":    redisHost,
			"user":     user,
			"password": password,
			"dbname":   dbname,
		})
	}
}
