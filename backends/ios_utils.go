package backends

import (
	log "git.chunyu.me/golang/cyutils/utils/rolling_log"
	"fmt"
	"net/url"
	"gopkg.in/flosch/pongo2.v3"
)

func init() {
	pongo2.RegisterFilter("itemservice_url", GenerateItemServiceUrlFilter)
}

func GenerateItemServiceUrlFilter(in *pongo2.Value, param *pongo2.Value) (out *pongo2.Value, err *pongo2.Error) {
	plist := in.String()
	result := GenerateItemServiceUrl(plist)
	return pongo2.AsValue(result), nil
}
func GenerateItemServiceUrl(plist string) (out string) {

	out = fmt.Sprintf("itms-services://?action=download-manifest&url=%s", url.QueryEscape(plist))
	//out = "测试"
	log.Infof("Input: %s, Output: %s", plist, out)
	return out
}
