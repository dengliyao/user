package models

import (
	"encoding/json"
	"os"
)

func (u Userinfo) EncoderUsersDb() error {
	file, err := os.Create("users.json")
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	if err := encoder.Encode(UsersDb); err != nil {
		return err
	}
	return nil
}

func (u Userinfo) DecoderUsersDb() error {
	file, err := os.Open("users.json")
	if err != nil {
		return err
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&UsersDb); err != nil {
		return err
	}
	return nil
}
