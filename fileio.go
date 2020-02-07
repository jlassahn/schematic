package schematic

import (
	"encoding/json"
	"io/ioutil"
)

func LoadSchematic(filename string) (*Schematic, error) {

	jdata, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var schem Schematic
	err = json.Unmarshal(jdata, &schem)
	if err != nil {
		return nil, err
	}

	return &schem, nil
}

func (schem *Schematic) Save(name string) error {

	jdata, err := json.MarshalIndent(schem, "", "\t")
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(name, jdata, 0666)
	return err
}
