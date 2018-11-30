package collectors

// HTTP is the object that can read from an http data
type HTTP struct {
	URL                string `json:"url"`
	AuthorizationToken string `json:"token"`
}

// ListFiles returns the list of files in http endpoint
func (h *HTTP) ListFiles() ([]string, error) {
	return nil, nil
}
