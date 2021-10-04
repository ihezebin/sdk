package tencent

import (
	"fmt"
	"github.com/pkg/errors"

	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	tencentErrors "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	sms "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/sms/v20210111"
)

type Client interface {
	SendSms(msg *Message, telephones ...string) ([]FailedInfo, error)
}

type client struct {
	kernel *sms.Client
	config Config
}

// NewClient 实例化一个认证客户端，入参需要传入腾讯云账户密钥对secretId，secretKey
func NewClient(config Config) (Client, error) {
	if config.SecretKey == "" || config.SecretId == "" {
		return nil, errors.New("create sms client must need secret_id and secret_key")
	}
	if config.Region == "" {
		config.Region = "ap-guangzhou"
	}

	credential := common.NewCredential(config.SecretId, config.SecretKey)
	/* 非必要步骤:
	 * 实例化一个客户端配置对象，可以指定超时时间等配置 */
	cpf := profile.NewClientProfile()
	if config.HttpProfile != nil {
		if config.HttpProfile.ReqMethod != "" {
			cpf.HttpProfile.ReqMethod = config.HttpProfile.ReqMethod
		}

		if config.HttpProfile.ReqTimeout > 0 {
			cpf.HttpProfile.ReqTimeout = config.HttpProfile.ReqTimeout
		}

		if config.HttpProfile.Endpoint != "" {
			cpf.HttpProfile.Endpoint = config.HttpProfile.Endpoint
		}
	}
	if config.SignMethod != "" {
		cpf.SignMethod = config.SignMethod
	}

	/* 实例化要请求产品(以sms为例)的client对象 */
	c, err := sms.NewClient(credential, config.Region, cpf)
	if err != nil {
		return nil, err
	}

	return &client{kernel: c, config: config}, nil
}

// FailedInfo 发送失败的短信信息
type FailedInfo struct {
	Telephone string `json:"telephone"`
	Message   string `json:"message"`
}

// SendSms 发送短信
// msg: 短信消息体, 包含短信签名和模板等信息, 使用 NewMessage() 生成
// telephones: 下发手机号码，采用 E.164 标准，+[国家或地区码][手机号]
//	 * 示例如：+8613711112222， 其中前面有一个+号 ，86为国家码，13711112222为手机号，最多不要超过200个手机号
func (c *client) SendSms(msg *Message, telephones ...string) ([]FailedInfo, error) {
	/* 实例化一个请求对象，根据调用的接口和实际情况，可以进一步设置请求参数
	 * 你可以直接查询SDK源码确定接口有哪些属性可以设置
	 * 属性可能是基本类型，也可能引用了另一个数据结构
	 * 推荐使用IDE进行开发，可以方便的跳转查阅各个接口和数据结构的文档说明 */
	request := sms.NewSendSmsRequest()

	if msg != nil {
		request.SmsSdkAppId = common.StringPtr(msg.SmsSdkAppId)
		request.SignName = common.StringPtr(msg.SignName)
		request.SenderId = common.StringPtr(msg.SenderId)
		request.SessionContext = common.StringPtr(msg.SessionContext)
		request.ExtendCode = common.StringPtr(msg.ExtendCode)
		request.TemplateId = common.StringPtr(msg.TemplateId)
		request.TemplateParamSet = common.StringPtrs(msg.TemplateParamSet)
	}

	request.PhoneNumberSet = common.StringPtrs(telephones)

	// 调用接口，传入请求对象
	response, err := c.kernel.SendSms(request)
	// 处理异常
	if _, ok := err.(*tencentErrors.TencentCloudSDKError); ok {
		return nil, errors.Wrap(err, "an tencent sms api error has returned: %s")
	}
	// 非SDK异常
	if err != nil {
		return nil, err
	}
	failedInfos := make([]FailedInfo, 0)
	if response != nil && response.Response != nil && response.Response.SendStatusSet != nil {
		for _, status := range response.Response.SendStatusSet {
			if status.Code != nil && *status.Code != "Ok" {
				phone := ""
				message := ""
				if status.PhoneNumber != nil {
					phone = *status.PhoneNumber
				}
				if status.Message != nil {
					message = *status.Message
				}
				failedInfos = append(failedInfos, FailedInfo{Telephone: phone, Message: message})
			}
		}
	}
	if len(failedInfos) > 0 {
		return failedInfos, errors.New(fmt.Sprintf("send sms to some phones failed. deatils in failedInfos"))
	}

	return nil, nil
}
