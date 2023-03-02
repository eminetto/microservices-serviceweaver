package feedback

import (
	"github.com/ServiceWeaver/weaver"
	"github.com/google/uuid"
)

type Feedback struct {
	ID    uuid.UUID
	Email string
	Title string `json:"title"`
	Body  string `json:"body"`
	weaver.AutoMarshal
}
