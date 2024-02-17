package models

const (
	QueryOrganization string = "organization"
	QueryRepository   string = "repository"
	QuerySpan         string = "span"
)

type UsageRequest struct {
	Organization string `json:"organization"`
	Repository   string `json:"repository"`
	Span         int    `json:"span"`
}

type UsageResponse struct {
	Usages    []Usage `json:"usages"`
	IsDisable bool    `json:"isDisable"`
}

type Usage struct {
	PodName string `json:"podName"`
	CPU     int64  `json:"cpu"`
	MEM     int64  `json:"memory"`
	Storage int64  `json:"storage"`
}
