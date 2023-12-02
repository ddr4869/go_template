package common

type NewCommonErrorParam struct {
	Error      error
	StatusCode int
	Reasons    []string
}
