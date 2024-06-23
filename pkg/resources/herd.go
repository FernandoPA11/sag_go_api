package resources

import "gorm.io/gorm"

type Herd struct {
	gorm.Model

	TagNumber  string    `gorm:"unique;not null;default:null" json:"tag_number"`
	Weight     int       `gorm:"not null;default:null" json:"weight"`
	BirthDate  string    `gorm:"not null;default:null" json:"birth_date"`
	BreedID    int       `gorm:"not null" json:"breed_id"`
	Procedence string    `gorm:"not null;default:null" json:"procedence"`
	SpecieID   int       `gorm:"not null" json:"specie_id"`
	RanchID    int       `gorm:"not null" json:"ranch_id"`
	Pictures   []Picture `gorm:"foreignKey:HerdID;references:ID"`

	Breed  Breed  `gorm:"foreignKey:BreedID;references:ID"`
	Specie Specie `gorm:"foreignKey:SpecieID;references:ID"`
	Ranch  Ranch  `gorm:"foreignKey:RanchID;references:ID"`
}

type Picture struct {
	ID     int    `gorm:"primaryKey;auto increment;default:null" json:"id"`
	HerdID int    `gorm:"not null" json:"cattle_id"`
	Url    string `gorm:"not null" json:"url"`

	herd herd `gorm:"foreignKey:HerdID;references:ID"`
}
