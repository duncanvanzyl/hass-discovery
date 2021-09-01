package discovery

import (
	"encoding/base64"
	"fmt"
	"strings"
)

// hash creates a hash of the object. Will be repeatable.
// There must be a better way to do this.
func hash(d interface{}) string {
	enc := []byte{}
	src := []byte(fmt.Sprintf("%v", d))
	base64.StdEncoding.Encode(src, enc)
	h := string(enc)
	h = strings.ReplaceAll(h, "=", "")
	if len(h) > 63 {
		h = h[:63]
	}
	return h
}
