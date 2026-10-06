package appearance

type AssetName string

const (
	MenuLogo        AssetName = "menu-logo"
	LoginLogo       AssetName = "login-logo"
	LoginBackground AssetName = "login-bg"
)

type UpdateResponse struct {
	Result struct {
		Success bool `json:"success"`
	} `json:"result"`
	ResultStatus ResultStatus `json:"resultStatus"`
}
type ResultStatus struct {
	Code        int    `json:"code"`
	Description string `json:"description"`
}
