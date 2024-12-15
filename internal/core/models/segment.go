package models

type Segment struct {
	Slug string `json:"slug"`
	AutoAddPercent int  `json:"auto_add_percent"`
}

type SegmentResponse struct {
    Status string      `json:"status"`
	Id     int         `json:"id"`
    Data   Segment 		`json:"segment"`
}



type UserSegment struct {
	UserId int `json:"user_id"`
	SegmentId int `json:"segment_id"`
	AutoAdded bool  `json:"auto_added"`
	ExpiresAt string  `json:"expires_at"`
}

