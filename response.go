package livestatus

// Response represents a Livestatus query response.
type Response struct {
	Status  int
	Message string
	Records []Record
}

// Len returns the number of records present in the response.
func (r Response) Len() int {
	return len(r.Records)
}
