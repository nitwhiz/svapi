package importer

type Item struct {
	ID           string            `json:"id"`
	Category     int               `json:"category"`
	Type         string            `json:"type"`
	DisplayNames map[string]string `json:"displayNames"`
}

type Items struct {
	Version string `json:"version"`
	Objects []Item `json:"objects"`
}
