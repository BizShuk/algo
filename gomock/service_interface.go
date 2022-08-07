package gomock

type DBobj interface {
	Call() []int
}

// [Pattern]: [Go Mock Testing] by inject mock object into service component
type Service struct {
	db DBobj
}

func (s *Service) query() {
	s.db.Call()
}
