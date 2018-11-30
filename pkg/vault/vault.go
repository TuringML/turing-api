package vault

import (
	"github.com/hashicorp/vault/api"
)

const (
	secretPath = "secret/data/"
)

// Vault is a simple wrapper object to deal with the Vault APIs
type Vault struct {
	client *api.Client
}

// New creates a new Vault object. DO NOT USE the ROOT token in Production
func New(token, vaultAddress string) *Vault {
	client, err := api.NewClient(&api.Config{
		Address: vaultAddress,
	})
	if err != nil {
		panic(err)
	}

	// set token for current user
	client.SetToken(token)

	return &Vault{client: client}
}

// ReadSecret returns the data from the /secret/data/:secret path in Vault
func (v *Vault) ReadSecret(secret string) (map[string]interface{}, error) {
	s, err := v.client.Logical().Read(secretPath + secret)
	if err != nil {
		return nil, err
	}
	return s.Data, nil
}

// WriteSecret store the data into the /secret/data/:secret path in Vault
func (v *Vault) WriteSecret(secret string, data map[string]interface{}) error {
	_, err := v.client.Logical().Write(secretPath+secret, data)
	if err != nil {
		return err
	}
	return nil
}
