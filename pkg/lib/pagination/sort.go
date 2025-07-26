package pagination

func GetSortField(sortKey, defaultField string, fieldMap map[string]string) string {
	if field, ok := fieldMap[sortKey]; ok {
		return field
	}
	return defaultField
}

func GetOrder(order string) string {
	val := "ASC"
	if order == "DESC" {
		val = order
	}
	return val
}
