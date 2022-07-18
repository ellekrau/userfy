package contracts

type Response struct {
	ID string `json:"ID"`
}

func NewResponse(id string) Response {
	return Response{ID: id}
}
