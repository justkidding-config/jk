package cache

import (
	"io/ioutil"
	"testing"

	"github.com/google/go-containerregistry/pkg/name"
	"github.com/stretchr/testify/assert"
)

func mustParseRef(ref string) name.Reference {
	parsed, err := name.ParseReference(ref)
	if err != nil {
		panic(err)
	}
	return parsed
}

func TestCacheOverlayCreation(t *testing.T) {
	c := New("testfiles/dotcache")
	ov, err := c.FileSystemForImage(mustParseRef("image-repo:image-tag"))
	assert.NoError(t, err)

	// Open a file known to be present. This verifies that the overlay
	// has been constructed properly, but could also fail if say the
	// overlay implementation is broken.
	f, err := ov.Open("/foo/bar")
	assert.NoError(t, err)
	defer f.Close()

	contents, err := ioutil.ReadAll(f)
	assert.NoError(t, err)
	// and by prior arrangement ...
	assert.Equal(t, "This is the content in the top-most layer.\n", string(contents))
}
