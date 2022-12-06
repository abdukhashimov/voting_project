package domain

type Status struct {
	ID    string `json:"id"`
	Title string `json:"title"`
}

type StatusAllResp struct {
	Statuss []Status `json:"statuss"`
	Count   int      `json:"count"`
}

type StatusCreate struct {
	Title string `json:"title"`
}
