package model

import (
	uuid "github.com/satori/go.uuid"
	"log"
	"time"
)

type BaseEntity struct {
	Id        string    `gorm:"type:varchar(36);primary_key;NOT NULL"json:"id"`
	CreatedAt time.Time `gorm:"type:datetime;NOT NULL;DEFAULT:CURRENT_TIMESTAMP;"json:"createdAt"`
	CreatedBy string    `gorm:"type:varchar(36);NOT NULL"json:"createdBy"`
	UpdatedAt time.Time `gorm:"type:datetime;NOT NULL;DEFAULT:CURRENT_TIMESTAMP;"json:"updatedAt"`
	UpdatedBy string    `gorm:"type:varchar(36);NOT NULL"json:"updatedBy"`
}

/**
初始化entity
*/
func (e *BaseEntity) InitEntity(userId string) {
	e.Id = uuid.NewV4().String()
	cstZone := time.FixedZone("CST", 8*3600)
	e.CreatedBy = userId
	e.CreatedAt = time.Now().In(cstZone)
	e.UpdatedBy = userId
	log.Println("date setting: ", time.Now().In(cstZone))
	e.UpdatedAt = time.Now().In(cstZone)
}

func (e *BaseEntity) DoneUpdate(userId string) {
	cstZone := time.FixedZone("CST", 8*3600)
	e.UpdatedBy = userId
	log.Println("date setting: ", time.Now().In(cstZone))
	e.UpdatedAt = time.Now().In(cstZone)
}
