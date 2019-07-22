package main

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/thenguyenit/testing/mockingsms/mocks"
)

func TestCharge(t *testing.T) {
	smsServiceMock := new(mocks.MessageService)
	smsServiceMock.On("SendNotification", 100).Return(true)

	m := MobifoneService{smsServiceMock}
	err := m.Charge(100)
	assert.NoError(t, err)

	smsServiceMock.AssertExpectations(t)
}
