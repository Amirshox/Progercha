package schemas

type SchemaArticle struct {
	ID     string `json:"id" validate:"uuid"`
	Title  string `json:"title" validate:"required,lowercase"`
	Body   string `json:"body" validate:"required,lowercase"`
	Author string `json:"author" validate:"required,lowercase"`
}
