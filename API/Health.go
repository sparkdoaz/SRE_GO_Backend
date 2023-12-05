package main

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
)

func HealthCheckHandler(db *sql.DB, client *redis.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		dbStatus := "OK"
		redisStatus := "OK"

		if db == nil {
			dbStatus = "Not Connected"
		}

		if client == nil {
			redisStatus = "Not Connected"
		}

		status := gin.H{
			"database": dbStatus,
			"redis":    redisStatus,
		}

		if db == nil || client == nil {
			c.JSON(http.StatusServiceUnavailable, gin.H{"status": status})
			return
		}

		c.JSON(http.StatusOK, gin.H{"status": status})
	}
}
