package structs

import "github.com/google/uuid"

type PriceTable struct {
	PartnerID uuid.UUID
	Ranges    []Range
}

type Range struct {
	Start      int
	End        int
	Percentage float64
}
