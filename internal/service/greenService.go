package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"newProject/internal/domain"
)

type GreenService interface {
	GetSettings(idInstance string, apiTokenInstance string) (*domain.GetSettingsResponse, error)
	GetStateInstance(idInstance string, apiTokenInstance string) (*domain.GetStateInstanceResponse, error)
	SendMessage(idInstance string, apiTokenInstance string, chatId string, message string) (*domain.SendMessageResponse, error)
	SendFileByUrlResponse(idInstance string, apiTokenInstance string, chatId string, urlFile string) (*domain.SendFileByUrlResponse, error)
}

type greenService struct {
	UrlApi string
}

func (g *greenService) GetSettings(idInstance string, apiTokenInstance string) (*domain.GetSettingsResponse, error) {
	url := fmt.Sprintf("%s/waInstance%s/getSettings/%s", g.UrlApi, idInstance, apiTokenInstance)
	res, err := http.Get(url)
	defer res.Body.Close()
	if err != nil {
		return nil, err
	}

	var getSettingsRes domain.GetSettingsResponse
	if err = json.NewDecoder(res.Body).Decode(&getSettingsRes); err != nil {
		return nil, err
	}
	fmt.Println(getSettingsRes)

	return &getSettingsRes, nil
}

func (g *greenService) GetStateInstance(idInstance string, apiTokenInstance string) (*domain.GetStateInstanceResponse, error) {
	url := fmt.Sprintf("%s/waInstance%s/getStateInstance/%s", g.UrlApi, idInstance, apiTokenInstance)
	res, err := http.Get(url)
	defer res.Body.Close()
	if err != nil {
		return nil, err
	}

	var getStateInstanceRes domain.GetStateInstanceResponse
	if err = json.NewDecoder(res.Body).Decode(&getStateInstanceRes); err != nil {
		return nil, err
	}

	return &getStateInstanceRes, nil
}

func (g *greenService) SendMessage(idInstance string, apiTokenInstance string, chatId string, message string) (*domain.SendMessageResponse, error) {
	url := fmt.Sprintf("%s/waInstance%s/sendMessage/%s", g.UrlApi, idInstance, apiTokenInstance)
	sendMessReq := domain.SendMessageRequest{ChatId: chatId, Message: message}

	jsonData, err := json.Marshal(sendMessReq)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	res, err := client.Do(req)
	defer res.Body.Close()
	if err != nil {
		return nil, err
	}

	var sendMessageRes domain.SendMessageResponse
	if err = json.NewDecoder(res.Body).Decode(&sendMessageRes); err != nil {
		return nil, err
	}

	return &sendMessageRes, nil
}

func (g *greenService) SendFileByUrlResponse(idInstance string, apiTokenInstance string, chatId string, urlFile string) (*domain.SendFileByUrlResponse, error) {
	url := fmt.Sprintf("%s/waInstance%s/sendFileByUrl/%s", g.UrlApi, idInstance, apiTokenInstance)
	sendFileUrl := domain.SendFileByUrlRequest{ChatId: chatId, UrlFile: urlFile, FileName: "image.jpg"}

	jsonData, err := json.Marshal(sendFileUrl)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	res, err := client.Do(req)
	defer res.Body.Close()
	if err != nil {
		return nil, err
	}

	var sendFileUrlRes domain.SendFileByUrlResponse
	if err = json.NewDecoder(res.Body).Decode(&sendFileUrlRes); err != nil {
		return nil, err
	}

	return &sendFileUrlRes, nil
}

func New(urlApi string) GreenService {
	return &greenService{UrlApi: urlApi}
}
