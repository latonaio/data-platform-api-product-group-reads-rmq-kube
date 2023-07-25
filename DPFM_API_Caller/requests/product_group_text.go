package requests

type ProductGroupText struct {
	ProductGroup		string  `json:"ProductGroup"`
	Language			string  `json:"Language"`
	ProductGroupName	string  `json:"ProductGroupName"`
	CreationDate		string	`json:"CreationDate"`
	LastChangeDate		string	`json:"LastChangeDate"`
	IsMarkedForDeletion	*bool	`json:"IsMarkedForDeletion"`
}
