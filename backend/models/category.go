package models

// Category response category
type Category struct {
	CategoryID int         `json:"category_id" gorm:"primarykey"`
	Name       string      `json:"name"`
	Image      interface{} `json:"image"`
}

// Data ...
type Data struct {
	Categories []Category `json:"categories"`
	TextColor  string     `json:"text_color"`
}

// CatResponse ...
type CatResponse struct {
	Data     Data   `json:"data"`
	Error    int    `json:"error"`
	ErrorMsg string `json:"error_msg"`
}
