package types

type ModuleOptionsMenuFunction func(map[string]interface{}) string
type ModuleInfoMenuFunction func(map[string]interface{}) string
type ModuleSetOptions func([]string, map[string]interface{}) error
type ModuleRun func(map[string]interface{})
type ModuleLoadOptions func(map[string]interface{}, map[string]interface{})
