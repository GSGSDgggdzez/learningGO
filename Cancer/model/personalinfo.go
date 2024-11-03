package model

type PersonalInfo struct {
	ID          uint   `json:"id" gorm:"primaryKey"`
	FullName    string `json:"fullname"`
	Email       string `json:"email"`
	Phone       string `json:"phone"`
	DateOfBirth string `json:"date_of_birth"`
	Insurance   string `json:"insurance"`
	EmergencyNo string `json:"emergency"`
	Address     string `json:"address"`
	Occupation  string `json:"occupation"`
}
