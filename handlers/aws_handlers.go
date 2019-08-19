package handler

import (
  "log"
//   "os"
//   "strconv"
//   "time"
  "bytes"
//   "image/png"
	"net/http"
  "github.com/aws/aws-sdk-go/aws"
  "github.com/aws/aws-sdk-go/aws/session"
  "github.com/aws/aws-sdk-go/service/s3"
  "github.com/aws/aws-sdk-go/aws/credentials"
)

var(
  ACCESSKEY = ""
  SECRETKEY = ""
  //creds = credentials.NewStaticCredentials(ACCESSKEY, SECRETKEY)
  bucket = "tassel-app-images"
  region = "us-east-1"
  svc = s3.New(session.Must(session.NewSession(&aws.Config{
    Region: aws.String(region),
    Credentials: credentials.AnonymousCredentials,
  }))) 
)

type Aws3 struct {
	Sess 	*session.Session
	Bucket	string
	Region	string
}

func (a *Aws3) Init(bucket string, region string) {
	a.Sess = session.Must(session.NewSession(&aws.Config{
		Region: aws.String(region),
		Credentials: credentials.AnonymousCredentials,
	})) 
}

func (a *Aws3) GetImage(id string) (*s3.GetObjectOutput, error){
	log.Print("GET OBJECT ",id)
	svc := s3.New(a.Sess)
	input := &s3.GetObjectInput{
		Bucket: aws.String(a.Bucket),
		Key:    aws.String(id+".png"),
	}

	result, err := svc.GetObject(input)
	if err != nil {
			return nil,err
	}
	return result, nil
}


// https://github.com/kuzaxak/promalert/blob/master/s3.go
// https://docs.aws.amazon.com/sdk-for-go/api/service/s3/#PutObjectInput
// https://echo.labstack.com/cookbook/file-upload
// https://stackoverflow.com/questions/20602131/io-writeseeker-and-io-readseeker-from-byte-or-file
func (a *Aws3) PutImage(image []byte, key string) error {
	svc := s3.New(a.Sess)
	key = key + ".png"
	send_s3 := bytes.NewReader(image)

	_, err := svc.PutObject(&s3.PutObjectInput{
		Bucket:        aws.String(a.Bucket),
		Key:           aws.String(key),
		ACL:           aws.String("public-read"),
		Body:          send_s3,
	// 	ContentLength: aws.Int64(int64(size)),
		ContentType:   aws.String("image/png"),
  })
  if err != nil { return err} 

  // return link to new file
  // save to db (userName-tempFileName)
  log.Printf("https://s3.amazonaws.com/%s/%s", bucket, key,)
  return nil
}