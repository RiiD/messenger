package file

type line struct {
	Headers map[string][]string `json:"headers"`
	Body    []byte              `json:"body"`
}
