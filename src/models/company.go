package models

import (
	"gorm.io/gorm"
)

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
	gorm.Model
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

// Saves a company to the database
func (company *Company) Save() (*Company, error) {
	err := Database.Create(&company).Error
	if err != nil {
		return &Company{}, err
	}
	return company, nil
}

// Fetches all companies from the database
func FetchAllCompanies() (*[]Company, error) {
	var companies []Company
	err := Database.Find(&companies).Error
	if err != nil {
		return &[]Company{}, err
	}
	return &companies, nil
}

// Fetches a company from the database
func FetchCompany(id string) (*Company, error) {
	var company Company
	err := Database.Where("id = ?", id).First(&company).Error
	if err != nil {
		return &Company{}, err
	}
	return &company, nil
}

// Updates a company in the database
func UpdateCompany(id string, company *Company) (*Company, error) {
	err := Database.Model(&Company{}).Where("id = ?", id).Updates(company).Error
	if err != nil {
		return &Company{}, err
	}
	return company, nil
}

// Deletes a company from the database
func DeleteCompany(id string) error {
	err := Database.Where("id = ?", id).Delete(&Company{}).Error
	if err != nil {
		return err
	}
	return nil
}
