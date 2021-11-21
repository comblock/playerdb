package playerdb

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
)

type Config struct {
	SteamId         string `json:"steam_id"`
	MinecraftPlayer string `json:"minecraft_player"`
	XboxAccount     string `json:"xbox_account"`
}

func getConfig() Config {
	var config Config

	bytes, err := os.ReadFile("config.json")

	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(bytes, &config)

	if err != nil {
		panic(err)
	}

	return config
}

var config = getConfig()

func TestGetSteamAccount(t *testing.T) {
	data, err := GetSteamAccount(config.SteamId)

	if err != nil {
		t.Error(err)
	}

	fmt.Println(data)
}

func TestGetMinecraftAccount(t *testing.T) {
	data, err := GetMinecraftAccount(config.MinecraftPlayer)

	if err != nil {
		t.Error(err)
	}

	fmt.Println(data)
}

func TestGetXboxAccount(t *testing.T) {
	data, err := GetXboxAccount(config.XboxAccount)

	if err != nil {
		t.Error(err)
	}

	fmt.Println(data)
}
