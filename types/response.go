package types

type PaginatedData struct {
	PageCount int    `json:"page_count"`
	CurrPage  int    `json:"curr_page"`
	PrevPage  string `json:"prev_page"`
	NextPage  string `json:"next_page"`
	Pages     []int  `json:"pages"`
}
type Response struct {
	Code int
	Body any // can be data in read apis or string messages in create, update and delete apis
	Page PaginatedData
}
