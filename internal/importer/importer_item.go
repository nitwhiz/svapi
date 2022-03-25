package importer

type Item struct {
	ID           string            `json:"id"`
	InternalID   int               `json:"internalId"`
	Category     int               `json:"category"`
	Type         string            `json:"type"`
	DisplayNames map[string]string `json:"displayNames"`
}
