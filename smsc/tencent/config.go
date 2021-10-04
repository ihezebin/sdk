package tencent

type Config struct {
	SecretId  string `json:"secret_id"`
	SecretKey string `json:"secret_key"`
	// Region 地域信息，可以直接填写字符串ap-guangzhou，或者引用预设的常量
	Region string `json:"region"`
	// HttpProfile 客户端配置对象，可以指定超时时间等配置
	HttpProfile *HttpProfile `json:"http_profile"`
	/* SignMethod SDK默认用TC3-HMAC-SHA256进行签名，非必要请不要修改这个字段 */
	SignMethod string `json:"sign_method"`
}

type HttpProfile struct {
	/*
	   ReqMethod
	   SDK默认使用POST方法。
	 * 如果你一定要使用GET方法，可以在这里设置。GET方法无法处理一些较大的请求 */
	ReqMethod string `json:"req_method"`
	/*
	   ReqTimeout
	   SDK有默认的超时时间，非必要请不要进行调整
	 * 如有需要请在代码中查阅以获取最新的默认值 */
	ReqTimeout int `json:"req_timeout"`
	/*
	   Endpoint
	   SDK会自动指定域名。通常是不需要特地指定域名的，但是如果你访问的是金融区的服务
	 * 则必须手动指定域名，例如sms的上海金融区域名： sms.ap-shanghai-fsi.tencentcloudapi.com */
	Endpoint string `json:"endpoint"`
}
