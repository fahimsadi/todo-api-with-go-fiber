package routes

type Todo struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Status string `json:"status"`
}

var todos = []Todo{} // Slices as test database for now. Will replace with redis or other db
var idCounter = 1
