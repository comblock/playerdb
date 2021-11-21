package playerdb

import (
	"encoding/json"
	"errors"
)

func GetXboxAccount(player string) (XboxAccount, error) {
	var account XboxAccount

	data, err := get("https://playerdb.co/api/player/xbox/" + player)

	if err != nil {
		return account, err
	}

	err = json.Unmarshal(data, &account)

	if !account.Success {
		return account, errors.New(account.Code)
	}

	return account, err
}

type XboxAccount struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Data    struct {
		Player struct {
			ID   string `json:"id"`
			Meta struct {
				Gamerscore       string `json:"gamerscore"`
				AccountTier      string `json:"accountTier"`
				XboxOneRep       string `json:"xboxOneRep"`
				PreferredColor   string `json:"preferredColor"`
				RealName         string `json:"realName"`
				Bio              string `json:"bio"`
				TenureLevel      string `json:"tenureLevel"`
				Watermarks       string `json:"watermarks"`
				Location         string `json:"location"`
				ShowUserAsAvatar string `json:"showUserAsAvatar"`
			} `json:"meta"`
			Username string `json:"username"`
			Avatar   string `json:"avatar"`
		} `json:"player"`
	} `json:"data"`
	Success bool `json:"success"`
	Error   bool `json:"error"`
}
