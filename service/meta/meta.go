package meta

import (
	"encoding/json"

	"github.com/spf13/viper"

	"github.com/spf13/cast"
	"github.com/unicolony/hayes/source/redis"
)

// Meta model
type Meta struct {
	Title       string
	Canonical   string
	Robots      string
	Description string
	Keyword     string
	Image       string
	FBAppID     string
	FBAdmin     string
	Site        string
	Creator     string
	Alternate   string `json:"alternate"`
}

// Category Meta
type Category struct {
	ID              string `json:"id"`
	Description     string `json:"description"`
	MetaTitle       string `json:"meta_title"`
	MetaKeyword     string `json:"meta_keyword"`
	MetaDescription string `json:"meta_description"`
	MetaRobotTag    string `json:"meta_robottag"`
	Canonical       string `json:"canonical"`
	Alternate       string `json:"alternate"`
}

// GetMetaCategory get
func GetMetaCategory(id int) *Meta {
	var meta = NewMeta()
	db, err := redis.Connect("master")
	if err != nil {
		return meta
	}

	var cat = &Category{}

	res, err := db.HGet("categories", cast.ToString(id)).Result()
	if err != nil {
		return meta
	}
	err = json.Unmarshal([]byte(res), cat)
	if err != nil {
		return meta
	}
	meta.Title = cat.MetaTitle
	meta.Canonical = cat.Canonical
	meta.Robots = cat.MetaRobotTag
	meta.Description = cat.MetaDescription
	meta.Keyword = cat.MetaKeyword
	meta.Alternate = cat.Canonical
	return meta
}

// GetMetaByURL page
func GetMetaByURL(url string) *Meta {
	var meta = NewMeta()
	db, err := redis.Connect("master")
	if err != nil {
		return meta
	}
	res, err := db.HGet("metatags:url", url).Result()
	if err != nil {
		return meta
	}
	err = json.Unmarshal([]byte(res), meta)
	if err != nil {
		return meta
	}
	meta.Alternate = meta.Canonical
	return meta
}

// NewMeta NewMeta
func NewMeta() *Meta {
	seo := viper.GetStringMapString("seo")
	return &Meta{
		Title:     "Home",
		Canonical: seo["canonical"],
		Image:     seo["image"],
		FBAppID:   seo["fb_app_id"],
		FBAdmin:   seo["fb_admin"],
		Site:      seo["site"],
		Creator:   seo["creator"],
	}
}
