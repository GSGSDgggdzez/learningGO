package model

type MedicalInfo struct {
	ID                uint   `json:"id" gorm:"primaryKey"`
	MainDoctor        string `json:"main_doctor"`
	InsuranceProvider string `json:"insurance_provider"`
	InsuranceNumber   string `json:"insurance_number"`
	MedicalHistory    string `json:"medical_history"`
	Medications       string `json:"medications"`
	FamilyMedications string `json:"Family_medications"`

	Allergies string `json:"allergies"`
}
