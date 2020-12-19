package models

//Productos servira para dar formato a los productos
type Productos struct {
	ID          int    `json:"id"`
	ProductCode string `json:"product_code"`
	Description string `json:"description"`
}
