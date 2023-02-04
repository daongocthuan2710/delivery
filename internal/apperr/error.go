package apperr

import "github.com/friendsofgo/errors"

var (
	OrderCodeExisted = errors.New("order_code_existed")
	PartnerNotFound  = errors.New("partner_not_found")
)
