package main

import (
	"desafio/structs"
	"github.com/google/uuid"
	"log/slog"
	"os"
	"time"

	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	log := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(log)
}

func getPartners() []structs.Partner {
	return []structs.Partner{
		{
			ID:   uuid.MustParse("d1445b3f-7e22-4123-909f-801bf65f1ea3"),
			Name: "Partner 1",
		},
		{
			ID:   uuid.MustParse("c2bd75df-7100-4061-b71a-1fa4a6f38112"),
			Name: "Partner 2",
		},
		{
			ID:   uuid.MustParse("29783785-404b-4765-8a3a-ede4e9df873d"),
			Name: "Partner 3",
		},
	}
}

func getPriceTable() []structs.PriceTable {
	return []structs.PriceTable{
		{
			PartnerID: uuid.MustParse("d1445b3f-7e22-4123-909f-801bf65f1ea3"),
			Ranges: []structs.Range{
				{
					Start:      1,
					End:        30,
					Percentage: 5.,
				},
				{
					Start:      31,
					End:        60,
					Percentage: 10.,
				},
				{
					Start:      61,
					End:        9999999,
					Percentage: 30.,
				},
			},
		},
		{
			PartnerID: uuid.MustParse("c2bd75df-7100-4061-b71a-1fa4a6f38112"),
			Ranges: []structs.Range{
				{
					Start:      1,
					End:        30,
					Percentage: 3.,
				},
				{
					Start:      31,
					End:        60,
					Percentage: 7.,
				},
				{
					Start:      61,
					End:        9999999,
					Percentage: 10.,
				},
			},
		},
		{
			PartnerID: uuid.MustParse("29783785-404b-4765-8a3a-ede4e9df873d"),
			Ranges: []structs.Range{
				{
					Start:      1,
					End:        30,
					Percentage: 2.,
				},
				{
					Start:      31,
					End:        60,
					Percentage: 5.,
				},
				{
					Start:      61,
					End:        9999999,
					Percentage: 12.,
				},
			},
		},
	}
}

func getPaidDocs() []structs.PaidDocs {
	return []structs.PaidDocs{
		{
			Value:       100.,
			DueDate:     time.Date(2025, 02, 01, 10, 0, 0, 0, time.UTC),
			PaymentDate: time.Date(2025, 03, 02, 10, 0, 0, 0, time.UTC),
			PartnerID:   uuid.MustParse("d1445b3f-7e22-4123-909f-801bf65f1ea3"),
		},
		{
			Value:       100.,
			DueDate:     time.Date(2025, 02, 01, 10, 0, 0, 0, time.UTC),
			PaymentDate: time.Date(2025, 04, 02, 10, 0, 0, 0, time.UTC),
			PartnerID:   uuid.MustParse("d1445b3f-7e22-4123-909f-801bf65f1ea3"),
		},
		{
			Value:       100.,
			DueDate:     time.Date(2025, 01, 01, 10, 0, 0, 0, time.UTC),
			PaymentDate: time.Date(2025, 04, 02, 10, 0, 0, 0, time.UTC),
			PartnerID:   uuid.MustParse("c2bd75df-7100-4061-b71a-1fa4a6f38112"),
		},
		{
			Value:       100.,
			DueDate:     time.Date(2025, 02, 01, 10, 0, 0, 0, time.UTC),
			PaymentDate: time.Date(2025, 02, 01, 10, 0, 0, 0, time.UTC),
			PartnerID:   uuid.MustParse("29783785-404b-4765-8a3a-ede4e9df873d"),
		},
	}
}
