package backends

import (
	"git.chunyu.me/feiwang/appserver/models"
	"os"
	"io/ioutil"
	"path"
	log "git.chunyu.me/golang/cyutils/utils/rolling_log"
	"fmt"
	"encoding/json"
	"github.com/astaxie/beego"
	"sort"
)

var gDirScanned bool = false
var gIosAppDirs[]*models.IosAppDirMeta
var gAndroidAppDirs[]*models.AndroidAppDirMeta

// http://stackoverflow.com/questions/12518876/how-to-check-if-a-file-exists-in-go
func IsExist(filePath string) bool {
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return false
	} else {
		return true
	}
}



//获取指定目录下的所有文件，不进入下一级目录搜索，可以匹配后缀过滤。
func ListAppDir(appRootDir string) (iosAppDirs[]*models.IosAppDirMeta, androidAppDirs []*models.AndroidAppDirMeta, err error) {
	log.Infof("AppRootDir: %s", appRootDir)

	// 如果已经扫描了，则直接返回
	if gDirScanned {
		return gIosAppDirs, gAndroidAppDirs, nil
	} else {

		err = ScanAppRootDir(appRootDir)

		// 返回结果
		if err != nil {
			return nil, nil, err
		} else {
			return gIosAppDirs, gAndroidAppDirs, nil
		}

	}

}

func ScanAppRootDir(appsRootDir string) error {

	log.Infof("%s %s", GreenF("Begin Scan Root Dir"), appsRootDir)
	// 扫描目录

	dir, err := ioutil.ReadDir(appsRootDir)
	if err != nil {
		return err
	}

	iosAppDirs := make([]*models.IosAppDirMeta, 0, 10)
	androidAppDirs := make([]*models.AndroidAppDirMeta, 0, 10)
	appHost := beego.AppConfig.String("server_host")

	apiBase := fmt.Sprintf("https://%s/api", appHost)

	// 遍历所有的目录
	for _, fi := range dir {
		if !fi.IsDir() {
			// 忽略目录
			continue
		}
		appId := fi.Name()
		appDir := path.Join(appsRootDir, appId)
		ipaPath := path.Join(appDir, "app.ipa")
		androidPath := path.Join(appDir, "app.apk")


		// 判断是否为 iOs目录
		if IsExist(ipaPath) {
			appMeta := parseIosAppDir(apiBase, appId, appDir)
			if appMeta != nil {
				iosAppDirs = append(iosAppDirs, appMeta)
			}
		}

		// 判断是否为 Android目录
		if IsExist(androidPath) {
			appMeta := parseAndroidAppDir(apiBase, appId, appDir)
			if appMeta != nil {
				androidAppDirs = append(androidAppDirs, appMeta)
			}
		}

	}

	// 按照Released的时间排序
	sort.Sort(models.AndroidAppDirMets(androidAppDirs))
	sort.Sort(models.IosAppDirMetas(iosAppDirs))

	// 记录扫描的结果
	gIosAppDirs = iosAppDirs
	gAndroidAppDirs = androidAppDirs
	gDirScanned = true
	return nil
}

func parseIosAppDir(apiBase string, appId string, appDir string) *models.IosAppDirMeta {
	ipaPath := path.Join(appDir, "app.ipa")

	if !IsExist(path.Join(appDir, "app.plist")) {
		return nil
	}

	metaInfo, _ := ParseIpa(ipaPath, "chunyu")

	state, _ := os.Stat(ipaPath)
	size := fmt.Sprintf("%.2fM", float32(state.Size() / 1024.0 / 1024.0))



	appMeta := &models.IosAppDirMeta{
		Id: appId,
		Plist: fmt.Sprintf("%s/plist/%s", apiBase, appId),
		MobileProvision: fmt.Sprintf("%s/mp/%s", apiBase, appId),
		AppIcon: fmt.Sprintf("%s/icon/%s", apiBase, appId),

		Name: metaInfo["CFBundleDisplayName"].(string),
		Version: metaInfo["CFBundleShortVersionString"].(string),
		ReleaseDate: state.ModTime().Format("2006-01-02 15:04"),
		Size:size,
	}

	has_provinsion := IsExist(path.Join(appDir, "app.mobileprovision"))
	if !has_provinsion {
		appMeta.MobileProvision = ""
	}

	return appMeta
}

func parseAndroidAppDir(apiBase string, appId string, appDir string) *models.AndroidAppDirMeta {
	apkPath := path.Join(appDir, "app.apk")

	if !IsExist(path.Join(appDir, "app.json")) {
		return nil
	}

	state, _ := os.Stat(apkPath)
	size := fmt.Sprintf("%.2fM", float32(state.Size() / 1024.0 / 1024.0))

	appJsonFile := path.Join(appDir, "app.json")
	data, _ := ioutil.ReadFile(appJsonFile)
	var appJson map[string]interface{} = make(map[string]interface{})
	json.Unmarshal(data, &appJson)

	appMeta := &models.AndroidAppDirMeta{
		Id: appId,
		AppIcon: fmt.Sprintf("%s/icon/%s", apiBase, appId),
		Apk: fmt.Sprintf("%s/apk/%s", apiBase, appId),
		ReleaseDate: state.ModTime().Format("2006-01-02 15:04"),
		Size:size,
		Name: appJson["title"].(string),
		Version: appJson["versionName"].(string),
	}
	return appMeta
}