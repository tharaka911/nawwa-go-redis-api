package controllers

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/tharaka911/go-redis-api/initializers"
	"github.com/tharaka911/go-redis-api/models"
)

var ctx = context.Background()

// create post
func PostCreate(c *gin.Context) {

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

	result := initializers.DB.HMSet(ctx, "posts:"+post.Id, postMap)

	if result.Err() != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "post creating have error"})
		return
	} else {
		// Increment the post ID in Redis
		_, err := initializers.DB.Incr(ctx, "post-id").Result()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "could not increment post ID"})
			return
		}
	}
	//return post
	printForTest(post)

	c.JSON(200, gin.H{
		"message": post,
	})
}

// get all posts
func PostGetAll(c *gin.Context) {
	// Get all post keys
	keys, err := initializers.DB.Keys(ctx, "posts:*").Result()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not get posts"})
		return
	}

	// Create posts list to store posts
	var posts []models.PostRedis

	// Get all posts
	for _, key := range keys {
		postMap, err := initializers.DB.HGetAll(ctx, key).Result()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "could not get post"})
			return
		}

		post := models.PostRedis{
			Id:           strings.TrimPrefix(key, "posts:"),
			Title:        postMap["Title"],
			Body:         postMap["Body"],
			CreationTime: postMap["CreationTime"],
			UpdatingTime: postMap["UpdatingTime"],
		}

		posts = append(posts, post)
	}

	// Return posts
	c.JSON(200, gin.H{
		"message": posts,
	})
}

// get a post
func PostGet(c *gin.Context) {
	// Get ID from URL
	id := c.Param("id")

	// Get post
	postMap, err := initializers.DB.HGetAll(ctx, "posts:"+id).Result()
	if err != nil || len(postMap) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "post not found"})
		return
	}

	post := models.PostRedis{
		Id:           id,
		Title:        postMap["Title"],
		Body:         postMap["Body"],
		CreationTime: postMap["CreationTime"],
		UpdatingTime: postMap["UpdatingTime"],
	}

	// Return post
	c.JSON(200, gin.H{
		"message": post,
	})
}

// update a post
func PostUpdate(c *gin.Context) {
	// Get ID from URL
	id := c.Param("id")

	// Get data off req body
	var body struct {
		Body  string
		Title string
	}
	c.Bind(&body)

	// Check if post exists
	exists, err := initializers.DB.Exists(ctx, "posts:"+id).Result()
	if err != nil || exists == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "post not found"})
		return
	}

	// Update post
	_, err = initializers.DB.HSet(ctx, "posts:"+id, map[string]interface{}{
		"Title":        body.Title,
		"Body":         body.Body,
		"UpdatingTime": time.Now().String(),
	}).Result()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not update post"})
		return
	}

	// Return success message
	c.JSON(200, gin.H{
		"message": "post updated successfully",
	})
}

// delete a post
func PostDelete(c *gin.Context) {
	// Get ID from URL
	id := c.Param("id")

	// Delete post
	result, err := initializers.DB.Del(ctx, "posts:"+id).Result()
	if err != nil || result == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "post not found"})
		return
	}

	// Return success message
	c.JSON(200, gin.H{
		"message": "post deleted successfully",
	})
}

func printForTest(post models.PostRedis) {
	fmt.Println(post.Title)
	fmt.Println(post.Body)
}
