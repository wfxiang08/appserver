package controllers

import (
	"github.com/astaxie/beego"
	"strings"
	"path"
	"io/ioutil"
	log "git.chunyu.me/golang/cyutils/utils/rolling_log"
	"fmt"
	"net/url"
	"git.chunyu.me/feiwang/appserver/backends"
	"github.com/oal/beego-pongo2"
)

type MainController struct {
	beego.Controller
}

//
// @Router /
//
func (this *MainController) Get() {
	userAgent := this.Ctx.Request.Header.Get("User-Agent")


	isAndroid := strings.Index(userAgent, "Android") != -1
	isIos := strings.Index(userAgent, "iPhone") != -1

	platform := this.GetString("platform", "Android")

	appsRoot := beego.AppConfig.String("apps_root")
	iosAppDirs, androidDirs, _ := backends.ListAppDir(appsRoot)

	// 参考: https://github.com/oal/beego-pongo2
	context := pongo2.Context{
		"platform": platform,
		"is_android": isAndroid,
		"is_ios": isIos,
		"is_web": !isIos && !isAndroid,
		"ios_app_dirs": iosAppDirs,
		"android_app_dirs": androidDirs,
	}
	pongo2.Render(this.Ctx, "index.html", context)
}

//
// @Router /api/icon/:app_id/
//
func (this*MainController)AppIcon() {
	appId := this.Ctx.Input.Param(":app_id")
	appsRoot := beego.AppConfig.String("apps_root")
	appRoot := path.Join(appsRoot, appId)
	appIcon := path.Join(appRoot, "app.png")

	var bodyBytes []byte
	var err error
	bodyBytes, err = ioutil.ReadFile(appIcon)
	if err != nil {
		this.Ctx.Output.Status = 404
		return
	}

	output := this.Ctx.Output
	output.Header("Content-Type", "image/png")
	output.Header("Content-Length", fmt.Sprintf("%d", len(bodyBytes)))
	this.Ctx.ResponseWriter.Write(bodyBytes)
}

//
// @Router /api/ipa/:app_id/
//
func (this*MainController)AppIpa() {
	appId := this.Ctx.Input.Param(":app_id")
	appsRoot := beego.AppConfig.String("apps_root")
	appRoot := path.Join(appsRoot, appId)
	appIcon := path.Join(appRoot, "app.ipa")

	var bodyBytes []byte
	var err error
	bodyBytes, err = ioutil.ReadFile(appIcon)
	if err != nil {
		log.Errorf("Error: %v", err)
		this.Ctx.Output.Status = 404
		return
	}

	output := this.Ctx.Output
	output.Header("Content-Disposition", fmt.Sprintf("attachment; filename=%s", url.QueryEscape(fmt.Sprintf("%s.ipa", appId))))
	output.Header("Content-Type", "application/octet-stream")
	output.Header("Content-Transfer-Encoding", "binary")
	output.Header("Content-Length", fmt.Sprintf("%d", len(bodyBytes)))
	this.Ctx.ResponseWriter.Write(bodyBytes)
}

//
// @Router /api/apk/:app_id/
//
func (this*MainController)AndroidApk() {
	appId := this.Ctx.Input.Param(":app_id")
	appsRoot := beego.AppConfig.String("apps_root")
	appRoot := path.Join(appsRoot, appId)
	appIcon := path.Join(appRoot, "app.apk")

	var bodyBytes []byte
	var err error
	bodyBytes, err = ioutil.ReadFile(appIcon)
	if err != nil {
		log.Errorf("Error: %v", err)
		this.Ctx.Output.Status = 404
		return
	}

	output := this.Ctx.Output
	output.Header("Content-Disposition", fmt.Sprintf("attachment; filename=%s", url.QueryEscape(fmt.Sprintf("%s.apk", appId))))
	output.Header("Content-Type", "application/octet-stream")
	output.Header("Content-Transfer-Encoding", "binary")
	output.Header("Content-Length", fmt.Sprintf("%d", len(bodyBytes)))
	this.Ctx.ResponseWriter.Write(bodyBytes)
}



//
// @Title 下载App的plist文件
// @Router /api/plist/:app_id/
//
func (this*MainController)PlistFile() {
	appId := this.Ctx.Input.Param(":app_id")
	appsRoot := beego.AppConfig.String("apps_root")
	appRoot := path.Join(appsRoot, appId)
	pListFile := path.Join(appRoot, "app.plist")

	var bodyBytes []byte
	var err error
	bodyBytes, err = ioutil.ReadFile(pListFile)
	if err != nil {
		log.Errorf("Error: %v", err)
		this.Ctx.Output.Status = 404
		return
	}

	body := string(bodyBytes)

	appHost := beego.AppConfig.String("server_host")
	ipaUrl := fmt.Sprintf("<![CDATA[https://%s/api/ipa/%s]]>", appHost, appId);
	iconUrl := fmt.Sprintf("<![CDATA[https://%s/api/icon/%s]]>", appHost, appId);
	body = strings.Replace(body, "__URL__", ipaUrl, -1)

	index := strings.Index(body, "</array>")

	// appPng := path.Join(appRoot, "app.png")

	imageUrl := fmt.Sprintf(`<dict>
		<key>kind</key>
		<string>display-image</string>
		<key>needs-shine</key>
		<false/>
		<key>url</key>
		<string>%s</string>
	</dict>`, iconUrl)
	body = body[0:index] + imageUrl + body[index:len(body)]

	output := this.Ctx.Output
	output.Header("Content-Type", "content-type: application/xml")
	output.Header("Content-Length", fmt.Sprintf("%d", len(body)))
	this.Ctx.ResponseWriter.Write([]byte(body))
}
//
// @Router /api/mobileprovision/:app_id/
//
func (this*MainController)MobileProvision4Key() {
	appId := this.Ctx.Input.Param(":app_id")

	appsRoot := beego.AppConfig.String("apps_root")
	appRoot := path.Join(appsRoot, appId)
	provision := path.Join(appRoot, "app.mobileprovision")

	var bodyBytes []byte
	var err error
	bodyBytes, err = ioutil.ReadFile(provision)
	if err != nil {
		log.Errorf("Error: %v", err)
		this.Ctx.Output.Status = 404
		return
	}

	output := this.Ctx.Output
	output.Header("Content-Disposition", fmt.Sprintf("attachment; filename=%s", url.QueryEscape(path.Base(provision))))
	output.Header("Content-Type", "application/octet-stream")
	output.Header("Content-Transfer-Encoding", "binary")
	output.Header("Content-Length", fmt.Sprintf("%d", len(bodyBytes)))
	this.Ctx.ResponseWriter.Write(bodyBytes)

}


// https://adhoc.chunyu.me/api/2/apps/com.chunyu.SymptomCheckerOnline?format=mobileprovision