package dto

type Order struct {
	ID        string `json:"id"`
	ProductId string `json:"product_id"`
	UserId    string `json:"user_id"`
	Amount    int64  `json:"amount"`
}
