package todo

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
)

type Item struct {
	Text     string
	Priority int
	position int
	Done     bool
}

type ByPri []Item

func (s ByPri) Len() int      { return len(s) }
func (s ByPri) Swap(i, j int) { s[i], s[j] = s[j], s[i] }
func (s ByPri) Less(i, j int) bool {
	if s[i].Done != s[j].Done {
		return s[j].Done
	}
	if s[i].Priority == s[j].Priority {
		return s[i].position < s[j].position
	}
	return s[i].Priority < s[j].Priority
}

func (i *Item) SetPriority(pri int) {
	if pri < 1 {
		i.Priority = 1
	} else if pri > 255 {
		i.Priority = 255
	} else {
		i.Priority = pri
	}
}

func (i *Item) PrettyPrintPriority() string {
	return fmt.Sprintf("( %03d )", i.Priority)
}

func (i *Item) PrettyPrintDone() string {
	if i.Done {
		return "X"
	}
	return ""
}

func (i *Item) Label() string {
	return strconv.Itoa(i.position) + "."
}

func ReadItems(filename string) ([]Item, error) {
	b, err := os.ReadFile(filename)
	if err != nil {
		return []Item{}, err
	}

	var items []Item
	if err := json.Unmarshal(b, &items); err != nil {
		return []Item{}, err
	}

	for i, _ := range items {
		items[i].position = i + 1
	}

	return items, nil
}

func SaveItems(filename string, items []Item) error {
	b, err := json.Marshal(items)
	if err != nil {
		return err
	}

	f, _ := os.Create(filename)
	defer f.Close()

	_, err = f.Write(b)
	if err != nil {
		return err
	}
	return nil
}
