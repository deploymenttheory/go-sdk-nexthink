package amplify

const (
	EndpointSearch              = "/apigateway/ast/assist-extension-be/api/v1/search"
	EndpointSearchDevices       = "/apigateway/ast/assist-extension-be/api/v1/search/device"
	EndpointGetDeviceProperties = "/apigateway/ast/assist-extension-be/api/v2/device/%s/properties"
	EndpointGetDeviceUsers      = "/apigateway/ast/assist-extension-be/api/v1/device/%s/users"
	EndpointGetDevicePackages   = "/apigateway/ast/assist-extension-be/api/v1/device/%s/packages"
	EndpointGetUserProperties   = "/apigateway/ast/assist-extension-be/api/v2/user/%s/properties"
	EndpointGetUserDevices      = "/apigateway/ast/assist-extension-be/api/v2/user/%s/devices"
	EndpointGetConfiguration    = "/apigateway/ast/assist-admin-be/api/v1/admin-conf"
	EndpointCreateConfiguration = EndpointGetConfiguration
	EndpointUpdateConfiguration = EndpointGetConfiguration + "/%s"
	EndpointPostInsights        = "/apigateway/ast/assist-extension-be/api/v1/insights"
)
