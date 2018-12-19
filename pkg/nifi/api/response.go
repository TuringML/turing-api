package api

type Response struct {
	Permissions struct {
		CanRead  bool `json:"canRead"`
		CanWrite bool `json:"canWrite"`
	} `json:"permissions"`
	ProcessGroupFlow struct {
		ID         string `json:"id"`
		URI        string `json:"uri"`
		Breadcrumb struct {
			ID          string `json:"id"`
			Permissions struct {
				CanRead  bool `json:"canRead"`
				CanWrite bool `json:"canWrite"`
			} `json:"permissions"`
			Breadcrumb struct {
				ID   string `json:"id"`
				Name string `json:"name"`
			} `json:"breadcrumb"`
		} `json:"breadcrumb"`
		Flow struct {
			ProcessGroups       []interface{} `json:"processGroups"`
			RemoteProcessGroups []interface{} `json:"remoteProcessGroups"`
			Processors          []interface{} `json:"processors"`
			InputPorts          []interface{} `json:"inputPorts"`
			OutputPorts         []interface{} `json:"outputPorts"`
			Connections         []interface{} `json:"connections"`
			Labels              []interface{} `json:"labels"`
			Funnels             []interface{} `json:"funnels"`
		} `json:"flow"`
		LastRefreshed string `json:"lastRefreshed"`
	} `json:"processGroupFlow"`
}
