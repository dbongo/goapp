package err

// ValidationErr ...
type ValidationErr struct {
	Message string
}

// Msg ...
func (err *ValidationErr) Msg() string {
	return err.Message
}

// HTTPErr ...
type HTTPErr struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
	URL        string `json:"url"`
}

// Msg ...
func (err *HTTPErr) Msg() string {
	return err.Message
}
