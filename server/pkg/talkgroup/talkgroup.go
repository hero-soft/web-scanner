package talkgroup

import (
	"fmt"
	"os"

	"github.com/gocarina/gocsv"
	"github.com/spf13/viper"
)

type Talkgroup struct {
	ID          string   `json:"id,omitempty"`
	Name        string   `json:"name,omitempty"`
	Description string   `json:"description,omitempty"`
	Tags        []string `json:"tags,omitempty"`
	Categories  []string `json:"categories,omitempty"`
}

type record struct { // Our example struct, you can use "-" to ignore a field
	Decimal     string `csv:"Decimal"`
	Hex         string `csv:"Hex"`
	AlphaTag    string `csv:"Alpha Tag"`
	Mode        string `csv:"Mode"`
	Description string `csv:"Description"`
	Tag         string `csv:"Tag"`
	Category    string `csv:"Category"`
}

func GetAll() ([]*Talkgroup, error) {
	talkgroupsFile, err := os.OpenFile(viper.GetString("server.talkgroups_file"), os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		return nil, fmt.Errorf("could not open talkgroups file: %v", err)
	}
	defer talkgroupsFile.Close()

	talkgroups := []*record{}

	if err := gocsv.UnmarshalFile(talkgroupsFile, &talkgroups); err != nil { // Load clients from file
		return nil, fmt.Errorf("could not unmarshal talkgroups file: %v", err)
	}

	var selectedTGs []*Talkgroup

	for _, tg := range talkgroups {
		selectedTGs = append(selectedTGs, &Talkgroup{
			ID:          tg.Decimal,
			Name:        tg.AlphaTag,
			Description: tg.Description,
			// Categories:  tg.Category,
		})
	}

	return selectedTGs, nil

}

func Lookup(talkgroupID string, fallback string) (Talkgroup, error) {
	talkgroupsFile, err := os.OpenFile(viper.GetString("server.talkgroups_file"), os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		return Talkgroup{ID: talkgroupID, Name: fallback}, fmt.Errorf("could not open talkgroups file: %v", err)
	}
	defer talkgroupsFile.Close()

	talkgroups := []*record{}

	if err := gocsv.UnmarshalFile(talkgroupsFile, &talkgroups); err != nil { // Load clients from file
		return Talkgroup{ID: talkgroupID, Name: fallback}, fmt.Errorf("could not unmarshal talkgroups file: %v", err)
	}

	selectedTG := Talkgroup{
		ID:   talkgroupID,
		Name: fallback,
	}

	for _, tg := range talkgroups {
		if tg.Decimal == talkgroupID {
			selectedTG.Name = tg.AlphaTag
			selectedTG.Description = tg.Description
			// selectedTG.Categories = tg.Category
			break
		}
	}

	return selectedTG, nil

}
