package appearance

import "encoding/json"

// LegacyAsset preserves the identifier/version encoding returned by the portal.
// Its image body is base64Blob, unlike saveAsset's blob request field.
type LegacyAsset struct {
	ID         json.RawMessage `json:"id"`
	Version    json.RawMessage `json:"version"`
	Name       AssetName       `json:"name,omitempty"`
	MIMEType   string          `json:"mimeType"`
	Base64Blob string          `json:"base64Blob"`
}
type SaveLegacyAssetRequest struct {
	ID       json.RawMessage `json:"id"`
	Version  json.RawMessage `json:"version"`
	Name     AssetName       `json:"name"`
	Filename string          `json:"filename"`
	// Blob contains base64 bytes without a data: URL prefix.
	Blob string `json:"blob"`
}
type LegacyAssetError struct {
	Code    json.RawMessage `json:"code"`
	Message string          `json:"message,omitempty"`
}
type SaveLegacyAssetResult struct {
	Error *LegacyAssetError `json:"error,omitempty"`
	// Raw retains unspecified server result fields.
	Raw json.RawMessage `json:"-"`
}

func (r *SaveLegacyAssetResult) UnmarshalJSON(data []byte) error {
	type plain SaveLegacyAssetResult
	var v plain
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	*r = SaveLegacyAssetResult(v)
	r.Raw = append(r.Raw[:0], data...)
	return nil
}
func (r SaveLegacyAssetResult) MarshalJSON() ([]byte, error) {
	if len(r.Raw) > 0 {
		return r.Raw, nil
	}
	type plain SaveLegacyAssetResult
	return json.Marshal(plain(r))
}

type SaveLegacyAssetResponse struct {
	Result       *SaveLegacyAssetResult `json:"result"`
	ResultStatus *ResultStatus          `json:"resultStatus,omitempty"`
}
