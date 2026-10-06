package benchmark

import "encoding/json"

type ProfileType struct {
	Type        string `json:"type"`
	DataVersion string `json:"dataVersion"`
}
type ProfilePropertiesParams struct {
	ProfileType
	BinaryName  string `json:"binaryName"`
	ProductName string `json:"productName"`
}
type ProductPropertiesParams struct {
	ProfileType
	ProductName string `json:"productName"`
}
type ProfileVersionsParams struct {
	ProfilePropertiesParams
	RequestedAttributes []string `json:"requestedAttributes"`
	RequestedMetrics    []string `json:"requestedMetrics"`
}
type ProfileParams struct {
	ProfileVersionsParams
	Version string `json:"version"`
}

// Benchmark filters depend on the selected metric and retain their native JSON shape.
type BenchmarkParams struct {
	Metric  string          `json:"metric"`
	Filters json.RawMessage `json:"filters"`
}

type GetBinaryProductMapsRequest struct {
	Params *ProfileType `json:"params" required:"true"`
}

type GetBinaryProductMapsResponseBinaryProductMapsBinaryToProductsMap struct {
	Key    *string         `json:"key"`
	Values json.RawMessage `json:"values"`
}

type GetBinaryProductMapsResponseBinaryProductMapsProductToBinariesMap struct {
	Key    *string         `json:"key"`
	Values json.RawMessage `json:"values"`
}

type GetBinaryProductMapsResponseBinaryProductMaps struct {
	ResultCode           *string                                                             `json:"resultCode"`
	BinaryToProductsMap  []GetBinaryProductMapsResponseBinaryProductMapsBinaryToProductsMap  `json:"binaryToProductsMap"`
	ProductToBinariesMap []GetBinaryProductMapsResponseBinaryProductMapsProductToBinariesMap `json:"productToBinariesMap"`
}

type GetBinaryProductMapsResponse struct {
	BinaryProductMaps *GetBinaryProductMapsResponseBinaryProductMaps `json:"binaryProductMaps"`
}

type GetProfileRequest struct {
	Params *ProfileParams `json:"params" required:"true"`
}

type GetProfileResponseProfileMetrics struct {
	Name *string  `json:"name"`
	Mean *float64 `json:"mean"`
	P25  *float64 `json:"p25"`
	P95  *float64 `json:"p95"`
}

type GetProfileResponseProfile struct {
	ResultCode *string                            `json:"resultCode"`
	Tenants    *int64                             `json:"tenants"`
	Devices    *int64                             `json:"devices"`
	Metrics    []GetProfileResponseProfileMetrics `json:"metrics"`
}

type GetProfileResponse struct {
	Profile *GetProfileResponseProfile `json:"profile"`
}

type GetProfileSearchItemsRequest struct {
	Params *ProfileType `json:"params" required:"true"`
}

type GetProfileSearchItemsResponseProfileSearchItemsItems struct {
	Name  *string         `json:"name"`
	Extra json.RawMessage `json:"extra"`
}

type GetProfileSearchItemsResponseProfileSearchItems struct {
	ResultCode *string                                                `json:"resultCode"`
	Items      []GetProfileSearchItemsResponseProfileSearchItemsItems `json:"items"`
}

type GetProfileSearchItemsResponse struct {
	ProfileSearchItems *GetProfileSearchItemsResponseProfileSearchItems `json:"profileSearchItems"`
}

type GetProfileVersionsRequest struct {
	Params *ProfileVersionsParams `json:"params" required:"true"`
}

type GetProfileVersionsResponseProfileVersionsVersionsVersionInfoMetrics struct {
	Name *string  `json:"name"`
	Mean *float64 `json:"mean"`
}

type GetProfileVersionsResponseProfileVersionsVersionsVersionInfo struct {
	ResultCode   *string                                                               `json:"resultCode"`
	Tenants      *int64                                                                `json:"tenants"`
	Devices      *int64                                                                `json:"devices"`
	DevicesRatio *float64                                                              `json:"devicesRatio"`
	Metrics      []GetProfileVersionsResponseProfileVersionsVersionsVersionInfoMetrics `json:"metrics"`
}

type GetProfileVersionsResponseProfileVersionsVersions struct {
	Version     *string                                                       `json:"version"`
	VersionInfo *GetProfileVersionsResponseProfileVersionsVersionsVersionInfo `json:"versionInfo"`
}

type GetProfileVersionsResponseProfileVersions struct {
	ResultCode *string                                             `json:"resultCode"`
	Versions   []GetProfileVersionsResponseProfileVersionsVersions `json:"versions"`
}

type GetProfileVersionsResponse struct {
	ProfileVersions *GetProfileVersionsResponseProfileVersions `json:"profileVersions"`
}

type LookupBenchmarkRequest struct {
	Params *BenchmarkParams `json:"params" required:"true"`
}

type LookupBenchmarkResponseLookupLookupMatchesKey struct {
	Key   *string         `json:"key"`
	Value json.RawMessage `json:"value"`
}

type LookupBenchmarkResponseLookupLookupMatches struct {
	Key        []LookupBenchmarkResponseLookupLookupMatchesKey `json:"key"`
	Value      json.RawMessage                                 `json:"value"`
	Weight     *float64                                        `json:"weight"`
	Origins    []string                                        `json:"origins"`
	Precedence *int64                                          `json:"precedence"`
}

type LookupBenchmarkResponseLookup struct {
	ResultCode    *string                                      `json:"resultCode"`
	LookupMatches []LookupBenchmarkResponseLookupLookupMatches `json:"lookupMatches"`
}

type LookupBenchmarkResponse struct {
	Lookup *LookupBenchmarkResponseLookup `json:"lookup"`
}

type ProductPropertiesRequest struct {
	Params *ProductPropertiesParams `json:"params" required:"true"`
}

type ProductPropertiesResponseProductPropertiesProperties struct {
	Name  *string         `json:"name"`
	Value json.RawMessage `json:"value"`
}

type ProductPropertiesResponseProductProperties struct {
	ResultCode  *string                                                `json:"resultCode"`
	ProductName *string                                                `json:"productName"`
	Properties  []ProductPropertiesResponseProductPropertiesProperties `json:"properties"`
}

type ProductPropertiesResponse struct {
	ProductProperties *ProductPropertiesResponseProductProperties `json:"productProperties"`
}

type ProfilePropertiesRequest struct {
	Params *ProfilePropertiesParams `json:"params" required:"true"`
}

type ProfilePropertiesResponseProfilePropertiesProfileKey struct {
	BinaryName  *string `json:"binaryName"`
	ProductName *string `json:"productName"`
}

type ProfilePropertiesResponseProfilePropertiesProperties struct {
	Name  *string         `json:"name"`
	Value json.RawMessage `json:"value"`
}

type ProfilePropertiesResponseProfileProperties struct {
	ResultCode *string                                                `json:"resultCode"`
	ProfileKey *ProfilePropertiesResponseProfilePropertiesProfileKey  `json:"profileKey"`
	Properties []ProfilePropertiesResponseProfilePropertiesProperties `json:"properties"`
}

type ProfilePropertiesResponse struct {
	ProfileProperties *ProfilePropertiesResponseProfileProperties `json:"profileProperties"`
}
