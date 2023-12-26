package models

import "gorm.io/gorm"

type Company struct {
	gorm.Model
	ID          string      `json:"id"`
	Name        string      `json:"name"`
	LogoURL     string      `json:"logo_url"`
	WebsiteURL  string      `json:"website_url"`
	SocialMedia SocialMedia `json:"social_media"  gorm:"embedded"`
	Slogan      string      `json:"slogan"`
	Industry    string      `json:"industry"`
	CEO         string      `json:"ceo"`
	Since       string      `json:"since"`
	Headquater  Office      `json:"headquarter" gorm:"embedded"`
}

type Office struct {
	Country string `json:"country"`
	Town    string `json:"town"`
	Venue   string `json:"venue"`
}

type SocialMedia struct {
	Facebook  string `json:"facebook"`
	Whatsapp  string `json:"whatsapp"`
	X         string `json:"x"`
	LinkedIn  string `json:"linkedin"`
	Instagram string `json:"instagram"`
}
