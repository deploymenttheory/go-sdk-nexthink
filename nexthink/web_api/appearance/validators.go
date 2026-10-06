package appearance

import (
	"fmt"
	"mime"
	"strings"
)

func validateName(n AssetName) error {
	switch n {
	case MenuLogo, LoginLogo, LoginBackground:
		return nil
	}
	return fmt.Errorf("unsupported appearance asset")
}
func validateUpload(n AssetName, typ string, data []byte) error {
	if err := validateName(n); err != nil {
		return err
	}
	media, _, err := mime.ParseMediaType(typ)
	if err != nil || !strings.HasPrefix(media, "image/") {
		return fmt.Errorf("an image media type is required")
	}
	if len(data) == 0 {
		return fmt.Errorf("image bytes are required")
	}
	return nil
}
