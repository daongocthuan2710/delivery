package v1

import (
	"fmt"
	"io"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/thoas/go-funk"

	"delivery/internal/constant"
	"delivery/internal/response"
	"delivery/internal/service"
)

func NewWebhook(svc service.IDelivery) *Webhook {
	return &Webhook{deliverySvc: svc}
}

type Webhook struct {
	deliverySvc service.IDelivery
}

func (w Webhook) init(g *echo.Group) {
	r := g.Group("/webhook")

	r.POST("/:partnerCode", w.Webhook)
}

func (w Webhook) Webhook(c echo.Context) error {
	var (
		partnerIdentityCode = strings.ToUpper(c.Param("partnerCode"))
		b                   []byte
		partnerCodes        = []string{
			constant.TPLCodeGHN,
		}
	)
	if !funk.ContainsString(partnerCodes, partnerIdentityCode) {
		return response.R404(c, nil, "")
	}

	if c.Request().Body != nil {
		b, _ = io.ReadAll(c.Request().Body)
	}
	fmt.Println("------------------")
	fmt.Println("Webhook from: ", partnerIdentityCode)
	fmt.Println(string(b))
	if err := w.deliverySvc.UpdateStatusFromWebhook(b, partnerIdentityCode, c.RealIP()); err != nil {
		return response.R400(c, nil, err.Error())
	}
	return response.R200(c, nil, "")

}
