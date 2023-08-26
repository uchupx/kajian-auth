package model

type BaseModel struct{}

func (m *BaseModel) TableName() string {
	return ""
}
