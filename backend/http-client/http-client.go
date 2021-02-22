package httpclient

// HTTPClient ...
type HTTPClient interface {
	Get(url string) ([]byte, error)
}
