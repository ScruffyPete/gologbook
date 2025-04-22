package api

type ProjectRequestBody struct {
	Title string `json:"title"`
}

// type GetProjectParams struct {
// 	ID   string `json:"id"`
// 	Name string `json:"name"`
// }

// type GetProjectResponse struct {
// 	Code    int     `json:"code:"`
// 	Project Project `json:"project"`
// }

type Error struct {
	Code    int
	Message string
}
