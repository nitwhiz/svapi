package importer

type GiftTasteByNpc struct {
	NpcID        string   `json:"npcId"`
	DislikeItems []string `json:"dislikeItems"`
	HateItems    []string `json:"hateItems"`
	LikeItems    []string `json:"likeItems"`
	LoveItems    []string `json:"loveItems"`
	NeutralItems []string `json:"neutralItems"`
}

type GiftTastes struct {
	Version     string           `json:"version"`
	TastesByNpc []GiftTasteByNpc `json:"tastesByNpc"`
}
