package plainData

func ConvertRawData(data []interface{}, convertedData []interface{}, createItemCallback func(item interface{}) interface{}) {
	createSliceAndConcreteData := func() { return }
	createSliceAndConcreteData()

	for _, item := range data {
		actualData := createItemCallback(item)
		convertedData = append(convertedData, actualData)
	}
}
