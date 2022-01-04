package common

import "fmt"

//分页获取相关
const (
	APIValueDirectionDesc = 1  //降序
	APIValueDirectionASC  = 2  //生序
	APIMaxPageSize        = 50 //分页限制
)

const (
	ASC  = "ASC"
	DESC = "DESC"
)

//Page 页码参数
type Page struct {
	PageNo    int `json:"pageNo"    form:"pageNo"`
	PageSize  int `json:"pageSize"  form:"pageSize"`
	Direction int `json:"direction" form:"direction"`
}

//1：新的排在前；2：旧的排在前
func (p *Page) SortWord() string {
	if p.Direction != 2 {
		return DESC
	} else {
		return ASC
	}
}

func (p *Page) pageNo() int {
	page := p.PageNo
	if page < 1 {
		page = 1
	}
	return page - 1
}

func (p *Page) Offset() int {
	return p.pageNo() * p.Limit()
}

func (p *Page) Limit() int {
	limit := p.PageSize
	if limit < 1 {
		limit = 10
	}
	return limit
}

func (p *Page) Order(filed string) string {
	return fmt.Sprintf("%s %s", filed, p.SortWord())
}

func (p *Page) Check() WebError {
	if p.PageSize > APIMaxPageSize {
		return NewError(StatusInvalidParam, MsgInvalidParam+": pageSize is too large")
	}
	if p.Direction > APIValueDirectionASC {
		return NewError(StatusInvalidParam, MsgInvalidParam+": direction err")
	}
	if p.PageNo < 1 {
		p.PageNo = 1
	}
	if p.PageSize < 1 {
		p.PageSize = 10
	}

	return nil
}

func CheckPage(page *Page) error {
	if page.PageSize > APIMaxPageSize {
		return NewError(StatusInvalidParam, MsgInvalidParam+": pageSize")
	}

	if page.PageNo < 1 {
		page.PageNo = 1
	}
	if page.PageSize < 1 {
		page.PageSize = 10
	}

	return nil
}

//通用分页关键字查询参数
type Select struct {
	Page
	SelectValue string `json:"selectValue" form:"selectValue"`
}
