package ossc

import (
	"context"
	"fmt"
	"github.com/ihezebin/sdk/logger"
	"github.com/ihezebin/sdk/ossc/qiniu"
	"github.com/ihezebin/sdk/ossc/tencent"
	"github.com/ihezebin/sdk/ossc/ucloud"
	"os"
	"strings"
	"testing"
)

func TestTencent(t *testing.T) {
	ctx := context.Background()
	client := tencent.NewClientWithConfig(tencent.Config{
		SecretID:  "AKIDjLkvsc8QVLxLUP5WjhRwCDzXrxoRpzua",
		SecretKey: "UjjV2KNBl6D5j8LK5btwlMiqjOo01nKc",
		BucketURL: "http://picture.whereabouts.icu",
	})
	url, err := client.Upload(ctx, strings.NewReader("test file3"), "test3.txt")
	if err != nil {
		logger.Fatal(err)
	}
	logger.Info(url)
	err = client.Delete(ctx, "test.txt")
	if err != nil {
		logger.Fatal(err)
	}
	faileds, err := client.DeleteMulti(ctx, "test1.txt", "test2.txt", "test3.txt")
	if err != nil {
		logger.Fatal(err)
	}
	if len(faileds) > 0 {
		logger.Error(faileds)
	}
	logger.Info("upload succeed")
}

func TestQiniu(t *testing.T) {
	client := qiniu.NewClientWithConfig(qiniu.Config{
		Zone:      qiniu.ZoneHuanan,
		AccessKey: "AccessKey",
		SecretKey: "SecretKey",
		Bucket:    "c4lms",
		Domain:    "http://image-c4lms-qiniu.whereabouts.icu",
	})
	file, err := os.Open("C:\\Users\\Korbin\\Pictures\\hzb.jpg")
	if err != nil {
		logger.Println(err)
		return
	}
	url, err := client.Upload(file, "Korbin.jpg")
	if err != nil {
		logger.Println(err, url)
		return
	}
	fmt.Println(url)
}

func TestUCloud(t *testing.T) {
	client := ucloud.NewClientWithConfig(ucloud.Config{
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
		logger.Println(err, url)
		return
	}
	fmt.Println(url)
}
