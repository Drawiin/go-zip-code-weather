package client

type WebClient interface {
	Get(url string) (string, error)
}

type WebClientImpl struct{}

func (w WebClientImpl) Get(url string) (string, error) {
	panic("implement me")
}

func NewWebClient() WebClient {
	return WebClientImpl{}
}
