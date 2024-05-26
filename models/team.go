package models

type Team struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Page        string    `json:"page"` //аватар команды
	Description string    `json:"description"`
	Teamate     []Teamate `json:"teamate"` // можно использовать мапу с ключем айдишником пользователя
	IsFully     bool      `json:"isFully"`
	WhoNeed     []string  `json:"whoNeed"` //роли которые требуются для заполнения команды
	Hachatons   int       `json:"hachatons"`
	Won         int       `json:"won"`
	Story       string    `json:"story"` //перечисление хакатонов в которых учавствовала команда
}
