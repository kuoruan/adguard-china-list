package service_test

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.kuoruan.net/adguard-upstream/service"
)

func TestChinaList_Update(t *testing.T) {
	s := service.NewChinaList()

	err := s.Update()

	assert.NoError(t, err)
	assert.NotEmpty(t, s.AcceleratedDomains)
}

func TestChinaList_Transform(t *testing.T) {
	s := service.NewChinaList()

	s.Transform()

	assert.Empty(t, s.AcceleratedDomains)
}

func TestChinaList_Client(t *testing.T) {
	s := service.NewChinaList()

	assert.IsType(t, http.Client{}, s.Client)
}

func TestChinaList_AcceleratedDomains(t *testing.T) {
	s := service.NewChinaList()

	assert.Empty(t, s.AcceleratedDomains)

	s.Update()

	assert.NotEmpty(t, s.AcceleratedDomains)
}
