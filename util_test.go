package genmarkdown_test

import (
	"io"
	"io/ioutil"

	. "github.com/onsi/gomega"
)

func Slurp(r io.Reader) string {
	data, err := ioutil.ReadAll(r)
	Expect(err).To(BeNil())
	return string(data)
}
