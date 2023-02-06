package response

// Order ...
const (
	OrderNotSupportDelivery         = "order_not_support_delivery"
	OrderNoDeliveryService          = "order_no_delivery_service_found"
	OrderInvalidCODValue            = "order_invalid_cod_value"
	OrderInvalidValue               = "order_invalid_value"
	OrderInvalidCode                = "order_invalid_code"
	OrderCodeRequired               = "order_code_required"
	OrderNotFound                   = "order_not_found"
	OrderInvalidStatus              = "order_invalid_status"
	OrderExisted                    = "order_existed"
	OrderCannotApplyDeliveryService = "order_cannot_apply_delivery_service"
	OrderCannotUpdateFee            = "order_cannot_update_fee"
	OrderCannotUpdateStatus         = "order_cannot_update_status"
	OrderUnknownApproveMethod       = "order_unknown_approve_method"
	OrderCannotApprove              = "order_cannot_approve"
	OrderCannotCancelDelivery       = "order_cannot_cancel_delivery"
	OrderCannotCancel               = "order_cannot_cancel"
)

// 200 - 399
var order = []Code{
	{
		Key:     OrderNotSupportDelivery,
		Message: "không hỗ trợ vận chuyển đến địa chỉ này",
		Code:    200,
	},
	{
		Key:     OrderNoDeliveryService,
		Message: "không tìm thấy dịch vụ phù hợp",
		Code:    201,
	},
	{
		Key:     OrderInvalidCODValue,
		Message: "tiền thu hộ không hợp lệ",
		Code:    202,
	},
	{
		Key:     OrderInvalidValue,
		Message: "giá trị đơn hàng không hợp lệ",
		Code:    203,
	},
	{
		Key:     OrderInvalidCode,
		Message: "mã đơn không hợp lệ",
		Code:    204,
	},
	{
		Key:     OrderNotFound,
		Message: "không tìm thấy đơn hàng",
		Code:    205,
	},
	{
		Key:     OrderInvalidStatus,
		Message: "trạng thái không hợp lệ",
		Code:    206,
	},
	{
		Key:     OrderExisted,
		Message: "đơn hàng đã tồn tại",
		Code:    207,
	},
	{
		Key:     OrderCannotApplyDeliveryService,
		Message: "đơn hàng không thỏa điều kiện của dịch vụ vận chuyển",
		Code:    208,
	},
	{
		Key:     OrderCannotUpdateFee,
		Message: "không thể cập nhật phí của đơn hàng duyệt qua API",
		Code:    209,
	},
	{
		Key:     OrderCannotUpdateStatus,
		Message: "không thể cập nhật trạng thái của đơn chưa được duyệt hoặc duyệt qua API",
		Code:    210,
	},
	{
		Key:     OrderUnknownApproveMethod,
		Message: "không xác định phương thức duyệt đơn",
		Code:    211,
	},
	{
		Key:     OrderCannotApprove,
		Message: "chỉ duyệt đơn ở trạng thái chờ lấy hoặc tìm tài xế thất bại",
		Code:    212,
	},
	{
		Key:     OrderCannotCancelDelivery,
		Message: "chỉ hủy vận chuyển khi đơn duyệt qua api và các trạng thái chưa kết thúc",
		Code:    213,
	},
	{
		Key:     OrderCannotCancel,
		Message: "chỉ hủy đơn ở trạng thái chờ lấy hoặc tìm tài xế thất bại",
		Code:    213,
	},
	{
		Key:     OrderCodeRequired,
		Message: messageRequired("mã đơn hàng"),
		Code:    214,
	},
}
