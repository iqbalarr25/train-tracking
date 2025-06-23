package helper

const (
	UserRoleAdmin               = "ADMIN"
	UserRoleOperator            = "OPERATOR"
	UserRoleSetter              = "SETTER"
	SuccessCreatedDataMessage   = "Data created successfully"
	SuccessUpdatedDataMessage   = "Data updated successfully"
	SuccessDeletedDataMessage   = "Data deleted successfully"
	SuccessRetrievedDataMessage = "Data retrieved successfully"
	AreaRoleBlock               = "BLOCK"
	AreaRoleFiningLine          = "FINING LINE"
	ExcavatorStatusActive       = "ACTIVE"
	ExcavatorStatusInactive     = "INACTIVE"
	ExcavatorStatusMaintenance  = "MAINTENANCE"
	SatelliteLogStatusQueue     = "QUEUE"
	SatelliteLogStatusSent      = "SENT"
)

var (
	FileType = map[string]string{
		"image/png":  "png",
		"image/jpeg": "jpg",
		"image/bmp":  "bmp",
		"image/webp": "webp",
	}
	DynamicFieldConditionQuery = "%s = ?"
)
