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
	var meta = &Meta{}
	db, err := redis.Connect("master")
	if err != nil {
		return NewMeta(meta, "")
	}

	var cat = &Category{}

	res, err := db.HGet("categories", cast.ToString(id)).Result()
	if err != nil {
		return NewMeta(meta, "")
	}
	err = json.Unmarshal([]byte(res), cat)
	if err != nil {
		return NewMeta(meta, "")
	}
	meta.Title = cat.MetaTitle
	meta.Canonical = cat.Canonical
	meta.Robots = cat.MetaRobotTag
	meta.Description = cat.MetaDescription
	meta.Keyword = cat.MetaKeyword
	meta.Alternate = cat.Canonical
	return NewMeta(meta, "")
}

// GetMetaByURL page
func GetMetaByURL(url string) *Meta {
	var meta = &Meta{}
	db, err := redis.Connect("master")
	if err != nil {
		return NewMeta(meta, url)
	}
	res, err := db.HGet("metatags:url", url).Result()
	if err != nil {
		return NewMeta(meta, url)
	}

	err = json.Unmarshal([]byte(res), meta)
	if err != nil {
		return NewMeta(meta, url)
	}
	return NewMeta(meta, url)
}

// NewMeta NewMeta
func NewMeta(m *Meta, url string) *Meta {
	seo := viper.GetStringMapString("seo")
	if m.Canonical == "" {
		m.Canonical = seo["canonical"]
	}
	if m.Title == "" {
		m.Title = seo["title"]
	}
	if m.Image == "" {
		m.Image = seo["image"]
	}
	if m.FBAppID == "" {
		m.FBAppID = seo["fb_app_id"]
	}
	if m.FBAdmin == "" {
		m.FBAdmin = seo["fb_admin"]
	}
	if m.Site == "" {
		m.Site = seo["site"]
	}
	if m.Creator == "" {
		m.Site = seo["creator"]
	}
	if m.Alternate == "" {
		m.Alternate = m.Canonical
	}
	if m.Robots == "" {
		m.Robots = seo["robots"]
	}
	return m
}
