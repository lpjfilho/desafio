package structs

import (
	"github.com/google/uuid"
	"time"
)

type PaidDocs struct {
	Value       float64
	DueDate     time.Time
	PaymentDate time.Time
	PartnerID   uuid.UUID
}
