package dto

// FooCreateRequest is request DTO for create foo with http
type FooCreateRequest struct {
	Title string `json:"title"`
}

// FooUpdateRequest is request DTO for update foo with http
type FooUpdateRequest struct {
	ID    string `json:"id"`
	Title string `json:"title"`
}

// FooResponse is response DTO for foo with http
type FooResponse struct {
	ID    string `json:"id"`
	Title string `json:"title"`
}

// FooResponse is response DTO for multiple foo
type FoosResponse struct {
	Foos []FooResponse `json:"foos"`
}
