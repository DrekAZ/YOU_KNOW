package main
import (
  "fmt"
  "log"
  "context"
  "os"
  "io"
  "io/ioutil"
  "time"
  "strings"
  "flag"
  "encoding/json"

  "github.com/gin-gonic/gin"
  "cloud.google.com/go/firestore"
  "cloud.google.com/go/storage"
  "google.golang.org/api/iterator"

  "YOU_KNOW/Go/storage"
)

func gcpInit_Push() {
  // Sets your Google Cloud Platform project ID.
  projectID := "you-know-275301"
  // Get a Firestore client.
  ctx := context.Background()
  client, err := firestore.NewClient(ctx, projectID)
  if err != nil {
    log.Fatalf("Failed to create client: %v", err)
  }

  // Close client when done.
  defer client.Close()
  _, _, err = client.Collection("users").Add(ctx, map[string]interface{}{
    "first": "Ada",
    "last":  "Lovelace",
    "born":  1815,
  })
  if err != nil {
    log.Fatalf("Failed adding alovelace: %v", err)
  }


  iter := client.Collection("users").Documents(ctx)
  for {
    doc, err := iter.Next()
    if err == iterator.Done {
      break
    }
    if err != nil {
      log.Fatalf("Failed to iterate: %v", err)
    }
    fmt.Println(doc.Data())
  }
}

func main() {
  router := gin.Default()
  router.GET("/", func(ctx *gin.Context) {
    ctx.(200, "HELLO")
  })
  router.Run(":8080")
  fmt.Printf("DONE \n")
}