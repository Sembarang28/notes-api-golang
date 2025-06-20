package dto

type NotesReadAllRequestQuery struct {
	Search string   `query:"name"`
	Tags   []string `query:"tags"`
}
