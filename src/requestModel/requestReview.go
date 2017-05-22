package requestModel

type RequestReview struct {
	ReviewerId string `gorethink:"reviewerid" json:"reviewerid"`
	Title string `gorethink:"title" json:"title"`
	Description string `gorethink:"description" json:"description"`
	Rated int `gorethink:"rated" json:"rated"`
	For string `gorethink:"for" json:"for"`
}