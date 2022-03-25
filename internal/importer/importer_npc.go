package importer

type Npc struct {
	ID             string            `json:"id"`
	DisplayNames   map[string]string `json:"displayNames"`
	BirthdaySeason string            `json:"birthdaySeason"`
	BirthdayDay    int               `json:"birthdayDay"`
}
