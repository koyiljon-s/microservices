// job-service/internal/model/job.go
package model

import (
	"time"
	"gorm.io/gorm"
	"github.com/google/uuid"
)
// BaseModel
type BaseModel struct {
	ID             uuid.UUID      `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      gorm.DeletedAt `gorm:"index"`
}

// Status(open or closed)
type Status string
const (
	StatusOpen   Status = "open"
	StatusClosed Status = "closed"
)

// Employment type 
type EmploymentType string
const (
	FullTime   EmploymentType = "full_time"
	Intern     EmploymentType = "intern"
)

// Salary 
type SalaryType string
const (
	SalaryExact      SalaryType = "exact"
	SalaryNegotiable SalaryType = "negotiable"
)

// Experience level
type ExperienceType string
const (
	ExperienceNewbie     ExperienceType = "newbie"
	ExperienceYears      ExperienceType = "years"
)

// Industry/Role Type
type IndustryType string
const (
	IndustryIT                 IndustryType = "it_technologies"
	IndustryRetailEcommerce    IndustryType = "retail_ecommerce"
	IndustryManufacturing      IndustryType = "manufacturing"
	IndustryConstruction       IndustryType = "construction_engineering"
	IndustryAgricultureFood    IndustryType = "agriculture_food_processing"
	IndustryEducationTraining  IndustryType = "education_training"
	IndustryHealthcare         IndustryType = "healthcare_medical"
	IndustryFinanceBanking     IndustryType = "finance_banking"
	IndustryHospitalityTourism IndustryType = "hospitality_tourism"
	IndustryTransportLogistics IndustryType = "transport_logistics"
	IndustryGovernmentSociety  IndustryType = "government_society"
	IndustrySalesMarketing     IndustryType = "sales_marketing"
	IndustryManagementOffice   IndustryType = "management_office"
	IndustryCreativeDesign     IndustryType = "creative_design"
	IndustryOther              IndustryType = "other_services"
)


type Job struct {
	BaseModel
    
	Title             string            `json:"title" gorm:"not null;index"`
	Status            Status            `json:"status" gorm:"default:'open';index"`
	CompanyName       string            `json:"company_name" gorm:"not null;index"`
    
	EmploymentType    EmploymentType    `json:"employment_type" gorm:"not null"`
	SalaryType        SalaryType        `json:"salary_type" gorm:"not null"`
	Salary            *int              `json:"salary,omitempty"`

	ExperienceLevel   ExperienceType    `json:"experience_level" gorm:"not null"`
	ExperienceYears   *int              `json:"experience_years,omitempty"`

	IndustryType      IndustryType      `json:"industry_type" gorm:"not null;index"`
	IndustryCustom    *string           `json:"industry_custom,omitempty"`

	Location          string            `json:"location" gorm:"not null;index"`

	Requirements      string            `json:"requirements" gorm:"type:text;not null"`
}