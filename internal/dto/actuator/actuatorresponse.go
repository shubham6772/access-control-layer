package actuatordto

type HealthResponse struct {
	Status string `json:"status"`
}

type ReadyResponse struct {
	Status string `json:"status"`
	Redis  string `json:"redis,omitempty"`
}

type InfoResponse struct {
	Service string `json:"service"`
	Version string `json:"version"`
}
