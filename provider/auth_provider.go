package provider

type TokenDataResponse struct {
	IsValid          bool
	Provider         string
	ProviderUsername string
	UserName         *string
}

type AuthProvider interface {
	Provider
	GetTokenData(string) (*TokenDataResponse, error)
}
