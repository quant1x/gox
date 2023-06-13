package api

import (
	"gitee.com/quant1x/gox/util/homedir"
	"github.com/gocarina/gocsv"
	"os"
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
	return err
}

// SlicesToCsv struct切片转csv文件
func SlicesToCsv[S ~[]E, E any](filename string, s S) error {
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
	defer CloseQuietly(csvFile)
	err = gocsv.MarshalFile(s, csvFile)
	return err
}
