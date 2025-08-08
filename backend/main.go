package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"math/rand"
	"time"
)

type Random struct {
	Id        int    `json:"id"`
	Random    string `json:"random"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func main() {

	r := gin.Default()

	r.GET("/backend/actuator/health", func(c *gin.Context) {
		c.JSON(200, "backend OK")
	})

	r.GET("/backend/random", func(c *gin.Context) {
		rand.Seed(time.Now().UnixNano()) // 设置随机种子

		res := make([]*Random, 0)
		for i := 0; i < 10; i++ {
			data := &Random{
				Id:        i + 1,
				Random:    generateRandomString(20, 30),
				CreatedAt: time.Now().Format("2006-01-02 15:04:05"),
				UpdatedAt: time.Now().Format("2006-01-02 15:04:05"),
			}
			res = append(res, data)
		}
		c.JSON(200, res)
	})

	log.Fatal(r.Run(":8080"))
}

func generateRandomString(minLen, maxLen int) string {
	letters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	length := rand.Intn(maxLen-minLen+1) + minLen
	result := make([]rune, length)
	for i := range result {
		result[i] = letters[rand.Intn(len(letters))]
	}
	return string(result)
}
