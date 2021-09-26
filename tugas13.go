package main

import "fmt"
import "encoding/base64"
import "crypto/sha1"

func main(){
  myname:= "I Made Ramdhawa Yudapatha"
  encode:= base64.StdEncoding.EncodeToString([]byte(myname))
  fmt.Println("encode dengan base64 :",encode)
  sha := sha1.New()
  sha.Write([]byte(encode))
  enkripsi := sha.Sum(nil)
  hasil := fmt.Sprintf("%x",enkripsi)
  fmt.Println("Enkripsi dengan sha1 :",hasil)
}
