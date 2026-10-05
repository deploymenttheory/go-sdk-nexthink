package product_shell

import "encoding/json"

// Envelope is the product shell's result/status response format.
type Envelope[T any] struct {
	Result       T            `json:"result"`
	ResultStatus ResultStatus `json:"resultStatus"`
}
type ResultStatus struct {
	Code        int    `json:"code"`
	Description string `json:"description"`
}
type (
	MenuResponse    = Envelope[Menu]
	UserResponse    = Envelope[User]
	ModulesResponse = Envelope[[]Module]
	FlagResponse    = Envelope[map[string]bool]
)

// Configuration retains nested product-specific flags and settings.
type (
	ConfigurationResponse = Envelope[json.RawMessage]
	Menu                  struct {
		Apps             []json.RawMessage `json:"apps"`
		Menus            []json.RawMessage `json:"menus"`
		MenuActions      []json.RawMessage `json:"menuActions"`
		Groups           []json.RawMessage `json:"groups"`
		Items            []json.RawMessage `json:"items"`
		ItemActions      []json.RawMessage `json:"itemActions"`
		TelemetryModules []json.RawMessage `json:"telemetryModules"`
	}
)
type User struct {
	UserID      string   `json:"userId"`
	Username    string   `json:"username"`
	FullName    string   `json:"fullName"`
	Email       string   `json:"email"`
	MTPRoles    []string `json:"mtpRoles"`
	ViewDomain  string   `json:"viewDomain"`
	UserType    string   `json:"userType"`
	Locale      string   `json:"locale"`
	LandingPage string   `json:"landingPage"`
	CanResetMFA bool     `json:"canResetMfa"`
}
type Module struct {
	ID               string            `json:"id"`
	Name             string            `json:"name"`
	Entrypoint       string            `json:"entrypoint"`
	Routes           []json.RawMessage `json:"routes"`
	Version          string            `json:"version"`
	CommandHandlers  json.RawMessage   `json:"commandHandlers"`
	EventSubscribers json.RawMessage   `json:"eventSubscribers"`
}
