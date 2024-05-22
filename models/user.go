package models

// User модель для представления информации о пользователе
type User struct {
	ID           int    `json:"id"`            // Идентификатор пользователя
	Login        string `json:"name"`          // Имя пользователя
	PasswordHash string `json:"password_hash"` // Хэш пароля пользователя
	Email        string `json:"email"`         // Email пользователя
	Role         string `json:"role"`          // Роль пользователя
	Address      string `json:"address"`       // Адрес пользователя
	PhoneNumber  string `json:"phone_number"`  // Номер телефона пользователя
}

// Price структура для представления цены в рублях и копейках
type Price struct {
	Rubles  int `json:"rubles"`  // Целая часть цены в рублях
	Kopecks int `json:"kopecks"` // Дробная часть цены в копейках
}

// Product модель для представления информации о продукте
type Product struct {
	ID          int    `json:"id"`          // Идентификатор продукта
	Name        string `json:"name"`        // Наименование продукта
	Page        string `json:"page"`        //изображение твоара
	Description string `json:"description"` // Описание продукта
	Price       Price  `json:"price"`       // Цена продукта в рублях и копейках
	Category    string `json:"category"`    // Категория продукта
}

// Order модель для представления информации о заказе
type Order struct {
	ID         int       `json:"id"`          // Идентификатор заказа
	UserID     int       `json:"user_id"`     // Идентификатор пользователя, оформившего заказ
	Products   []Product `json:"products"`    // Список продуктов в заказе
	TotalPrice Price     `json:"total_price"` // Общая цена заказа в рублях и копейках
	Status     string    `json:"status"`      // Статус заказа
}

// CartItem модель для представления элемента корзины
type CartItem struct {
	ID       int     `json:"id"`       // Идентификатор элемента корзины
	Product  Product `json:"product"`  // Продукт в элементе корзины
	Quantity int     `json:"quantity"` // Количество продукта
}

// Cart модель для представления информации о корзине
type Cart struct {
	UserID     int        `json:"user_id"`     // Идентификатор пользователя, к которому относится корзина
	Items      []CartItem `json:"items"`       // Список элементов корзины
	Subtotal   Price      `json:"subtotal"`    // Промежуточная сумма корзины в рублях и копейках
	Total      Price      `json:"total"`       // Общая сумма корзины в рублях и копейках
	Tax        Price      `json:"tax"`         // Сумма налога в корзине в рублях и копейках
	Discount   Price      `json:"discount"`    // Сумма скидки в корзине в рублях и копейках
	GrandTotal Price      `json:"grand_total"` // Итоговая сумма корзины после применения скидки и налога в рублях и копейках
}
