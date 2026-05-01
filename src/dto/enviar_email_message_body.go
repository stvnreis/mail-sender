package dto

import "github.com/stvnreis/mail-sender/src/mapper"

type EnviarEmailMessageBody struct {
	To      string `json:"to"`
	Subject string `json:"subject"`
	Body    string `json:"body"`
}

func FromJson(jsonData []byte) (*EnviarEmailMessageBody, error) {

	var email EnviarEmailMessageBody
	err := mapper.FromJson(jsonData, &email)

	if err != nil {
		return nil, err
	}

	return &email, nil
}