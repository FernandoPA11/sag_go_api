package resources

import "gorm.io/gorm"

type Ranch struct {
	gorm.Model

	Name       string `gorm:"unique;not null;default:null" json:"name"`
	PostalCode string `gorm:"not null;default:null" json:"postal_code"`
	Country    string `gorm:"not null;default:null" json:"country"`
	Phone      string `gorm:"not null;default:null" json:"phone"`
}

type Corral struct {
	gorm.Model

	Name         string `gorm:"unique;not null;default:null" json:"name"`
	AnimalTypeID string `gorm:"not null;default:null" json:"animal_type"`
	RanchID      int    `gorm:"not null;default:null" json:"ranch_id"`

	AnimalType AnimalType `gorm:"foreignKey:AnimalTypeID;references:ID"`
	Ranch      Ranch      `gorm:"foreignKey:RanchID;references:ID"`
}

type Specie struct {
	gorm.Model

	Name string `gorm:"unique;not null;default:null" json:"name"` //Bovine, Ovine, Caprine, Porcine, Equine, Canine, Feline
}

type Breed struct {
	gorm.Model

	Name         string `gorm:"unique;not null;default:null" json:"name"`
	AnimalTypeID string `gorm:"not null" json:"animal_type"`

	AnimalType AnimalType `gorm:"foreignKey:AnimalTypeID;references:ID"`
}

type AnimalType struct {
	gorm.Model

	Name string `gorm:"unique;not null;default:null" json:"name"`
}
