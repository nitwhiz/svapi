package flags

import "fmt"

var currentIndex = byte(0)

type Flag struct {
	index byte
	Name  string `json:"name"`
}

func (f *Flag) String() string {
	return f.Name
}

func (f *Flag) MarshalJSON() ([]byte, error) {
	return []byte("\"" + f.Name + "\""), nil
}

var flagMap = map[string]*Flag{}

func register(name string) *Flag {
	f := &Flag{currentIndex, name}
	flagMap[name] = f
	currentIndex += 1

	return f
}

func Get(name string) *Flag {
	f, ok := flagMap[name]

	if !ok {
		return nil
	}

	return f
}

func (f *Flag) Val() byte {
	return 1 << f.index
}

func Count() int {
	return len(flagMap)
}

var IsGiftable = register("giftable")

var IsBigCraftable = register("bigCraftable")

var IsCooking = register("cooking")

func AppendToIndex(index []string, fs []*Flag) []string {
	flagVal := byte(0)

	for _, f := range fs {
		flagVal |= f.Val()
	}

	for b := range Count() {
		if flagVal&(1<<b) != 0 {
			index = append(index, fmt.Sprintf("%d", 1<<b))
		} else {
			index = append(index, "")
		}
	}

	return index
}
