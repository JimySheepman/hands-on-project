package swagger

import (
	"time"
)

type Body struct {
	Text string    `json:"text,omitempty"`
	Tags []string  `json:"tags,omitempty"`
	Due  time.Time `json:"due,omitempty"`
}
