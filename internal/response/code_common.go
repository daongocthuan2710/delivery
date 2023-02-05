package response

const (
	CommonSuccess             = "common_success"
	CommonBadRequest          = "common_bad_request"
	CommonUnauthorized        = "common_unauthorized"
	CommonNotFound            = "common_not_found"
	CommonInvalidChecksum     = "common_invalid_checksum"
	CommonInvalidPagination   = "common_invalid_pagination"
	CommonInvalidName         = "common_invalid_name"
	CommonInvalidPhone        = "common_invalid_phone"
	CommonInvalidEmail        = "common_invalid_email"
	CommonInvalidPassword     = "common_invalid_password"
	CommonInvalidPermissions  = "common_invalid_permissions"
	CommonExistedName         = "common_existed_name"
	CommonExistedPhone        = "common_existed_phone"
	CommonExistedEmail        = "common_existed_email"
	CommonInvalidDeliveryTime = "common_invalid_delivery_time"
	CommonInvalidDistance     = "common_invalid_distance"
	CommonInvalidPrice        = "common_invalid_price"
	CommonInvalidCourier      = "common_invalid_courier"
	CommonInvalidCode         = "common_invalid_code"
	CommonInvalidServiceCode  = "common_invalid_service_code"
	CommonInvalidWeight       = "common_invalid_weight"
	CommonInvalidVolume       = "common_invalid_volume"
	CommonInvalidProvince     = "common_invalid_province"
	CommonInvalidDistrict     = "common_invalid_district"
	CommonInvalidWard         = "common_invalid_ward"
	CommonInvalidMultiplier   = "common_invalid_multiplier"
	CommonInvalidTimeType     = "common_invalid_time_type"
	CommonInvalidTime         = "common_invalid_time"
	CommonInvalidDayOfWeek    = "common_invalid_day_of_week"
	CommonInvalidHour         = "common_invalid_hour"
	CommonInvalidApplyType    = "common_invalid_apply_type"
	CommonInvalidAddress      = "common_invalid_address"
	CommonInvalidQuantity     = "common_invalid_quantity"
	CommonInvalidAPIKey       = "common_invalid_api_key"
	CommonInvalidLocation     = "common_invalid_location"
	CommonInvalidWebhookURL   = "common_invalid_webhook_url"

	CommonDeliveryServiceNotFound = "common_delivery_service_not_found"
	CommonTokenRequired           = "common_token_required"
	CommonLoginGoogleFailed       = "common_login_google_failed"
	CommonNotExistsAccount        = "common_not_exists_account"
	CommonServiceNotFound         = "common_courier_service_not_found"
	CommonNotSupportCOD           = "common_not_support_cod"

	CommonCourierServiceApprovedViasRequired = "common_courier_service_approved_vias_required"
	CommonCourierCodesRequired               = "common_courier_codes_required"
	CommonConfigurationDoesNotExist          = "common_configuration_does_not_exist"
	CommonInvalidCash                        = "common_invalid_cash"

	CommonTargetRequired        = "common_target_required"
	CommonTargetIDRequired      = "common_target_id_required"
	CommonNoteRequired          = "common_note_required"
	CommonInvalidPercent        = "common_invalid_percent"
	CommonInvalidCOD            = "common_invalid_cod"
	CommonPaymentMethodRequired = "common_payment_method_required"
	CommonInvalidPaymentMethod  = "common_invalid_payment_method"
	CommonServiceCodeRequired   = "common_service_code_required"
	CommonErrorWhenHandler      = "common_error_when_handler"
	CommonServerIsProcessing    = "common_server_is_processing"
	CommonMissingAuthInfo       = "common_missing_auth_info"
	CommonInvalidStatus         = "common_invalid_status"
)

// 1 - 199
var common = []Code{
	{
		Key:     CommonSuccess,
		Message: "thành công",
		Code:    1,
	},
	{
		Key:     CommonBadRequest,
		Message: messageInvalid("dữ liệu"),
		Code:    2,
	},
	{
		Key:     CommonUnauthorized,
		Message: "bạn không có quyền thực hiện hành động này",
		Code:    3,
	},
	{
		Key:     CommonNotFound,
		Message: messageNotFound("dữ liệu"),
		Code:    4,
	},
	{
		Key:     CommonInvalidChecksum,
		Message: "xác thực dữ liệu thất bại",
		Code:    5,
	},
	{
		Key:     CommonInvalidPagination,
		Message: messageInvalid("phân trang"),
		Code:    6,
	},
	{
		Key:     CommonInvalidName,
		Message: messageInvalid("tên"),
		Code:    7,
	},
	{
		Key:     CommonInvalidPhone,
		Message: "số điện thoại không hợp lệ",
		Code:    8,
	},
	{
		Key:     CommonInvalidEmail,
		Message: "email không hợp lệ",
		Code:    9,
	},
	{
		Key:     CommonInvalidPassword,
		Message: "mật khẩu hợp lệ",
		Code:    10,
	},
	{
		Key:     CommonInvalidPermissions,
		Message: "danh sách quyền không hợp lệ",
		Code:    11,
	},
	{
		Key:     CommonExistedName,
		Message: "tên đã tồn tại trong hệ thống",
		Code:    12,
	},
	{
		Key:     CommonExistedPhone,
		Message: messageExisted("số điện thoại"),
		Code:    13,
	},
	{
		Key:     CommonExistedEmail,
		Message: messageExisted("email"),
		Code:    14,
	},
	{
		Key:     CommonInvalidDeliveryTime,
		Message: "thời gian giao hàng không hợp lệ",
		Code:    15,
	},
	{
		Key:     CommonInvalidDistance,
		Message: "khoảng cách không hợp lệ",
		Code:    16,
	},
	{
		Key:     CommonInvalidPrice,
		Message: "giá không hợp lệ",
		Code:    17,
	},
	{
		Key:     CommonInvalidCourier,
		Message: "hãng vận chuyển không hợp lệ",
		Code:    18,
	},
	{
		Key:     CommonInvalidCode,
		Message: "mã không hợp lệ",
		Code:    19,
	},
	{
		Key:     CommonInvalidServiceCode,
		Message: "mã dịch vụ không hợp lệ",
		Code:    20,
	},
	{
		Key:     CommonInvalidWeight,
		Message: "khối lượng không hợp lệ",
		Code:    21,
	},
	{
		Key:     CommonInvalidVolume,
		Message: "thể tích không hợp lệ: định dạng LxWxH",
		Code:    22,
	},
	{
		Key:     CommonInvalidProvince,
		Message: "tỉnh thành không hợp lệ",
		Code:    23,
	},
	{
		Key:     CommonInvalidMultiplier,
		Message: "hệ số không hợp lệ",
		Code:    24,
	},
	{
		Key:     CommonInvalidTimeType,
		Message: "loại thời gian không hợp lệ",
		Code:    25,
	},
	{
		Key:     CommonInvalidTime,
		Message: "thời gian không hợp lệ",
		Code:    26,
	},
	{
		Key:     CommonInvalidDayOfWeek,
		Message: "ngày trong tuần không hợp lệ",
		Code:    27,
	},
	{
		Key:     CommonInvalidHour,
		Message: "giờ không hợp lệ",
		Code:    28,
	},
	{
		Key:     CommonInvalidApplyType,
		Message: "loại áp dụng không hợp lệ",
		Code:    29,
	},
	{
		Key:     CommonInvalidDistrict,
		Message: "quận/huyện không hợp lệ",
		Code:    30,
	},
	{
		Key:     CommonInvalidWard,
		Message: "xã/phường không hợp lệ",
		Code:    31,
	},
	{
		Key:     CommonInvalidAddress,
		Message: "địa chỉ không hợp lệ",
		Code:    32,
	},
	{
		Key:     CommonInvalidQuantity,
		Message: "số lượng không hợp lệ",
		Code:    33,
	},
	{
		Key:     CommonInvalidAPIKey,
		Message: "API key không hợp lệ",
		Code:    34,
	},
	{
		Key:     CommonDeliveryServiceNotFound,
		Message: "không tìm thấy dịch vụ vận chuyển",
		Code:    35,
	},
	{
		Key:     CommonInvalidLocation,
		Message: "địa chỉ không đúng",
		Code:    36,
	},
	{
		Key:     CommonInvalidWebhookURL,
		Message: "Webhook URL không hợp lệ",
		Code:    37,
	},
	{
		Key:     CommonTokenRequired,
		Message: "token không được trống",
		Code:    38,
	},
	{
		Key:     CommonLoginGoogleFailed,
		Message: "không thể kết nối đến google, vui lòng thử lại sau",
		Code:    39,
	},
	{
		Key:     CommonNotExistsAccount,
		Message: "tài khoản không tồn tại",
		Code:    40,
	},
	{
		Key:     CommonServiceNotFound,
		Message: "không tìm thấy dịch vụ",
		Code:    41,
	},
	{
		Key:     CommonNotSupportCOD,
		Message: "dịch vụ không hỗ trợ COD",
		Code:    42,
	},
	{
		Key:     CommonCourierServiceApprovedViasRequired,
		Message: "hình thức phê duyệt của đơn vị vận chuyển không được trống",
		Code:    43,
	},
	{
		Key:     CommonCourierCodesRequired,
		Message: "mã của đơn vị vận chuyển không được trống",
		Code:    44,
	},
	{
		Key:     CommonConfigurationDoesNotExist,
		Message: "cấu hình không tồn tại",
		Code:    45,
	},
	{
		Key:     CommonInvalidCash,
		Message: "số tiền không hợp lệ",
		Code:    46,
	},
	{
		Key:     CommonTargetRequired,
		Message: messageRequired("target"),
		Code:    47,
	},
	{
		Key:     CommonTargetIDRequired,
		Message: messageRequired("target id"),
		Code:    48,
	},
	{
		Key:     CommonNoteRequired,
		Message: messageRequired("ghi chú"),
		Code:    49,
	},
	{
		Key:     CommonInvalidPercent,
		Message: messageInvalid("phần trăm"),
		Code:    50,
	},
	{
		Key:     CommonInvalidCOD,
		Message: messageInvalid("COD"),
		Code:    51,
	},
	{
		Key:     CommonInvalidPaymentMethod,
		Message: messageInvalid("phương thức thanh toán"),
		Code:    52,
	},
	{
		Key:     CommonPaymentMethodRequired,
		Message: messageRequired("phương thức thanh toán"),
		Code:    53,
	},
	{
		Key:     CommonServiceCodeRequired,
		Message: messageRequired("mã dịch vụ"),
		Code:    54,
	},
	{
		Key:     CommonErrorWhenHandler,
		Message: "có lỗi trong quá trình xử lý",
		Code:    55,
	},
	{
		Key:     CommonServerIsProcessing,
		Message: "Hệ thống đang xử lý yêu cầu, vui lòng thử lại sau giây lát!",
		Code:    56,
	},
	{
		Key:     CommonMissingAuthInfo,
		Message: "Thiếu thông tin xác thực, vui lòng kiếm tra lại",
		Code:    57,
	},
	{
		Key:     CommonInvalidStatus,
		Message: messageInvalid("trạng thái"),
		Code:    58,
	},
}
