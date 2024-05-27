package models

// User модель для представления информации о пользователе
type User struct {
	ID           int    `json:"id"`            // Идентификатор пользователя
	Login        string `json:"name"`          // Имя пользователя
	Page         string `json:"page"`          //аватар
	Password     string `json:"password"`      //  парол, пользователя
	PasswordHash string `json:"password_hash"` // Хэш пароля пользователя
	Email        string `json:"email"`         // Email пользователя
	Role         string `json:"role"`          // Роль пользователя
	Address      string `json:"address"`       // Адрес пользователя
	PhoneNumber  string `json:"phone_number"`  // Номер телефона пользователя
	Description  string `json:"description"`   //описаниие для профиле сплошным текстом либо рабитое на поля(образование,опыт,контакты...)
	TeamID       int    `json:"team"`          //0 если нет команды и по дефолту
	Skills       string `json:"skills"`
	Hachatons    int    `json:"hachatons"`
	Won          int    `json:"won"`
	Story        string `json:"story"` //перечисление хакатонов в которых учавствовала команда
}
