package common

type DefaultBinder struct{}

type Filter struct {
	Status   string `json:"status" query:"status"`
	Username string `json:"username" query:"username"`
	Role     string `json:"role" query:"role"`
}

type Paging struct {
	Page  int   `json:"page" query:"page"`
	Limit int   `json:"limit" query:"limit"`
	Total int64 `json:"total"`
}

func (p *Paging) Process() {
	if p.Page <= 0 {
		p.Page = 1
	}

	if p.Limit <= 0 || p.Limit > 50 {
		p.Limit = 10
	}
}
