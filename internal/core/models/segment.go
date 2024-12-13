package models

type Segment struct {
	ID int 
	Slug string
	AutoAddPercent int 
}



type UserSegment struct {
	UserId int
	SegmentId int
	AutoAdded bool 
	ExpiresAt string
}