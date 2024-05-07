package mapper

import (
	"encoding/xml"
	"io/fs"
	"log"
	"path/filepath"
	"reflect"
	"remember/config"
	"remember/entity"
	"remember/utils"
)

var MP []Mapper

type Select struct {
	XMLName    xml.Name `xml:"select"`
	Sql        string   `xml:",chardata"`
	Id         string   `xml:"id,attr"`
	ResultType string   `xml:"resultType,attr"`
}
type Mapper struct {
	XMLName   xml.Name `xml:"mapper"`
	NameSpace string   `xml:"namespace,attr"`
	Select    []Select `xml:"select"`
}

func ReadXml() []Mapper {
	files := make([]string, 0)
	err := filepath.Walk(config.Configure.DbConfig.Path, func(path string, info fs.FileInfo, err error) error {
		matched, err := filepath.Match(config.Configure.DbConfig.Matched, info.Name())
		if !info.IsDir() && matched {
			files = append(files, path)
		}
		return err
	})
	if err != nil {
		return nil
	}

	mp := make([]Mapper, 0)
	for _, file := range files {
		f := utils.ReadFile(file)
		decoder := xml.NewDecoder(f)
		//v := new([]Mapper)
		var v Mapper
		err := decoder.Decode(&v)
		if err != nil {
			log.Fatalln(err.Error())
		}
		mp = append(mp, v)
	}

	return mp
}

func GetMapper(mapperName string) map[string]Select {
	mp := make(map[string]Select)
	for _, mapper := range MP {
		if mapper.NameSpace == mapperName {
			sqls := mapper.Select
			for _, sql := range sqls {
				mp[sql.Id] = sql
			}
		}
	}
	return mp
}

func GetType(resultType string) reflect.Type {
	TypeMapping := map[string]reflect.Type{
		"entity.User":           reflect.TypeOf(entity.User{}),
		"entity.Users":          reflect.TypeOf([]entity.User{}),
		"entity.Bill":           reflect.TypeOf(entity.Bill{}),
		"entity.Bills":          reflect.TypeOf([]entity.Bill{}),
		"entity.SystemMessage":  reflect.TypeOf(entity.SystemMessage{}),
		"entity.SystemMessages": reflect.TypeOf([]entity.SystemMessage{}),
	}
	return TypeMapping[resultType]
}

func init() {
	MP = ReadXml()
}
