package api

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	mockdb "github.com/Kiminoto1412/simplebank/db/mock"
	db "github.com/Kiminoto1412/simplebank/db/sqlc"
	"github.com/Kiminoto1412/simplebank/util"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestGetAccountAPI(t *testing.T) {
	account := randomAccount()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	store := mockdb.NewMockStore(ctrl)

	// build stubs
	store.EXPECT().GetAccount(gomock.Any(), gomock.Eq(account.ID)).Times(1).Return(account, nil)

	// start test server and send request
	server := NewServer(store)
	// dont have to start a real http server so just use httptest.NewRecorder() to create a new ResponseRecorder
	recorder := httptest.NewRecorder()
	url := fmt.Sprintf("/accounts/%d", account.ID)
	// nil => for get don't have to have req.body
	request, err := http.NewRequest(http.MethodGet, url, nil)

	// validate
	require.NoError(t, err)

	server.router.ServeHTTP(recorder, request)

	// check response
	require.Equal(t, http.StatusOK, recorder.Code)
}

func randomAccount() db.Account {
	return db.Account{
		ID: util.RandomInt(1, 1000),
	}
}
