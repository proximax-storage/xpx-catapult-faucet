package Faucet

import (
	"github.com/json-iterator/go"
	"github.com/proximax-storage/xpx-catapult-faucet/utils"
	"io/ioutil"
	"os"
	"strconv"
)

var (
	json   = jsoniter.ConfigCompatibleWithStandardLibrary
	Config ConfigData
)

func LoadConfig(file *string) (*ConfigData, error) {
	// Open our jsonFile
	jsonFile, err := os.Open(*file)
	// if we os.Open returns an error then handle it
	if err != nil {
		return nil, err
	}
	utils.Logger(0, "Loading config from: ./resources/rest.json")
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	// read our opened xmlFile as a byte array.
	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return nil, err
	}

	// we unmarshal our byteArray which contains our
	// jsonFile's content into 'ConfigData' which we defined above
	err = json.Unmarshal(byteValue, &Config)
	if err != nil {
		return nil, err
	}

	return &Config, nil
}

func (c *ConfigData) FormatServer() string {
	port := strconv.Itoa(c.Server.Port)
	return c.Server.Host + ":" + port
}
