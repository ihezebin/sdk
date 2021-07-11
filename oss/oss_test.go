package oss

import (
	"fmt"
	"github.com/whereabouts/sdk-go/oss/qiniu"
	"github.com/whereabouts/sdk-go/oss/ucloud"
	"log"
	"os"
	"testing"
)

func TestQiniu(t *testing.T) {
	client := NewQiNiuClient(&qiniu.Config{
		Zone:      qiniu.ZoneHuanan,
		AccessKey: "AccessKey",
		SecretKey: "SecretKey",
		Bucket:    "c4lms",
		Domain:    "http://image-c4lms-qiniu.whereabouts.icu",
	})
	file, err := os.Open("C:\\Users\\Korbin\\Pictures\\hzb.jpg")
	if err != nil {
		log.Println(err)
		return
	}
	url, err := client.Upload(file, "Korbin.jpg")
	if err != nil {
		log.Println(err, url)
		return
	}
	fmt.Println(url)
}

func TestUCloud(t *testing.T) {
	client := NewUCloudClient(&ucloud.Config{
		PublicKey:  "PublicKey",
		PrivateKey: "PrivateKey",
		FileHost:   "cn-bj.ufileos.com",
		BucketName: "c4lms",
	})
	file, err := os.Open("C:\\Users\\Korbin\\Pictures\\hzb.jpg")
	if err != nil {
		fmt.Println(err)
		return
	}
	url, err := client.Upload(file, "Korbin.jpg")
	if err != nil {
		log.Println(err, url)
		return
	}
	fmt.Println(url)
}
