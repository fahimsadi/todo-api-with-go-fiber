package routes

type Todo struct {
	ID     string `json:"id"`
	Task   string `json:"task"`
	Status bool   `json:"status"`
}
