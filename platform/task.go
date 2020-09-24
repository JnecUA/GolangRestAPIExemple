package platform

import "time"

//Task type to work with db table
type Task struct {
	Name           string    `json:"name"`
	Decription     string    `json:"decription"`
	AttachmentsIds []string  `json:"attachmentsIds"`
	Deadline       time.Time `json:"deadline"`
	Price          int32     `json:"price"`
	Fine           int32     `json:"fine"`
}
