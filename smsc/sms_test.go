package smsc

import (
	"github.com/ihezebin/sdk/logger"
	"github.com/ihezebin/sdk/smsc/aliyun"
	"github.com/ihezebin/sdk/smsc/tencent"
	"github.com/ihezebin/sdk/utils/common"
	"testing"
)

func TestTencentSms(t *testing.T) {
	client, err := tencent.NewClientWithConfig(tencent.Config{
		SecretId:  "AKIDjLkv*******zXrxoRpzua",
		SecretKey: "UjjV2KNBl6D*********qjOo01nKc",
		Region:    "ap-guangzhou",
	})
	if err != nil {
		logger.Fatal(err)
	}
	msg := tencent.NewMessage().WithAppId("1400578890").WithSignName("whereabouts").
		WithTemplate("11477481", 123321, 10)
	faileds, err := client.SendSms(msg, "+8613518468111")
	if err != nil {
		logger.Error(faileds)
		logger.Fatal(err)
	}
	logger.Info("send sms succeed")
}

func TestAliyunSms(t *testing.T) {
	client, err := aliyun.NewClientWithConfig(aliyun.Config{
		AccessKeyId:     "",
		AccessKeySecret: "",
	})
	if err != nil {
		logger.Fatal(err)
	}
	msg := aliyun.NewMessage().WithSignName("sign").WithTemplate("code", common.Json{})
	err = client.SendSms(msg, "13518468111")
	if err != nil {
		logger.Fatal(err)
	}
	logger.Info("send sms succeed")
}
