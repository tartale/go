package filez

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPathWithoutExtension(t *testing.T) {

	assert.Equal(t, "", PathWithoutExtension(""))
	assert.Equal(t, "foo", PathWithoutExtension("foo"))
	assert.Equal(t, "foo", PathWithoutExtension("foo.bar"))
	assert.Equal(t, "foo/bar", PathWithoutExtension("foo/bar.baz"))
}

func TestNameWithoutExtension(t *testing.T) {

	assert.Equal(t, "", NameWithoutExtension(""))
	assert.Equal(t, "foo", NameWithoutExtension("foo"))
	assert.Equal(t, "foo", NameWithoutExtension("foo.bar"))
	assert.Equal(t, "bar", NameWithoutExtension("foo/bar.baz"))
}
