package response

type PageResult struct {
	List     interface{} `json:"list"`
	Count    uint64       `json:"count"`
	Page     int         `json:"page"`
	PageSize int         `json:"pageSize"`
}
