package models

// user model struct
type ListAs struct {
	ID   int    `json:"id"`
	Role string `json:"role" gorm:"type: varchar(255);primaryKey"`
}

type ListAsResponse struct {
	ID   int    `json:"id"`
	Role string `json:"role"`
}

func (ListAsResponse) TableName() string {
	return "list_as"
}
