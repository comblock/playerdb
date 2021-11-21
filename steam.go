package playerdb

import (
	"encoding/json"
	"errors"
)

func GetSteamAccount(id string) (SteamAccount, error) {
	var account SteamAccount

	data, err := get("https://playerdb.co/api/player/steam/" + id)

	if err != nil {
		return account, err
	}

	err = json.Unmarshal(data, &account)

	if !account.Success {
		return account, errors.New(account.Code)
	}

	return account, err
}

type SteamAccount struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Data    struct {
		Player struct {
			Meta struct {
				Steam2ID                 string `json:"steam2id"`
				Steam2IDNew              string `json:"steam2id_new"`
				Steam3ID                 string `json:"steam3id"`
				Steam64ID                string `json:"steam64id"`
				Steamid                  string `json:"steamid"`
				Communityvisibilitystate int    `json:"communityvisibilitystate"`
				Profilestate             int    `json:"profilestate"`
				Personaname              string `json:"personaname"`
				Profileurl               string `json:"profileurl"`
				Avatar                   string `json:"avatar"`
				Avatarmedium             string `json:"avatarmedium"`
				Avatarfull               string `json:"avatarfull"`
				Avatarhash               string `json:"avatarhash"`
				Personastate             int    `json:"personastate"`
				Primaryclanid            string `json:"primaryclanid"`
				Timecreated              int    `json:"timecreated"`
				Personastateflags        int    `json:"personastateflags"`
			} `json:"meta"`
			ID       string `json:"id"`
			Avatar   string `json:"avatar"`
			Username string `json:"username"`
		} `json:"player"`
	} `json:"data"`
	Success bool `json:"success"`
	Error   bool `json:"error"`
}
