package products

type Service interface {
	GetAll() ([]Product, error)
	Store(name, color string, price float64, stock int, code string, posted bool, dateCreated string) (Product, error)
	Update(id int, name, color string, price float64, stock int, code string, posted bool, dateCreated string) (Product, error)
	UpdateName(id int, name string) (Product, error)
	Delete(id int) error
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

func (s *service) GetAll() ([]Product, error) {
	ps, _ := s.repository.GetAll()
	return ps, nil
}

func (s *service) Store(name, color string, price float64, stock int, code string, posted bool, dateCreated string) (Product, error) {
	lastID, err := s.repository.LastID()

	if err != nil {
		return Product{}, err
	}

	lastID++

	product, err := s.repository.Store(lastID, name, color, price, stock, code, posted, dateCreated)
	if err != nil {
		return Product{}, err
	}

	return product, nil
}

func (s *service) Update(id int, name, color string, price float64, stock int, code string, posted bool, dateCreated string) (Product, error) {
	return s.repository.Update(id, name, color, price, stock, code, posted, dateCreated)
}

func (s *service) UpdateName(id int, name string) (Product, error) {
	return s.repository.UpdateName(id, name)
}

func (s *service) Delete(id int) error {
	return s.repository.Delete(id)
}
