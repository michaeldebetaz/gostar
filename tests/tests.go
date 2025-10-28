package tests

import (
	"testing"

	. "github.com/delaneyj/gostar/elements"
	"github.com/stretchr/testify/assert"
	"github.com/valyala/bytebufferpool"
)

type result struct {
	Expected string
	Actual   ElementRenderer
}

func run(t *testing.T, results []result) {
	for _, result := range results {
		buf := bytebufferpool.Get()
		e := result.Expected
		err := result.Actual.Render(buf)
		assert.NoError(t, err)
		a := buf.String()
		assert.Equal(t, e, a)
		bytebufferpool.Put(buf)
	}
}
