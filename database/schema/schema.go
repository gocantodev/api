package schema

import (
	"github.com/gocantodev/server/model"
)

type Schema struct {
	models *map[string]interface{}
}

func Make() Schema {
	schema := Schema{}
	models := make(map[string]interface{}, 2)

	models["author"] = model.Author{}
	models["post"] = model.Post{}

	schema.models = &models

	return schema
}

func (receiver Schema) GetModels() []interface{} {
	models := *receiver.models
	abstracts := make([]interface{}, 0, len(models))

	for _, value := range *receiver.models {
		abstracts = append(abstracts, value)
	}

	return abstracts
}
