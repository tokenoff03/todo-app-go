package todo

type User struct {
	ID       int    `json:"-" db:"id"`                   // Правильно заключенный тег
	Name     string `json:"name" binding:"required"`     // Правильно заключенные теги и добавлен знак двоеточия
	Username string `json:"username" binding:"required"` // Правильно заключенные теги и добавлен знак двоеточия
	Password string `json:"password" binding:"required"` // Правильно заключенные теги и добавлен знак двоеточия
}
