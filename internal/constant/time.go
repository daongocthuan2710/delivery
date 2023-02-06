package constant

import "time"

const (
	TimeLayoutDDMMYYYY = "02/01/2006"
)

var (
	TimezoneHCM, _ = time.LoadLocation("Asia/Ho_Chi_Minh")
)
