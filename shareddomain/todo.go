package shareddomain

type RequestCreate struct {
	Title string `json:"title"`
}

type ResponseWithoutData struct {
	Code    int    `json:"code"`
	Success bool   `json:"success"`
	Message string `json:"message"`
}