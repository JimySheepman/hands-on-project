package times

import (
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const layout = "2006-01-02"

func TimeConverter(stringTime string) primitive.DateTime {
	t, err := time.Parse(layout, stringTime)
	if err != nil {
		log.Println("Error: time is not converted")
	}
	return primitive.NewDateTimeFromTime(t.AddDate(-1, 0, 0))
}
