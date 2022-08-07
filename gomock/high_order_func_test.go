package gomock

import (
	"testing"

	"github.com/stretchr/testify/mock"
)

type MockObject struct {
	mock.Mock
}

func (m *MockObject) MockSayHi(name string) {
	m.Called(name)
}

func TestCallSayHi(t *testing.T) {
	type args struct {
		name    string
		mockObj *MockObject
	}

	tests := []struct {
		name string
		args args
	}{
		{
			name: "Test MockSayHi",
			args: args{
				name:    "Shuk",
				mockObj: new(MockObject),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockObj := tt.args.mockObj
			mockObj.On("MockSayHi", tt.args.name)

			CallSayHi(tt.args.name, mockObj.MockSayHi)

			mockObj.AssertCalled(t, "MockSayHi", "Shuk")
			mockObj.AssertNumberOfCalls(t, "MockSayHi", 1)
			mockObj.AssertExpectations(t)

		})
	}
}
