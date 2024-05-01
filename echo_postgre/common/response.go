package common

type Response struct {
	StatusCode int
	Message    string
	Data       interface{}
}

type ResponseWithPagination struct {
	Data       interface{}
	Pagination Paging
}

func HandleResponseWithPagination(data interface{}, paging Paging) *ResponseWithPagination {
	return &ResponseWithPagination{
		Data:       data,
		Pagination: paging,
	}
}
