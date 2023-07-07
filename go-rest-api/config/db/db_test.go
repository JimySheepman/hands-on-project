package db

import "testing"

func TestConnectDB(t *testing.T) {

	t.Run("mongo client conncetion error ", func(t *testing.T) {
		got, _ := ConnectDB()

		if got != nil {
			t.Errorf("got %v, want %v", got, nil)
		}
	})

	t.Run("mongo client ping error ", func(t *testing.T) {
		got, _ := ConnectDB()

		if got != nil {
			t.Errorf("got %v, want %v", got, nil)
		}
	})

	t.Run("mongo database name error ", func(t *testing.T) {
		got, _ := ConnectDB()

		if got != nil {
			t.Errorf("got %v, want %v", got, nil)
		}
	})
}
