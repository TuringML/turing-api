package collectors

// AzureDataLake is the object that can be used to get/store files
type AzureDataLake struct {
	Tenant       string `json:"tenant"`
	ClientID     string `json:"clientId"`
	ClientSecret string `json:"clientSecret"`
}

// ListFiles returns the list of files in Azure data lake bucket
func (a *AzureDataLake) ListFiles() ([]string, error) {
	return nil, nil
}
