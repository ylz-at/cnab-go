package credentials

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/cnabio/cnab-go/utils/crud"
)

func TestCredentialStore_HandlesNotFoundError(t *testing.T) {
	cs := NewMockStore()
	mockStore := cs.GetBackingStore().GetDataStore().(crud.MockStore)
	mockStore.ReadMock = func(itemType string, name string) (bytes []byte, err error) {
		// Change the default error message to test that we are checking
		// inside the error message and not matching it exactly
		return nil, errors.New("wrapping error message: " + crud.ErrRecordDoesNotExist.Error())
	}

	_, err := cs.Read("missing cred set")
	assert.EqualError(t, err, ErrNotFound.Error())
}
