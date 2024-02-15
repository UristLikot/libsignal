package tests

import (
	"fmt"
	"testing"

	"github.com/UristLikot/libsignal/util/keyhelper"
)

func TestRegistrationID(t *testing.T) {
	regID := keyhelper.GenerateRegistrationID()
	fmt.Println(regID)
}
