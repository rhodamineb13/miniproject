package dto

type AddBookDTO struct {
	ISBN        string `json:"isbn"`
	Title       string `json:"title"`
	Author      string `json:"author"`
	Description string `json:"description"`
	Quantity    uint   `json:"quantity"`
}

type ShowBooksDTO struct {
	ID          uint   `json:"id"`
	ISBN        string `json:"isbn"`
	Title       string `json:"title"`
	Author      string `json:"author"`
	Description string `json:"description"`
	Quantity    string `json:"quantity"`
}
