package db

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/IBM/sarama"
	"github.com/redis/go-redis/v9"
)

type SqlFields struct {
	Username string
	Password string
	Host     string
	Port     string
	Dbname   string
	Service  string
}

type RedisFields struct {
	Host     string
	Port     string
	Password string
	Dbname   int
	Service  string
}

type KafkaFields struct {
	Brokers []string
	Group   string
	Service string
}

// Example usage:
//
// db, err := MysqlConnector(mySqlArgs)
//
// defer db.Close()
//
// ... use db ...
//
// _, err = db.Exec(createTableQuery) //  Table creation moved out
func MysqlConnector(args SqlFields) (*sql.DB, error) {
	var err error
	var db *sql.DB
	// mysql connection string
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", args.Username, args.Password, args.Host, args.Port, args.Dbname)
	for range 10 {
		db, err = sql.Open("mysql", dsn)
		if err == nil {
			err = db.Ping()
			if err == nil {
				log.Printf("MySql Connection %s for %s service completed!\n", args.Dbname, args.Service)
				return db, nil
			}
		}
		log.Printf("Database connection failed: %v. Retrying in 5 seconds...\n", err)
		time.Sleep(5 * time.Second)
	}
	return nil, fmt.Errorf("could not connect to the database: %w", err)

}

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

// Example Usage:
// consumerGroup, err := KafkaConnector(myKafkaArgs)
//
//	if err != nil {
//	    log.Fatalf("Failed to connect to Kafka: %v", err)
//	}
//
// defer consumerGroup.Close()
// //  ... use consumerGroup ...
func KafkaConnector(args KafkaFields) (sarama.ConsumerGroup, error) {
	config := sarama.NewConfig()
	config.Consumer.Group.Rebalance.Strategy = sarama.NewBalanceStrategyRoundRobin()

	var client sarama.ConsumerGroup
	var err error

	for range 10 {
		client, err = sarama.NewConsumerGroup(args.Brokers, args.Group, config)
		if err == nil {
			log.Printf("Kafka Connection to %v for %s service completed!\n", args.Brokers, args.Service)
			return client, nil
		}
		log.Printf("Kafka connection failed: %v. Retrying in 5 seconds...\n", err)
		time.Sleep(5 * time.Second)
	}
	log.Fatal("Could not connect to Kafka")
	return nil, fmt.Errorf("could not connect to Kafka") //  Return nil and error
}
