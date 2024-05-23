package models

type Hachaton struct {
	id          int
	Page        string `json:"page"` //изображение
	description string `json:"description"`
}
