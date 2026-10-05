package software_metering

// ConfigurationInput is the UI save payload. Thresholds apply to each selected application type.
type ConfigurationInput struct {
	Name             string                `json:"name"`
	Description      string                `json:"description"`
	NQLID            string                `json:"nqlId"`
	LicenseType      string                `json:"licenseType"`
	ApplicationUUIDs []string              `json:"applicationUuids"`
	Thresholds       []Threshold           `json:"thresholds"`
	Packages         []ApplicationPackages `json:"packages,omitempty"`
}
type Threshold struct {
	AppType    string  `json:"appType"`
	MetricType string  `json:"metricType"`
	Value      float64 `json:"value"`
	TimeUnit   string  `json:"timeUnit"`
}
type Package struct {
	Name string `json:"name"`
}
type ApplicationPackages struct {
	AppUUID  string    `json:"appUuid"`
	Packages []Package `json:"packages"`
}
type Application struct {
	UUID string `json:"uuid"`
	Name string `json:"name"`
	Type string `json:"type"`
}
type Configuration struct {
	UUID         string                `json:"uuid"`
	Name         string                `json:"name"`
	NQLID        string                `json:"nqlId"`
	LicenseType  string                `json:"licenseType"`
	Description  string                `json:"description"`
	Thresholds   []Threshold           `json:"thresholds"`
	Applications []Application         `json:"applications"`
	Packages     []ApplicationPackages `json:"packages"`
}
type Summary struct {
	UUID             string `json:"uuid"`
	Name             string `json:"name"`
	LicenseType      string `json:"licenseType"`
	ApplicationCount int    `json:"applicationCount"`
}
type ListResponse struct {
	Configurations []Summary `json:"configurations"`
}
type GetResponse struct {
	Configuration *Configuration `json:"configuration"`
}

// Create returns a boolean, not a UUID. Find the created configuration using List.
type CreateResponse struct {
	Created *bool `json:"createConfiguration"`
}
type UpdateResponse struct {
	Updated *bool `json:"updateConfiguration"`
}
type DeleteResponse struct {
	Deleted *bool `json:"deleteConfiguration"`
}
