package platform

import "time"

type Task struct {
	name           string    `json:"name"`
	decription     string    `json:"decription"`
	attachmentsIds []string  `json:"attachmentsIds"`
	deadline       time.Time `json:"deadline"`
	price          int32     `json:"price"`
	fine           int32     `json:"fine"`
}
