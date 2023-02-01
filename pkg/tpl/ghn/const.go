package ghn

const (
	baseURLProd = "https://online-gateway.ghn.vn"
	baseURLDev  = "https://dev-online-gateway.ghn.vn"
)

const (
	pathEstimateFee  = "/shiip/public-api/v2/shipping-order/fee"
	pathCreateOrder  = "/shiip/public-api/v2/shipping-order/create"
	pathPreviewOrder = "/shiip/public-api/v2/shipping-order/preview"
	pathGetProvinces = "/shiip/public-api/master-data/province"
	pathGetDistricts = "/shiip/public-api/master-data/district"
	pathGetWards     = "/shiip/public-api/master-data/ward"
)

const (
	headerKeyToken       = "Token"
	headerKeyShopID      = "ShopId"
	headerKeyContentType = "Content-Type"

	headerValueContentTypeJSON = "application/json"
)

const (
	codeSuccess = 200
)

const (
	ServiceTypeExpress  = 1
	ServiceTypeStandard = 2
	ServiceTypeSaving   = 3
)

const (
	PaymentTypeSeller = 1
)
const (
	RequiredNoteCHOTHUHANG         = "CHOTHUHANG"
	RequiredNoteCHOXEMHANGKHONGTHU = "CHOXEMHANGKHONGTHU"
	RequiredNoteKHONGCHOXEMHANG    = "KHONGCHOXEMHANG"
)
