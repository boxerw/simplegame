package core

import "github.com/boxerw/simplegame/core/foundation"

type Data = foundation.Data

type DataModule struct {
	valueMap map[string]interface{}
}

func (dataModule *DataModule) SetValue(name string, value interface{}) {
	if dataModule.valueMap == nil {
		dataModule.valueMap = map[string]interface{}{}
	}
	dataModule.valueMap[name] = value
}

func (dataModule *DataModule) GetValue(name string) interface{} {
	if dataModule.valueMap == nil {
		return nil
	}
	return dataModule.valueMap[name]
}
