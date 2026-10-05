package data_management

type Device struct {
	UID  string `json:"uid"`
	Name string `json:"name"`
}
type DeleteDevicesRequest struct {
	Devices []Device `json:"devices"`
}
type DeviceStatus struct {
	UID    string `json:"uid"`
	Name   string `json:"name"`
	Status string `json:"status"`
}
type DeleteDevicesResponse struct {
	ScheduledCount int            `json:"scheduledCount"`
	Status         string         `json:"status"`
	Devices        []DeviceStatus `json:"devices"`
}
