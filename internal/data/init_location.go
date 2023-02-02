package data

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"regexp"
	"strings"

	"github.com/thoas/go-funk"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"

	"delivery/internal/config"
	"delivery/internal/constant"
	"delivery/internal/model/entity"
	"delivery/pkg/httpclient"
	"delivery/pkg/tpl/ghn"
)

type CommonLoc struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func Init(db *sql.DB, cfg *config.Config) {
	initLocation(db, cfg)
}

func initLocation(db *sql.DB, cfg *config.Config) {
	ctx := context.Background()
	total, err := entity.Provinces().Count(ctx, db)
	if err != nil {
		log.Println("Count province err:", err)
		return
	}
	if total > 0 {
		log.Println("province already init. Skip init...")
		return
	}

	provURL := "http://api.ecomx.vn/user/list/province"
	disURL := "http://api.ecomx.vn/user/list/district"
	wardURL := "http://api.ecomx.vn/user/list/ward"
	var provinces, districts, wards []CommonLoc
	httpClient := httpclient.NewClient().SetDebug(false)
	resp, err := httpClient.NewRequest().SetResult(provinces).Get(provURL)
	if err != nil {
		log.Println("Get province err: ", err)
		return
	}
	ppt := resp.Result().(*[]CommonLoc)
	ghnClient := ghn.New(cfg.GHN.Token, cfg.GHN.ShopID, cfg.App.ENV == constant.ENVProd)
	var ghnProvinces []ghn.Province
	ghnProv, err := ghnClient.GetProvinces(ctx)
	if err != nil {
		log.Println("GHN - province err: ", err)
	}
	if ghnProv != nil {
		ghnProvinces = ghnProv.Data
	}
	for _, p := range *ppt {
		pro := entity.Province{
			ID:   null.IntFrom(p.ID),
			Name: null.StringFrom(p.Name),
		}
		found := funk.Find(ghnProvinces, func(item ghn.Province) bool {
			return renameProvince(item.ProvinceName) == renameProvince(p.Name)
		})
		var ghnDistricts []ghn.District
		if found != nil {
			gp := found.(ghn.Province)
			pro.GHNID = null.IntFrom(gp.ProvinceID)
			pro.GHNCode = null.StringFrom(gp.Code)

			gd, err := ghnClient.GetDistricts(ctx, gp.ProvinceID)
			if err != nil {
				log.Println("GHN - district err: ", err)
			}
			if gd != nil {
				ghnDistricts = gd.Data
			}
		}
		if err := pro.Insert(ctx, db, boil.Infer()); err != nil {
			log.Println("Insert province err:", err)
		}
		// district
		resp, err = httpClient.NewRequest().
			SetQueryParam("id", fmt.Sprintf("%d", p.ID)).
			SetResult(districts).
			Get(disURL)
		dpt := resp.Result().(*[]CommonLoc)
		for _, dis := range *dpt {
			d := entity.District{
				ID:         null.IntFrom(dis.ID),
				Name:       null.StringFrom(dis.Name),
				ProvinceID: pro.ID,
			}
			foundD := funk.Find(ghnDistricts, func(item ghn.District) bool {
				return renameProvince(item.DistrictName) == renameProvince(dis.Name)
			})
			var ghnWards []ghn.Ward
			if foundD != nil {
				ghnD := foundD.(ghn.District)
				d.GHNID = null.IntFrom(ghnD.DistrictID)
				d.GHNCode = null.StringFrom(ghnD.Code)

				gw, err := ghnClient.GetWards(ctx, ghnD.DistrictID)
				if err != nil {
					log.Println("GHN - ward err: ", err)
				}
				if gw != nil {
					ghnWards = gw.Data
				}
			}
			if err := d.Insert(ctx, db, boil.Infer()); err != nil {
				log.Println("Insert district err:", err)
			}
			// ward
			resp, err = httpClient.NewRequest().
				SetQueryParam("id", fmt.Sprintf("%d", dis.ID)).
				SetResult(wards).
				Get(wardURL)

			wpt := resp.Result().(*[]CommonLoc)
			for _, loc := range *wpt {
				ward := entity.Ward{
					ID:         null.IntFrom(loc.ID),
					Name:       null.StringFrom(loc.Name),
					DistrictID: d.ID,
				}

				foundW := funk.Find(ghnWards, func(item ghn.Ward) bool {
					return renameProvince(item.WardName) == renameProvince(loc.Name)
				})
				if foundW != nil {
					w := foundW.(ghn.Ward)
					ward.GHNCode = null.StringFrom(w.WardCode)
				}
				if err := ward.Insert(ctx, db, boil.Infer()); err != nil {
					log.Println("Insert ward err:", err)
				}
			}
		}
	}
}

func renameDistrict(s string) string {
	keys := []string{
		"quận ",
		"quận ",
		"huyện ",
		"thị trấn ",
		"thị xã ",
		"thành phố ",
	}
	return rename(s, keys)
}

func renameWard(s string) string {
	keys := []string{
		"xã ",
		"phường 0",
		"phường ",
		"thị trấn ",
		"tt ",
	}
	return rename(s, keys)
}
func renameProvince(s string) string {
	keys := []string{
		"tỉnh ",
		"thành phố ",
		"tp. ",
	}
	return rename(s, keys)
}
func rename(s string, keys []string) string {
	for _, key := range keys {
		s = strings.Replace(strings.ToLower(s), key, "", 1)
	}
	s = strings.Replace(s, "hoà", "hòa", 1)
	s = strings.Replace(s, "thuỷ", "thủy", 1)
	s = strings.Replace(s, "thuỵ", "thụy", 1)
	s = strings.TrimSpace(s)
	// t := transform.Chain(norm.NFD, runes.Remove(runes.In(unicode.Mn)), norm.NFC)
	// s, _, _ = transform.String(t, s)

	s = replaceStringWithRegex(s, `[^a-zA-Z0-9\s]`, "")
	return replaceStringWithRegex(s, `\s+`, " ")
}

// replaceStringWithRegex ...
func replaceStringWithRegex(src string, regex string, replaceText string) string {
	reg := regexp.MustCompile(regex)
	return reg.ReplaceAllString(src, replaceText)
}
