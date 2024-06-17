package repository

import (
	"context"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var rdb *redis.Client
var mongoCLient *mongo.Client

func ConnectToRedis() {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "192.168.50.131:6379", // Thay thế bằng địa chỉ Redis thực tế
		Password: "",                    // Mật khẩu (nếu có)
		DB:       0,                     // Chọn cơ sở dữ liệu
	})
	ping, err := rdb.Ping(context.Background()).Result()
	if err != nil {
		fmt.Printf("Ping err: %v\n", err)
	}
	fmt.Printf("ping: %v\n", ping)
}

func GetRedisClient() *redis.Client {
	return rdb
}

func GetMongoClient() *mongo.Client {
	return mongoCLient
}

func ConnectToMongoDB() {
	clientOptions := options.Client().ApplyURI("mongodb://192.168.50.131:27017")
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")
	mongoCLient = client
}

// Ta muốn sinh ra 1 mảng bao gồm tất cả các ký tự có thể kết hợp
func Logger(c *gin.Context) {
	log := logrus.New()

	// Logging thông tin về yêu cầu
	log.WithFields(logrus.Fields{
		"clientIP": c.ClientIP(),
		"method":   c.Request.Method,
		"path":     c.FullPath(),
	}).Info("Request")

	// Thực hiện xử lý

	// Logging thông tin về phản hồi
	log.WithFields(logrus.Fields{
		"clientIP": c.ClientIP(),
		"status":   c.Writer.Status(),
	}).Info("Response")
}
