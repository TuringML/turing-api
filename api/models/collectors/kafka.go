package collectors

// ApacheKafka is the object in charge of reading the data stream from the topic
type ApacheKafka struct {
	Topic          string `json:"topic"`
	Authentication string `json:"auth"`
	Brokers        string `json:"brokers"`
}

// ListFiles returns the list of "files" in a Kafka topic
func (a *ApacheKafka) ListFiles() ([]string, error) {
	return nil, nil
}
