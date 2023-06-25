// ロジック
package model

type User struct {
	ID       int    `json:"id" gorm:"praimaly_key"`
	Name     string `json:"name"`
	Password string `json:"password"`
}
