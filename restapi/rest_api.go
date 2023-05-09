package restapi

type ApiConfig struct {
	ApiKey    string
	SecretKey string
	Password  string
	Simulate  bool // 模拟盘标识
	Proxy     string
}

func NewApiConfig(apiKey, secretKey, password, proxy string, simulate bool) *ApiConfig {
	return &ApiConfig{ApiKey: apiKey, SecretKey: secretKey, Password: password, Simulate: simulate, Proxy: proxy}
}
