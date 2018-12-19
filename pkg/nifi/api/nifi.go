package api

// NiFi base struct
// DisconnectedNodeAcknowledged default is false
type NiFi struct {
	Revision                     Revision  `json:"revision"`
	DisconnectedNodeAcknowledged bool      `json:"disconnectedNodeAcknowledged"`
	Component                    Component `json:"component"`
}

// Revision information of Version and ClientID
// ClientID is based per each client
// Version is by default 0
type Revision struct {
	ClientID string `json:"clientId"`
	Version  int    `json:"version"`
}

// Component which allows to set the component
// Type e.g. org.apache.nifi.processors.standard.LogMessage
type Component struct {
	Source                        Processor     `json:"source"`
	Destination                   Processor     `json:"destination"`
	SelectedRelationships         []string      `json:"selectedRelationships"`
	FlowFileExpiration            string        `json:"flowFileExpiration"`
	BackPressureDataSizeThreshold string        `json:"backPressureDataSizeThreshold"`
	BackPressureObjectThreshold   string        `json:"backPressureObjectThreshold"`
	Bends                         []interface{} `json:"bends"`
	Prioritizers                  []interface{} `json:"prioritizers"`
	LoadBalanceStrategy           string        `json:"loadBalanceStrategy"`
	LoadBalancePartitionAttribute string        `json:"loadBalancePartitionAttribute"`
	LoadBalanceCompression        string        `json:"loadBalanceCompression"`
	Type                          string        `json:"type"`
	Bundle                        Bundle        `json:"bundle"`
	Name                          string        `json:"name"`
	Position                      Position      `json:"position"`
}

// Processor represents a source or destination
type Processor struct {
	ID      string `json:"id"`
	GroupID string `json:"groupId"`
	Type    string `json:"type"`
}

// Bundle contains Group, Artifact and Version e.g.
// Group: org.apache.nifi
// Artifact: nifi-standard-nar
// Version: 1.8.0
type Bundle struct {
	Group    string `json:"group"`
	Artifact string `json:"artifact"`
	Version  string `json:"version"`
}

// Position is the place of the canvas
type Position struct {
	X int `json:"x"`
	Y int `json:"y"`
}
