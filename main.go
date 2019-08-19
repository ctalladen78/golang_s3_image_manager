package main

import (
  "log"
  "io/ioutil"
  "golang-projects/aws_s3_example/handlers"
  "github.com/lithammer/shortuuid"
)

func main(){
  log.Println("testing aws s3")

  a := &handler.Aws3{}
  a.Bucket = bucket
  a.Region = region
  a.Init(bucket, region)
  name := shortuuid.New()
  // read file from local dir
  image_byte, err := ioutil.ReadFile("test.png")
  if err != nil {log.Fatal(err)}
  err = a.PutImage(image_byte, name) // Aws3.PutObject(data *os.File, fileName string)
  if err != nil {
    log.Fatal("AWS S3 put error : ",err)
  }
  obj, err := a.GetImage(name)
  if err != nil {
    log.Fatal("AWS S3 get error : ",err)
  }
  log.Println("SUCCESS S3 RESULT ",obj)


}

