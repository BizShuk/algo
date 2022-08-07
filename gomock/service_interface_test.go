package gomock

import (
	"testing"

	"github.com/stretchr/testify/mock"
)

type MockDB struct {
	mock.Mock
}

func (m *MockDB) Call() []int {
	args := m.Called() // call mock object, and return result from On event
	return args.Get(0).([]int)
}

func TestService_query(t *testing.T) {
	tests := []struct {
		name   string
		mockDB *MockDB
	}{
		{
			name:   "Mock Service Component",
			mockDB: new(MockDB),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockDB.On("Call").Return([]int{1, 2, 3, 4})

			s := &Service{
				db: tt.mockDB,
			}

			s.query()
			tt.mockDB.AssertCalled(t, "Call")
			tt.mockDB.AssertNumberOfCalls(t, "Call", 1)
			tt.mockDB.AssertExpectations(t)
		})
	}
}
