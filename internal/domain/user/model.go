package user

import (
	"time"

	"github.com/evermos/boilerplate-go/shared/nuuid"
	"github.com/gofrs/uuid"
	"github.com/guregu/null"
)

type User struct {
	Id   			uuid.UUID		`db:"id"`
	Name 			string			`db:"name"`
	Type			string	 		`db:"type"`
	CreatedAt       	time.Time  	`db:"createdAt" `
	CreatedBy     	uuid.UUID   	`db:"createdBy"`
	UpdatedAt       	null.Time   `db:"updatedAt"`
	UpdatedBy     	nuuid.NUUID		`db:"updatedBy"`
	DeletedAt       	null.Time   `db:"deletedAt"`
	DeletedBy     	nuuid.NUUID		`db:"deletedBy"`
}



type UserRequestFormat struct {
	Name	string   `json:"name" validate:"required"`
	Type	string   `json:"type" validate:"required"`
}


type UserResponseFormat struct {
	Id   			uuid.UUID		`json:"id"`
	Name 			string			`json:"name"`
	Type			string	 		`json:"type"`
	CreatedAt       	time.Time  	`json:"createdAt" `
	CreatedBy     	uuid.UUID   	`json:"createdBy"`
	UpdatedAt       	null.Time   `json:"updatedAt,omitempty"`
	UpdatedBy     	uuid.UUID		`json:"updatedBy,omitempty"`
	DeletedAt       	null.Time   `json:"deletedAt,omitempty"`
	DeletedBy     	uuid.UUID		`json:"deletedBy,omitempty"`
}