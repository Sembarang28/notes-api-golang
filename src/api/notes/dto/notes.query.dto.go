package dto

type NotesReadAllRequestQuery struct {
	Search string   `query:"search"`
	Tags   []string `query:"tags"`
}
