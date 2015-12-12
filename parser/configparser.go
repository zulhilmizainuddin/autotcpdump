package parser

import (
	"os"
	"encoding/json"
)

type ConfigParser struct {
	CommandOptions string
	PcapLocation string
	WiresharkLocation string
}

func (this *ConfigParser) Parse(location string) error {
	configFile, err := os.Open(location)
	if err != nil {
		return err
	}

	jsonParser := json.NewDecoder(configFile)
	if err = jsonParser.Decode(this); err != nil {
		return err
	}

	return nil
}