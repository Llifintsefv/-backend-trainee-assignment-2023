package models

type Segment struct {
	Id             int    `json:"id"`
	Slug           string `json:"slug"`
	AutoAddPercent int    `json:"auto_add_percent"`
}

type SegmentResponse struct {
	Status string  `json:"status"`
	Id     int     `json:"id"`
	Data   Segment `json:"segment"`
}

type UserSegmentRequest struct {
	Add    []string          `json:"add"`
	Remove []string          `json:"remove,omitempty"`
	TTL    map[string]string `json:"ttl,omitempty"`
}
