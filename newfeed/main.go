package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jamessee/gin.git/newfeed/httpd/handler"
	"github.com/jamessee/gin.git/newfeed/httpd/platform/newsfeed"
)

func main() {
	feed := newsfeed.New()

	r := gin.Default()
	r.GET("/ping", handler.PingGet())
	r.GET("/newsfeed", handler.NewsFeedGet(feed))
	r.POST("/newsfeed", handler.NewsFeedPost(feed))
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

	// fmt.Println(feed)
	// feed.Add(newsfeed.Item{"hello", "World"})
	// fmt.Println(feed)

}
