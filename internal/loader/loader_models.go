package loader

type Category struct {
	ID           int               `json:"id"`
	DisplayNames map[string]string `json:"displayNames"`
}

type Categories struct {
	Version    string     `json:"version"`
	Categories []Category `json:"categories"`
}

type Npc struct {
	ID             string            `json:"id"`
	TextureName    string            `json:"textureName"`
	DisplayNames   map[string]string `json:"displayNames"`
	BirthdaySeason string            `json:"birthdaySeason"`
	BirthdayDay    int               `json:"birthdayDay"`
}

type Npcs struct {
	Version string `json:"version"`
	Npcs    []Npc  `json:"npcs"`
}

type Item struct {
	ID             string            `json:"id"`
	TextureName    string            `json:"textureName"`
	Category       int               `json:"category"`
	Type           string            `json:"type"`
	DisplayNames   map[string]string `json:"displayNames"`
	IsGiftable     bool              `json:"isGiftable"`
	IsBigCraftable bool              `json:"isBigCraftable"`
}

type Items struct {
	Version string `json:"version"`
	Objects []Item `json:"objects"`
}

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

type RecipeOutput struct {
	ItemIDs []string `json:"itemIds"`
	Amount  int      `json:"amount"`
}

type RecipeIngredient struct {
	ItemID   string `json:"itemId"`
	Quantity int    `json:"quantity"`
}

type Recipe struct {
	Name        string             `json:"name"`
	Output      RecipeOutput       `json:"output"`
	IsCooking   bool               `json:"isCooking"`
	Ingredients []RecipeIngredient `json:"ingredients"`
}

type Recipes struct {
	Version string   `json:"version"`
	Recipes []Recipe `json:"recipes"`
}
