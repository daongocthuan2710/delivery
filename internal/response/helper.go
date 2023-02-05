package response

import (
	"net/http"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/labstack/echo/v4"
)

func sendResponse(c echo.Context, httpCode int, success bool, data interface{}, message string, code int) error {
	if data == nil {
		data = echo.Map{}
	}

	return c.JSON(httpCode, Response{
		Success: success,
		Data:    data,
		Message: message,
		Code:    code,
	})
}

// InvalidChecksum ...
func InvalidChecksum(c echo.Context) error {
	key := CommonInvalidChecksum
	localeData := GetByKey(key)
	return sendResponse(c, http.StatusBadRequest, false, echo.Map{}, localeData.Message, localeData.Code)
}

// RouteValidation ...
// WARNING: only match with lib "github.com/go-ozzo/ozzo-validation/v4"
func RouteValidation(c echo.Context, err error) error {
	key := getMessage(err)

	// Return
	return R400(c, nil, key)
}

func getMessage(err error) string {
	err1, ok := err.(validation.Errors)
	if !ok {
		err2, ok := err.(validation.ErrorObject)
		if ok {
			return err2.Message()
		}
		return err.Error()
	}
	for _, item := range err1 {
		if item == nil {
			continue
		}
		return getMessage(item)
	}
	return err.Error()
}
