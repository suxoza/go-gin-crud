package models

// type Model struct {
// 	ID uint `gorm:"primaryKey"`
// }

// type ParentModel struct {
// 	ID        uint           `gorm:"primaryKey"`
// 	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
// }

type Post struct {
	ID uint `gorm:"primaryKey"`
	// ParentModel
	Title string `json:"title"`
	Body  string `json:"description"`
}
