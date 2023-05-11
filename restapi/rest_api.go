package restapi

type ApiConfig struct {
	ApiKey    string
	SecretKey string
	Password  string
	Simulate  bool // 模拟盘标识
	Proxy     string
	Timeout   int
}

func NewApiConfig(apiKey, secretKey, password, proxy string, timeout int, simulate bool) *ApiConfig {
	return &ApiConfig{ApiKey: apiKey, SecretKey: secretKey, Password: password, Simulate: simulate, Proxy: proxy, Timeout: timeout}
}
