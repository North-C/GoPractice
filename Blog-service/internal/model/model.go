package model


type Model struct{
	ID uint32 `gorm:"primary_key" json:"id"`
	CreatedBy string `json: "created_by"`
	CreatedOn uint32 `json: "created_on"`
	ModifiedBy string `json: "modified_by"`
	ModifiedOn uint32 `json: "modified_on"`
	DeletedOn uint32 `json: "deleted_on"`
	IsDel uint8 `json: "is_del"`
}