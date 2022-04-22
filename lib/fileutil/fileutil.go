package fileutil

import "strings"

func RemoveFileExtension(fileName string) string {
	return strings.SplitN(fileName, ".", 2)[0]
}

func RemoveFileNameFromDirectory(dir string) string {
	return strings.SplitN(dir, "_", 2)[0][4:]
}
