package models

import "errors"

func GetGun(gunType string) (IGun, error) {
	if gunType == "AK47" {
		return newAk47(), nil
	}
	if gunType == "musket" {
		return newMusket(), nil
	}

	return nil, errors.New("Wrong gun type")
}
