package models

type Hachaton struct {
	Id          int
	Page        string   `json:"page"` //изображение
	Description string   `json:"description"`
	Participant []string `json:"participant"`
}
