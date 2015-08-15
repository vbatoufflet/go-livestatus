package livestatus

// Response is a query response.
type Response struct {
	Status  int
	Records []Record
}

// Len returns the number of records present in the response.
func (r Response) Len() int {
	return len(r.Records)
}
