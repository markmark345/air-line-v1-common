package responses

import (
	"encoding/json"
	"io/ioutil"

	"github.com/markmark345/air-line-v1-common/api/responses/model"
)


func GetLookup(msgKey string) (*model.Lookup, *model.Response) {
	jsonData, err := ioutil.ReadFile("./lookup.json")
	if err != nil {
		res := &model.Response{
			Status: model.Status {
				Code: 500,
				Description: "Internal Server Error",
				Details: model.Details{
					Source: "common.GetLookup",
					Reason: "lookup read file error",
					Error: err,
				},
			},
		}
		return nil, res
	}

	var data map[string]model.Lookup
	if err := json.Unmarshal(jsonData, &data); err != nil {
		res := &model.Response{
			Status: model.Status {
				Code: 500,
				Description: "Internal Server Error",
				Details: model.Details{
					Source: "common.GetLookup",
					Reason: "lookup json unmarshal error",
					Error: err,
				},
			},
		}
		return nil, res
	}

	getKey, ok := data[msgKey]
	if !ok {
		res := &model.Response{
			Status: model.Status {
				Code: 500,
				Description: "Internal Server Error",
				Details: model.Details{
					Source: "common.GetLookup",
					Reason: "lookup json not found key",
				},
			},
		}
		return nil, res
	}

	result := &model.Lookup{
		Key: getKey.Key,
		HTTPCode: getKey.HTTPCode,
		Code: getKey.Code,
		DescEN: getKey.DescEN,
		DescTH: getKey.DescTH,
	}

	return result, nil
}