package models

type Pagination struct {
	Page      int `form:"page,default=0"`
	PageCount int `form:"pagecount,default=10"`
}
