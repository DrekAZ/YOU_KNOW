package main
import (
  "fmt"
  "log"
  "context"
  "os"
  //"io"
  "io/ioutil"
  "time"
  "unsafe"
  //"bufio"
  //"strings"
  //"flag"
  "github.com/gin-gonic/gin"
  "cloud.google.com/go/storage"
  //"golang.org/x/oauth2/google"
  "google.golang.org/api/option"
  //"google.golang.org/api/iterator"
)

type JsonData struct {
  Name  string `json:"name"`
  Markdown string `json:"markdown"`
}

func main() {
  path := "./YOU-KNOW-be4a1d88e2c3.json"

  ctx := context.Background()

  client, err := storage.NewClient(ctx, option.WithCredentialsFile(path))
  if err != nil {
    log.Fatalf("Failed to create client: %v", err)
  }

  // Sets the name for the new bucket.
  bucket := "you-know-275301.appspot.com"
  // Creates a Bucket instance.
  ///// bucket := client.Bucket(bucketName)

  // Creates the new bucket.
  ctx, cancel := context.WithTimeout(ctx, time.Second*1000)
  defer cancel()

  

  router := gin.Default()
  router.GET("/get", func(c *gin.Context) {
    name := c.Query("name")
    if name == "" {
      log.Fatalf("No query")
    }
    
    data, err := read(client, bucket, name)
    if err != nil {
      log.Fatalf("Cannot", err)
    }

    str := Byte2str(data)

    c.JSON(200, gin.H {
      "content": str,
    })
  })

  router.POST("/update", func(c *gin.Context) {
    var json_data JsonData
    if err := c.BindJSON(&json_data); err != nil {
      log.Fatal(err)
    }
    //fmt.Printf("--> ", json_data.Name, json_data.Markdown)
    if json_data.Name == "" {
      log.Fatalf("No query")
    }

    err := write(client, bucket, json_data.Name, json_data.Markdown)
    if err != nil {
      log.Fatalf("Cannot", err)
    }

    c.JSON(200, gin.H {
      "OK": true,
    })
  })
  fmt.Printf("DONE \n")

  router.Run()
}
func write(client *storage.Client, bucket, name string, markdown string) error {
  ctx := context.Background()
  f, err := os.Open(name)
  if err != nil {
    return err
  }
  defer f.Close()

  ctx, cancel := context.WithTimeout(ctx, time.Second*5000)
  defer cancel()
  wc := client.Bucket(bucket).Object(name).NewWriter(ctx)
  //if _, err = io.Copy(wc, markdown); err != nil {
  if _, err = fmt.Fprintf(wc, markdown); err != nil {
    return err
  }
  if err := wc.Close(); err != nil {
    return err
  }

  return nil
}
func read(client *storage.Client, bucket, name string) ([]byte, error) {
  ctx := context.Background()

  ctx, cancel := context.WithTimeout(ctx, time.Second*5000)
  defer cancel()

  rc, err := client.Bucket(bucket).Object(name).NewReader(ctx)
  if err != nil {
    return nil, err
  }
  defer rc.Close()

  data, err := ioutil.ReadAll(rc)
  if err != nil {
    return nil, err
  }

  /*f, err := os.Create(object)
  tee := io.TeeReader(rc, f)
  s := bufio.NewScanner(tee)
  for s.Scan() {
  }
  if err := s.Err(); err != nil {
    log.Fatalf("NO FILE ERROR", err)
  }*/


  return data, nil
}

func Byte2str(b []byte) (string) {
  return *(*string)(unsafe.Pointer(&b))
}