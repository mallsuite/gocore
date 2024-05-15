package ml

type BaseOrder struct {
	Sort string
	Sidx string
}

type WhereExt struct {
	Column string
	Val    any
	Symbol int
}

type BaseList struct {
	Page int `json:"page"  d:"1"  v:"min:0#分页号码错误"  dc:"分页号码"`
	Size int `json:"size" d:"10" v:"max:500#分页数量最大500条"  dc:"分页数量"`

	Sort string `json:"sort"  d:"asc"  dc:"排序条件"`
	Sidx string `json:"sidx"  dc:"排序字段"`

	Order     []*BaseOrder
	WhereExt  []*WhereExt
	WhereLike []*WhereExt
}

type BaseListOutput struct {
	Items   []*any
	Page    int
	Total   int
	Records int
	Size    int
}
