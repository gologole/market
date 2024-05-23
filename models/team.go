package models

type RateStatistics struct {
}

type Team struct {
	ID          int            `json:"id"`
	Name        string         `json:"name"`
	Page        string         `json:"page"` //аватар команды
	Description string         `json:"description"`
	Teamate     []Teamate      `json:"teamate"` // можно использовать мапу с ключем айдишником пользователя
	isFully     bool           `json:"isFully"`
	whoNeed     []string       `json:"whoNeed"` //роли которые требуются для заполнения команды
	Rating      RateStatistics `json:"rating"`
}
