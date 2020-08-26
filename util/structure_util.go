package util

func CloneStringBoolMap(inputMap map[string]bool) map[string]bool {

	outMap := make(map[string]bool, 0)

	for k, v := range inputMap {
		outMap[k] = v
	}

	return outMap
}
