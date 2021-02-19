package main

import (
	"log"

	data "github.com/akashkumar8/micproject/protos"
	"github.com/gin-gonic/gin"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func main() {
	r := gin.Default()
	r.GET("/", course)
	r.Run()
}

func course(g *gin.Context) {
	var conn *grpc.ClientConn

	conn, err := grpc.Dial(":8000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect:", err)
	}
	defer conn.Close()

	key := g.Query("key")
	c := data.NewCourseClient(conn)

	message := data.Request{
		Key: key,
	}

	response, err := c.Getcourse(context.Background(), &message)
	if err != nil {
		log.Fatal("error when calling say hello", err)

	}

	g.JSON(200, gin.H{
		"Response": response,
	})

	log.Printf("The '%s' course is '%s' and its traffic is '%d' in the '%sth hour' and it called overall times(calulated 1 for each hour) ", message.Key, response.Value, response.Count, response.Hour)

}
