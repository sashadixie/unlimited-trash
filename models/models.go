package models

import (
	"time"
)

type Tabler interface {
	TableName() string
}

func (Beer) TableName() string {
	return "beer_items"
}

type Beer struct {
	ID          int64     `gorm:"primary_key" json:"id"`
	Tradename   string    `json:"tradename"`
	Developerid int64     `json:"developerid"`
	Birthday    time.Time `gorm:"column:birthday;default:now()" json:"birthday"`
	Description string    `json:"description"`
}

type Developer struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	CountryId   int64  `json:"country_id"`
}

type Country struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	Legal    bool   `json:"legal"`
	TaxIndex int64  `json:"taxIndex"`
}
