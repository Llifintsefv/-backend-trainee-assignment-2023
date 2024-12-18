package models

import "time"

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

type GetUserSegmentsResponse struct {
	SegmentId int        `json:"segment_id"`
	CreatedAt time.Time  `json:"created_at"`
	DeletedAt *time.Time `json:"deleted_at"`
	Slug      string     `json:"slug"`
}