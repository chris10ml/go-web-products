/*
	Generar un Stub del Store cuya función “Read” retorne dos productos con las especificaciones que deseen.
	Comprobar que GetAll() retorne la información exactamente igual a la esperada.
*/

package products

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

// ------- STUB -------
type stubStore struct{}

func (s *stubStore) Read(data interface{}) error {
	p1 := Product{Id: 100, Name: "teclado", Color: "blanco", Price: 10.5, Stock: 10, Code: "bo100", Posted: true, DateCreated: "2022-01"}
	p2 := Product{Id: 200, Name: "mouse", Color: "negro", Price: 19.5, Stock: 15, Code: "te100", Posted: true, DateCreated: "2022-01"}
	listProducts := []Product{p1, p2}

	fileData, _ := json.Marshal(listProducts)

	return json.Unmarshal(fileData, &data)
}

func (s *stubStore) Write(data interface{}) error {
	return nil
}

// ------- MOCK -------

type mockStorage struct {
	Data          []Product
	ReadWasCalled bool
}

func (s *mockStorage) Read(data interface{}) error {
	s.ReadWasCalled = true

	fileData, _ := json.Marshal(s.Data)

	return json.Unmarshal(fileData, &data)
}

func (s *mockStorage) Write(data interface{}) error {
	return nil
}

func TestGetAll(t *testing.T) {
	db := stubStore{}
	repo := NewRepository(&db)

	req1 := Product{Id: 100, Name: "teclado", Color: "blanco", Price: 10.5, Stock: 10, Code: "bo100", Posted: true, DateCreated: "2022-01"}
	req2 := Product{Id: 200, Name: "mouse", Color: "negro", Price: 19.5, Stock: 15, Code: "te100", Posted: true, DateCreated: "2022-01"}

	expected := []Product{req1, req2}

	result, _ := repo.GetAll()

	assert.Equal(t, expected, result)
}

func TestUpdateName(t *testing.T) {
	beforeUpdate := Product{Id: 100, Name: "teclado", Color: "blanco", Price: 10.5, Stock: 10, Code: "bo100", Posted: true, DateCreated: "2022-01"}

	db := mockStorage{[]Product{beforeUpdate}, false}
	repo := NewRepository(&db)

	expected := "After Update"

	result, _ := repo.UpdateName(100, "After Update")

	assert.Equal(t, expected, result.Name)
	assert.True(t, db.ReadWasCalled)
}
