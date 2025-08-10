package api

import (
	"os"

	"gitee.com/quant1x/gox/util/homedir"
	"gitee.com/quant1x/pkg/gocsv"
)

const (
	DefaultTagName = "dataframe"
)

func init() {
	gocsv.TagName = DefaultTagName
}

// CsvToSlices CSV文件转struct切片
func CsvToSlices[S ~[]E, E any](filename string, pointer *S) error {
	filepath, err := homedir.Expand(filename)
	if err != nil {
		return err
	}
	csvFile, err := os.Open(filepath)
	if err != nil {
		return err
	}
	err = gocsv.Unmarshal(csvFile, pointer)
	CloseQuietly(csvFile)
	return err
}

// SlicesToCsv struct切片转csv文件
func SlicesToCsv[S ~[]E, E any](filename string, s S, force ...bool) error {
	filepath, err := homedir.Expand(filename)
	if err != nil {
		return err
	}
	// 检查目录, 不存在就创建
	_ = CheckFilepath(filepath, true)
	csvFile, err := os.Create(filepath)
	if err != nil {
		return err
	}
	err = gocsv.MarshalFile(s, csvFile)
	if err == nil {
		forceSync := false
		if len(force) > 0 && force[0] {
			forceSync = true
		}
		// 强制刷新内存副本到磁盘
		if forceSync {
			err = csvFile.Sync()
		}
	}
	CloseQuietly(csvFile)
	return err
}
