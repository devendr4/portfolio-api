package main

import (
	"net/http"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"

	// "github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	_ "github.com/joho/godotenv/autoload"
)

type TableBasics struct {
	DynamoDbClient *dynamodb.Client
	TableName      string
}

func main() {
	router := chi.NewRouter()
	dynamoClient := getClient()
	// accessKey := os.Getenv("ACCESS_KEY")
	// secret := os.Getenv("SECRET_ACCESS_KEY")
	// sess, err := session.NewSession(&aws.Config{Credentials: &creds, Region: &region})
	// fmt.Println(accessKey, sess, err)

	router.Use(render.SetContentType(render.ContentTypeJSON))
	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		projects := GetProjects(dynamoClient)
		render.JSON(w, r, projects)
	})

	http.ListenAndServe(":8000", router)
}
