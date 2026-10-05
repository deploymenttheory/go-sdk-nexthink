package data_management

import (
	"fmt"
	"strings"
)

func ValidateDeleteDevicesRequest(req *DeleteDevicesRequest) error {
	if req == nil || len(req.Devices) == 0 || len(req.Devices) > 100 {
		return fmt.Errorf("devices must contain 1 to 100 entries")
	}
	for i, d := range req.Devices {
		if strings.TrimSpace(d.UID) == "" || strings.TrimSpace(d.Name) == "" {
			return fmt.Errorf("devices[%d]: uid and name are required", i)
		}
	}
	// Malformed nonempty UIDs are deliberately passed to the server: the API
	// reports INVALID per device while scheduling the rest of the batch.
	return nil
}
