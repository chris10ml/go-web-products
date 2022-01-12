package products

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

// ------- MOCK -------

type mockStore struct {
	Data          []Product
	ReadWasCalled bool
}

var input Product = Product{Id: 100, Name: "teclado", Color: "blanco", Price: 10.5, Stock: 10, Code: "bo100", Posted: true, DateCreated: "2022-01"}

func (s *mockStore) Read(data interface{}) error {
	s.ReadWasCalled = true

	fileData, _ := json.Marshal(s.Data)

	return json.Unmarshal(fileData, &data)
}

func (s *mockStore) Write(data interface{}) error {
	return nil
}

func TestUpdate(t *testing.T) {

	db := mockStore{[]Product{input}, false}
	repo := NewRepository(&db)

	expected := Product{Id: 100, Name: "After Update", Color: "blanco", Price: 10.5, Stock: 10, Code: "bo100", Posted: true, DateCreated: "2022-01"}

	result, _ := repo.Update(100, "After Update", "blanco", 10.5, 10, "bo100", true, "2022-01")

	assert.Equal(t, expected, result)
	assert.True(t, db.ReadWasCalled)
}

func TestDelete(t *testing.T) {
	db := mockStore{[]Product{input}, false}
	repo := NewRepository(&db)

	err := repo.Delete(100)

	assert.Nil(t, err)
}
