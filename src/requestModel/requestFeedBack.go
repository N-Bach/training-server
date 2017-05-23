package requestModel

type RequestFeedback struct {
	Id string `gorethink:"id,omitempty" json:"id"`
	AuthorId string `gorethink:"authorId" json:"authorId"`
	Title string `gorethink:"title" json:"title"`
	Description string `gorethink:"description" json:"description"`
}