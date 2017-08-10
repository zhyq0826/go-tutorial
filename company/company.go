package company

// +gen slice:"Where,GroupBy[string],DistinctBy,SortBy,Select[string]"
type Company struct {
	Name    string
	Country string
	City    string
}
