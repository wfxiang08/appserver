package models

type  IosAppDirMeta  struct {
	Id              string
	Plist           string
	AppIcon         string
	MobileProvision string
	Ipa             string

	Name            string
	Version         string
	Author          string
	ReleaseDate     string
	Size            string
}

type  AndroidAppDirMeta  struct {
	Id          string
	AppIcon     string
	Apk         string

	Name        string
	Version     string
	Author      string
	ReleaseDate string
	Size        string
}

type IosAppDirMetas []*IosAppDirMeta

func (a IosAppDirMetas) Len() int {
	return len(a)
}
func (a IosAppDirMetas) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}
// 按照降序排序
func (a IosAppDirMetas) Less(i, j int) bool {
	return a[j].ReleaseDate < a[i].ReleaseDate
}

type AndroidAppDirMets []*AndroidAppDirMeta

func (a AndroidAppDirMets) Len() int {
	return len(a)
}
func (a AndroidAppDirMets) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}
// 按照降序排序
func (a AndroidAppDirMets) Less(i, j int) bool {
	return a[j].ReleaseDate < a[i].ReleaseDate
}
