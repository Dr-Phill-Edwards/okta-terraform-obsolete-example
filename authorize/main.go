package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"cloud.google.com/go/firestore"
	"github.com/gin-gonic/gin"
	"google.golang.org/api/iterator"
)

var client *firestore.Client

func GetUsers(c *gin.Context) {
	var result []map[string]interface{}
	iter := client.Collection("user-roles").Documents(context.Background())
	for {
		doc, err := iter.Next()
		fmt.Println(err)
		if err == iterator.Done {
			break
		}
		result = append(result, doc.Data())
	}
	c.JSON(http.StatusOK, result)
}

func PostUser(c *gin.Context) {
	status := http.StatusCreated
	email := c.PostForm("email")
	response := email
	role := c.PostForm("role")
	_, _, err := client.Collection("user-roles").Add(context.Background(), map[string]interface{}{"email": email, "role": role})
	if err != nil {
		status = http.StatusBadRequest
		response = err.Error()
	}
	c.String(status, response)
}

func main() {
	var err error
	fmt.Println("ID " + os.Getenv("GOOGLE_PROJECT_ID"))
	client, err = firestore.NewClient(context.Background(), os.Getenv("GOOGLE_PROJECT_ID"))
	if err != nil {
		fmt.Println("Error " + err.Error())
	}
	r := gin.Default()
	r.GET("/api/authorize", GetUsers)
	r.POST("/api/authorize", PostUser)
	r.Run(":8080")
}
