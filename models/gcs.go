package models

// GCloudStore is the object that can be used to get/store files
type GCloudStore struct {
	ProjectID         string `json:"projectId"`
	ServiceAccountKey string `json:"serviceAccountKey"`
	Bucket            string `json:"bucket"`
	Prefix            string `json:"prefix"`
}

// ListFiles returns the list of files in GCloudStorage bucket
func (g *GCloudStore) ListFiles() ([]string, error) {
	return nil, nil
}
