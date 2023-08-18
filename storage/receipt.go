package storage

var receiptStorage map[string]int

func init() {
	receiptStorage = make(map[string]int)
}

func Save(id string, points int) {
	receiptStorage[id] = points
}

func IsContains(id string) bool {
	_, ok := receiptStorage[id]
	return ok
}

func GetPoints(id string) int {
	return receiptStorage[id]
}
