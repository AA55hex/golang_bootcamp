package connection

import (
	"errors"
	"testing"

	"github.com/AA55hex/golang_bootcamp/server/config"
	"github.com/stretchr/testify/mock"
	"github.com/upper/db/v4"
	"github.com/upper/db/v4/adapter/mysql"
)

type mocked_adapter struct {
	mock.Mock
	isFailTest bool
}

var fail_error = errors.New("")

func (m *mocked_adapter) Open(connURL db.ConnectionURL) (db.Session, error) {
	m.Called(connURL)
	if m.isFailTest {
		return nil, fail_error
	}
	return nil, nil

}

var db_settings = &mysql.ConnectionURL{
	Database: config.MySQL.Database,
	Host:     config.MySQL.Host,
	User:     config.MySQL.User,
	Password: config.MySQL.Password,
}

func TestOpenSessionOnSuccess(t *testing.T) {
	adp := new(mocked_adapter)
	adapter = adp

	adp.On("Open", db_settings).Return(nil, nil)

	OpenSession(db_settings, 1)

	adp.AssertExpectations(t)
}

func TestOpenSessionOnFail(t *testing.T) {
	adp := &mocked_adapter{isFailTest: true}
	adapter = adp

	adp.On("Open", db_settings).Return(nil, fail_error)

	OpenSession(db_settings, 2)

	adp.AssertExpectations(t)
}
