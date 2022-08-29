package repository

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"genesis_test_case/src/pkg/domain"
	"genesis_test_case/src/pkg/utils"
)

type imageRepository struct{}

func NewImageRepository() domain.ImageRepository {
	return &imageRepository{}
}

func (i *imageRepository) GetCryptoBannerUrl(chart []float64, rate *domain.CurrencyRate) (string, error) {
	generationUrl, err := i.getGenerationBannerURL(chart, rate)
	if err != nil {
		return "", err
	}
	return i.waitAndExtractBannerURL(generationUrl)
}

func (i *imageRepository) addBannerBearer(r *http.Request) {
	r.Header.Add("Authorization", "Bearer "+os.Getenv("BANNER_API_TOKEN"))
}

func (i *imageRepository) getGenerationBannerURL(chart []float64, rate *domain.CurrencyRate) (string, error) {
	reqBody, err := i.getBannerRequestBody(chart, rate)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest("POST",
		os.Getenv("BANNER_API_URL"),
		bytes.NewBuffer(reqBody),
	)
	if err != nil {
		return "", err
	}
	i.addBannerBearer(req)
	req.Header.Set("Content-Type", "application/json")

	generationURL := struct {
		Self string `json:"self"`
	}{}
	err = utils.DoHttpAndParseBody(req, &generationURL)
	if err != nil {
		return "", err
	}

	return generationURL.Self, nil
}

func (i *imageRepository) getBannerRequestBody(chart []float64, rate *domain.CurrencyRate) ([]byte, error) {
	return json.Marshal(&domain.BannerRequest{
		Template:    os.Getenv("CRYPTO_TEMPLATE"),
		Transparent: false,
		WebHookUrl:  nil,
		Metadata:    nil,
		Modifications: []domain.BannerModifications{
			{
				Name:       "bg",
				Background: nil,
			},
			{
				Name: "title",
				Text: fmt.Sprintf("%s/%s",
					rate.BaseCurrency,
					rate.QuoteCurrency,
				),
				Background: nil,
			},
			{
				Name: "subtitle",
				Text: fmt.Sprintf("1 %s = %s %s",
					rate.BaseCurrency,
					rate.Price,
					rate.QuoteCurrency,
				),
				Background: nil,
			},
			{
				Name:      "chart",
				ChartData: chart,
			},
		},
	})
}

func (i *imageRepository) waitAndExtractBannerURL(imageURL string) (string, error) {
	var (
		jpgUrl string
		err    error
	)

	for jpgUrl == "" {
		time.Sleep(time.Second)
		jpgUrl, err = i.extractBannerURL(imageURL)
		if err != nil {
			fmt.Println(err.Error())
			return "", err
		}
	}

	return jpgUrl, nil
}

func (i *imageRepository) extractBannerURL(url string) (string, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}

	i.addBannerBearer(req)
	jpg := struct {
		URL string `json:"image_url_jpg"`
	}{}
	err = utils.DoHttpAndParseBody(req, &jpg)
	if err != nil {
		return "", err
	}

	return jpg.URL, nil
}
