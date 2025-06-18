package dto

type NotesResponse struct {
	ID           string   `json:"id"`
	Name         string   `json:"name"`
	Notes        string   `json:"notes"`
	CategoryID   string   `json:"category_id"`
	CategoryName string   `json:"category_name"`
	Tags         []string `json:"tags"`
}
