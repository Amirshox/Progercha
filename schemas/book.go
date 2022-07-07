package schemas

type CreateBookSchema struct {
	Title  string `json:"title" binding:"required"`
	Author string `json:"author" binding:"required"`
}

type UpdateBookSchema struct {
	Title  string `json:"title"`
	Author string `json:"author"`
}
