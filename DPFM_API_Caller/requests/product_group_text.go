package requests

type ProductGroupText struct {
	ProductGroup     string  `json:"ProductGroup"`
	Language         string  `json:"Language"`
	ProductGroupName *string `json:"ProductGroupName"`
}
