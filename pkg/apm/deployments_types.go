package apm

// Deployment represents information about a New Relic application deployment.
type Deployment struct {
	Links       *DeploymentLinks `json:"links,omitempty"`
	ID          int              `json:"id,omitempty"`
	Revision    string           `json:"revision"`
	Changelog   string           `json:"changelog,omitempty"`
	Description string           `json:"description,omitempty"`
	User        string           `json:"user,omitempty"`
	Timestamp   string           `json:"timestamp,omitempty"`
}

// DeploymentLinks contain the application ID for the deployment.
type DeploymentLinks struct {
	ApplicationID int `json:"application,omitempty"`
}
