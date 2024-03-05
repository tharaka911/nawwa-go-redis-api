package controllers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/tharaka911/go-redis-api/initializers"
	"github.com/tharaka911/go-redis-api/models"
	"github.com/tharaka911/go-redis-api/utils"
)

func WriteMessageToQueue(c *gin.Context) {

	//Get data off req body
	var body struct {
		Body  string
		Title string
	}
	c.Bind(&body)

	// Get the post ID from Redis
	postIDStr, err := initializers.DB.Get(ctx, "post-id").Result()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not get post ID"})
		return
	}

	post := models.PostRedis{
		Id:           postIDStr,
		Title:        body.Title,
		Body:         body.Body,
		CreationTime: time.Now().String(),
		UpdatingTime: time.Now().String(),
	}

	postMap := map[string]interface{}{
		"Title":        post.Title,
		"Body":         post.Body,
		"CreationTime": post.CreationTime,
		"UpdatingTime": post.UpdatingTime,
	}

	jsonString := fmt.Sprintf(`{"title": "%s", "body": "%s", "creation_time": "%s", "updating_time": "%s", "id": "%s"}`, postMap["Title"], postMap["Body"], postMap["CreationTime"], postMap["UpdatingTime"], post.Id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not write message to queue"})
		return
	}

	// Write message to queue in goroutine
	go utils.WriteMessageToQueue(jsonString)

	c.JSON(200, gin.H{
		"message": "message sent to queue",
	})

}
