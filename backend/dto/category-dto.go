package dto

// CategoryDTO ...
type CategoryDTO struct {
	CategoryID int    `json:"category_id"`
	Name       string `json:"name"`
}

// CategoryData ...
type CategoryData struct {
	Categories []CategoryDTO `json:"categories"`
}

// CategoryResponse ...
type CategoryResponse struct {
	Data     CategoryData `json:"data"`
	Error    int          `json:"error"`
	ErrorMsg string       `json:"error_msg"`
}
