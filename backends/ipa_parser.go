package backends

import (
	"archive/zip"
	"errors"
	"github.com/DHowett/go-plist"
	"io"
	"log"
	"strings"
)

//ParseIpa : It parses the given ipa and returns a map from the contents of Info.plist in it
func ParseIpa(name string, bundleFilter string) (map[string]interface{}, error) {
	r, err := zip.OpenReader(name)
	if err != nil {
		log.Println("Error opening ipa/zip ", err.Error())
		return nil, err
	}
	defer r.Close()

	for _, file := range r.File {
		// fmt.Printf("FileName: %s\n", file.FileInfo().Name())

		// 存在多个Info.plist 如何处理呢?
		if file.FileInfo().Name() == "Info.plist" {
			rc, err := file.Open()
			if err != nil {
				log.Println("Error opening Info.plist in zip", err.Error())
				return nil, err
			}
			buf := make([]byte, file.FileInfo().Size())
			_, err = io.ReadFull(rc, buf)
			if err != nil {
				log.Println("Error reading Info.plist", err.Error())
				return nil, err
			}


			var info_map map[string]interface{}
			_, err = plist.Unmarshal(buf, &info_map)
			if err != nil {
				log.Println("Error reading Info.plist", err.Error())
				return nil, err
			}


			if id, ok := info_map["CFBundleIdentifier"].(string); !ok || !strings.Contains(strings.ToLower(id), bundleFilter) {
				continue
			}

			return info_map, nil
		}
	}
	return nil, errors.New("Info.plist not found")
}