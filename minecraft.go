package playerdb

import (
	"encoding/json"
	"errors"
)

func GetMinecraftAccount(player string) (SteamAccount, error) {
	var account SteamAccount

	data, err := get("https://playerdb.co/api/player/minecraft/" + player)

	if err != nil {
		return account, err
	}

	err = json.Unmarshal(data, &account)

	if !account.Success {
		return account, errors.New(account.Code)
	}

	return account, err
}

type MinecraftAccount struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Data    struct {
		Player struct {
			Meta struct {
				NameHistory []struct {
					Name        string `json:"name"`
					ChangedToAt int64  `json:"changedToAt,omitempty"`
				} `json:"name_history"`
			} `json:"meta"`
			Username string `json:"username"`
			ID       string `json:"id"`
			RawID    string `json:"raw_id"`
			Avatar   string `json:"avatar"`
		} `json:"player"`
	} `json:"data"`
	Success bool `json:"success"`
	Error   bool `json:"error"`
}
