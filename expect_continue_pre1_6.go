// Show that we can't use `Expect: 100-continue` in go versions < 1.6
//
// +build !go1.6

package swift

import "net/http"

const disableExpectContinue = true

// SetExpectContinueTimeout sets ExpectContinueTimeout in the
// transport to the default value if not set.
func SetExpectContinueTimeout(transport *http.Transport) {
	// can't do anything so ignore
}
