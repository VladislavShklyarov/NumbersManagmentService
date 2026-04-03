package requests

type ImportRequest struct {
	Numbers []string `json:"numbers"`
	Source  string   `json:"source"`
}
