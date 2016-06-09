// Show that we can use `Expect: 100-continue` in go versions >= 1.6
//
// +build go1.6

package swift

import (
	"net/http"
	"time"
)

const disableExpectContinue = false

// SetExpectContinueTimeout sets ExpectContinueTimeout in the
// transport to the default value if not set.
func SetExpectContinueTimeout(transport *http.Transport) {
	if transport.ExpectContinueTimeout == 0 {
		transport.ExpectContinueTimeout = 1 * time.Second
	}
}
