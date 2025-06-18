package dto

type NotesRequest struct {
	Name       string   `json:"name" validate:"required"`
	Notes      string   `json:"notes" validate:"required"`
	CategoryID string   `json:"category_id" validate:"required"`
	Tags       []string `json:"tags" validate:"required"`
}
