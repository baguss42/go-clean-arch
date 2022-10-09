package dto

// ListOption store data related option of list method
type ListOption struct {
	// Query is used to where params in database query
	Query  string
	Limit  int
	Offset int
}
