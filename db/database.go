package db

import "example/test/model"

var database = map[string]model.Company{
	"golang":  {Name: "golang", Tel: "012-345-6789", Email: "golang-nuts@googlegroups.com"},
	"golang1": {Name: "golang1", Tel: "012-345-6781", Email: "golang1-nuts@googlegroups.com"},
	"golang2": {Name: "golang2", Tel: "012-345-6782", Email: "golang2-nuts@googlegroups.com"},
	"golang3": {Name: "golang3", Tel: "012-345-6783", Email: "golang3-nuts@googlegroups.com"},
	"golang4": {Name: "golang4", Tel: "012-345-6784", Email: "golang4-nuts@googlegroups.com"},
}

func FindAll() []model.Company {
	items := make([]model.Company, len(database))
	for _, v := range database {
		items = append(items, v)
	}

	return items
}

func FindBy(key string) (interface{}, bool) {
	com, ok := database[key]

	return com, ok
}

func Save(key string, item model.Company) {
	database[key] = item
}

func Remove(key string) {
	delete(database, key)
}

// func Init() {
// 	database =
// }
