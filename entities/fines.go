package entities

import (
	"time"
)

type Fines struct {
	FinesId     *uint   `json:"finesId,omitempty" gorm:"type:serial;not null;primary_key;auto_increment;unique_index"`
	BuildingsId *uint   `json:"ฺbuildingsId,omitempty" gorm:"type:int;not null;index"`
	BillType    *string `json:"billtype,omitempty" gorm:"varchar(255);not null"`
	Title       *string `json:"title,omitempty" gorm:"type:varchar(255);not null"`
	Agreement   *string `json:"agreement,omitempty" gorm:"type:varchar(255);not null"`
	Note        *string `json:"note,omitempty" gorm:"type:text"`

	CreatedAt *time.Time `json:"createdAt,omitempty" gorm:"type:timestamp;not null;default:current_timestamp"`
	UpdatedAt *time.Time `json:"updatedAt,omitempty" gorm:"type:timestamp;not null;default:current_timestamp"`
	DeletedAt *time.Time `json:"deletedAt,omitempty" gorm:"type:timestamp"`

	ArrBuildingsId *[]uint `json:"ฺarrBuildingsId,omitempty" gorm:"-"`

	// Buildings Buildings `json:"buildings,omitempty" gorm:"foreignkey:BuildingsId;association_foreignkey:BuildingsId"`
}
