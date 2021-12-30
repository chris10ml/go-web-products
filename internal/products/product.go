package products

type Product struct {
	Id          int     `json:"id"`
	Name        string  `json:"name"`
	Color       string  `json:"color"`
	Price       float64 `json:"price"`
	Stock       int     `json:"stock"`
	Code        string  `json:"code"`
	Posted      bool    `json:"posted"`
	DateCreated string  `json:"date_created"`
}
