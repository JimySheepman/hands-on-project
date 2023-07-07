package times

import (
	"reflect"
	"testing"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestTimeConverter(t *testing.T) {
	stringTime := "2016-01-26"
	got := TimeConverter(stringTime)
	var want primitive.DateTime
	if reflect.TypeOf(got) != reflect.TypeOf(want) {
		t.Errorf("Time converter is not succeed: got %v want %v", reflect.TypeOf(got), reflect.TypeOf(want))
	}
}
