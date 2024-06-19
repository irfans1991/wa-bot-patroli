package domain

type Key struct {
	Type                       string `json:"type"`
	Project_id                 string `json:"project_id"`
	PrivateKeyId               string `json:"private_key_id"`
	PrivateKey                 string `json:"private_key"`
	ClientEmail                string `json:"client_email"`
	ClientId                   string `json:"client_id"`
	AuthUri                    string `json:"auth_uri"`
	TokenUri                   string `json:"token_uri"`
	AuthProvider_x509_cert_url string `json:"auth_provider_x509_cert_url"`
	Client_x509_cert_url       string `json:"client_x509_cert_url"`
}
	