package ghn

import (
	"context"
	"delivery/pkg/httpclient"
	"fmt"
)

func New(token string, shopID string, isProd bool) *Client {
	return &Client{
		Token:        token,
		ShopID:       shopID,
		IsProduction: isProd,
	}
}

type Client struct {
	Token        string
	ShopID       string
	IsProduction bool
}

func (c *Client) CreateOrder(ctx context.Context, req CreateOrderReq) (*CreateOrderRes, error) {
	api := c.getAPIURL(pathCreateOrder)
	resp, err := httpclient.NewRequest().
		SetContext(ctx).
		SetHeader(headerKeyShopID, c.ShopID).
		SetHeader(headerKeyToken, c.Token).
		SetBody(req).
		SetResult(CreateOrderRes{}).
		SetError(CreateOrderRes{}).
		Post(api)
	if err != nil {
		return nil, fmt.Errorf("ghn: create_order: %v", err)
	}
	result := resp.Result().(*CreateOrderRes)
	if resp.IsError() || result.Code != codeSuccess {
		return nil, fmt.Errorf("ghn: create_order: %s", result.Message)
	}
	return result, nil
}

func (c *Client) PreviewOrder(ctx context.Context, req PreviewOrderReq) (*PreviewOrderRes, error) {
	api := c.getAPIURL(pathPreviewOrder)
	resp, err := httpclient.NewRequest().
		SetContext(ctx).
		SetHeader(headerKeyShopID, c.ShopID).
		SetHeader(headerKeyToken, c.Token).
		SetHeader(headerKeyContentType, headerValueContentTypeJSON).
		SetBody(req).
		SetResult(PreviewOrderRes{}).
		SetError(PreviewOrderRes{}).
		Post(api)
	if err != nil {
		return nil, fmt.Errorf("ghn: preview_order: %v", err)
	}
	result := resp.Result().(*PreviewOrderRes)
	if resp.IsError() || result.Code != codeSuccess {
		return nil, fmt.Errorf("ghn: preview_order: %s", result.Message)
	}
	return result, nil
}

func (c *Client) GetProvinces(ctx context.Context) (*ProvinceRes, error) {
	api := c.getAPIURL(pathGetProvinces)
	resp, err := httpclient.NewRequest().
		SetContext(ctx).
		SetHeader(headerKeyShopID, c.ShopID).
		SetHeader(headerKeyToken, c.Token).
		SetHeader(headerKeyContentType, headerValueContentTypeJSON).
		SetResult(ProvinceRes{}).
		SetError(ProvinceRes{}).
		Get(api)
	if err != nil {
		return nil, fmt.Errorf("ghn: preview_order: %v", err)
	}
	result := resp.Result().(*ProvinceRes)
	if resp.IsError() || result.Code != codeSuccess {
		return nil, fmt.Errorf("ghn: get_province: %s", result.Message)
	}
	return result, nil
}

func (c *Client) GetDistricts(ctx context.Context, provinceID int) (*DistrictRes, error) {
	api := c.getAPIURL(pathGetDistricts)
	resp, err := httpclient.NewRequest().
		SetContext(ctx).
		SetHeader(headerKeyShopID, c.ShopID).
		SetHeader(headerKeyToken, c.Token).
		SetHeader(headerKeyContentType, headerValueContentTypeJSON).
		SetResult(DistrictRes{}).
		SetError(DistrictRes{}).
		Get(api)
	if err != nil {
		return nil, fmt.Errorf("ghn: get_district: %v", err)
	}
	result := resp.Result().(*DistrictRes)
	if resp.IsError() || result.Code != codeSuccess {
		return nil, fmt.Errorf("ghn: get_district: %s", result.Message)
	}
	return result, nil
}

func (c *Client) GetWards(ctx context.Context, districtID int) (*WardRes, error) {
	api := c.getAPIURL(pathGetWards)
	resp, err := httpclient.NewRequest().
		SetContext(ctx).
		SetHeader(headerKeyShopID, c.ShopID).
		SetHeader(headerKeyToken, c.Token).
		SetHeader(headerKeyContentType, headerValueContentTypeJSON).
		SetResult(WardRes{}).
		SetError(WardRes{}).
		Get(api)
	if err != nil {
		return nil, fmt.Errorf("ghn: get_ward: %v", err)
	}
	result := resp.Result().(*WardRes)
	if resp.IsError() || result.Code != codeSuccess {
		return nil, fmt.Errorf("ghn: get_ward: %s", result.Message)
	}
	return result, nil
}

func (c *Client) getAPIURL(path string) string {
	baseURL := baseURLDev
	if c.IsProduction {
		baseURL = baseURLProd
	}
	return baseURL + path
}
