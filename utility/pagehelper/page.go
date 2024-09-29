package pagehelper

type PageReq struct {
	PageNum  int    `json:"pageNum" description:"当前页码" d:"1"`
	PageSize int    `json:"pageSize" description:"每页数" d:"10"`
	OrderBy  string `json:"orderBy" description:"排序方式"`
}

func (p *PageReq) GetPageNum() int {
	return (p.PageNum - 1) * p.PageSize
}

func (p *PageReq) GetPageSize() int {
	return p.PageSize
}

func (p *PageReq) GetOrderBy() string {
	// todo maybe deal
	return p.OrderBy
}

type PageRes struct {
	Total interface{} `json:"total"`
}
