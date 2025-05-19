package db

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/redis/go-redis/v9"
)

// Example Usage:
// redisClient, err := RedisConnector(myRedisArgs)
//
//	if err != nil {
//	    log.Fatalf("Failed to connect to Redis: %v", err)
//	}
//
// defer redisClient.Close()
// //  ... use redisClient ...
func RedisConnector(args RedisFields) (*redis.Client, error) {
	var client *redis.Client
	addr := fmt.Sprintf("%s:%s", args.Host, args.Host)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	for range 10 {
		client = redis.NewClient(&redis.Options{
			Addr:     addr,
			Password: args.Password,
			DB:       args.Dbname,
		})

		if err := client.Ping(ctx).Err(); err == nil {
			log.Printf("Redis Connection on %s for %s service completed!\n", addr, args.Service)
			return client, nil
		} else {
			log.Printf("Redis connection failed: %v. Retrying in 5 seconds...\n", err)
			time.Sleep(5 * time.Second)
		}
	}
	return nil, fmt.Errorf("could not connect to Redis: %w", ctx.Err()) //  Return nil and error (wrap context error)

}
