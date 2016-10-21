package backends

import (
	"fmt"
	// log "git.chunyu.me/golang/cyutils/utils/rolling_log"
	"github.com/stretchr/testify/assert"
	"testing"
)

var plistData string = `<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE plist PUBLIC "-//Apple//DTD PLIST 1.0//EN" "http://www.apple.com/DTDs/PropertyList-1.0.dtd">
<plist version="1.0">
  <dict>
    <key>items</key>
    <array>
      <dict>
        <key>assets</key>
        <array>
          <dict>
            <key>kind</key>
            <string>software-package</string>
            <key>url</key>
            <string>__URL__</string>
          </dict>
        </array>
        <key>metadata</key>
        <dict>
          <key>bundle-identifier</key>
          <string>com.chunyu.DiabetesManagerUser</string>
          <key>bundle-version</key>
          <string>1.4.0</string>
          <key>subtitle</key>
          <string>1.4.0_Unversioned directory</string>
          <key>kind</key>
          <string>software</string>
          <key>title</key>
          <string>春雨糖管家(Online)</string>
        </dict>
      </dict>
    </array>
  </dict>
</plist>`

//
// go test backends -v -run "TestPlist"
//
func TestPlist(t *testing.T) {

	//pwd, _ := filepath.Abs(".")
	//fmt.Printf("pwd: %v\n", pwd)
	//plist, err := UnmarshalFile("backends/testdata/app.plist")

	var plist Plist
	err := Unmarshal([]byte(plistData), &plist)

	assert.NoError(t, err)

	fmt.Printf("Root: %v\n", plist.Root)
	root := plist.Root.(Dict)
	items := root["items"].(Array)[0].(Dict)
	metadata := items["metadata"].(Dict)

	fmt.Printf("metadata: %v\n", metadata)
}
