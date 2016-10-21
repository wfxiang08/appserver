package backends

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"fmt"
)

//
// go test git.chunyu.me/feiwang/appserver/backends -v -run "TestIpaParser"
//
func TestIpaParser(t *testing.T) {
	//TODO
	path := "/Users/feiwang/goprojects/apps/src/git.chunyu.me/feiwang/appserver/apps_root/app.zip"
	metaInfo, err := ParseIpa(path, "chunyu")

	assert.NoError(t, err)

	fmt.Printf("CFBundleDisplayName: %s\n", metaInfo["CFBundleDisplayName"])
	fmt.Printf("CFBundleShortVersionString: %s\n", metaInfo["CFBundleShortVersionString"])
}
