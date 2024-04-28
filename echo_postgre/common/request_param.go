package common

type Filter struct {
	Status   string `json:"status" form:"status"`
	Username string `json:"username" form:"username"`
	Role     string `json:"role" form:"role"`
}

type Paging struct {
	Page  int `json:"page" form:"page"`
	Limit int `json:"limit" form:"limit"`
	Total int `json:"total" form:"-"`
}

func (p *Paging) Process() {
	if p.Page <= 0 {
		p.Page = 1
	}

	if p.Limit <= 0 || p.Limit > 50 {
		p.Limit = 10
	}
}
