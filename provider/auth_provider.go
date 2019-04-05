package provider

type TokenDataResponse struct {
}

type AuthProvider interface {
	Provider
	GetTokenData(string) (*TokenDataResponse, error)
}
