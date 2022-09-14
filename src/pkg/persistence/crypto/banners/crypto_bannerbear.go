package banners

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"genesis_test_case/src/pkg/domain"
	"genesis_test_case/src/pkg/usecase"
	"genesis_test_case/src/pkg/utils"
)

type bannerBearRequest struct {
	Transparent   bool                      `json:"transparent"`
	Template      string                    `json:"template"`
	Metadata      interface{}               `json:"metadata"`
	WebHookUrl    interface{}               `json:"webHook_url"`
	Modifications []bannerBearModifications `json:"modifications"`
}

type bannerBearModifications struct {
	Name       string      `json:"name"`
	Text       string      `json:"text,omitempty"`
	ChartData  []float64   `json:"chart_data,omitempty"`
	Background interface{} `json:"background"`
}

type bannerBearRepo struct {
	bearer      string
	apiEndpoint string
	templateId  string
}

func NewCryptoBannerBearRepository(bearer, enpoint, template string) usecase.CryptoBannerRepository {
	return &bannerBearRepo{
		bearer:      bearer,
		apiEndpoint: enpoint,
		templateId:  template,
	}
}

func (b *bannerBearRepo) GetCryptoBannerUrl(chart []float64, rate *domain.CurrencyRate) (string, error) {
	generationUrl, err := b.getGenerationBannerURL(chart, rate)
	if err != nil {
		return "", err
	}
	return b.waitAndExtractBannerURL(generationUrl)
}

func (b *bannerBearRepo) addBannerBearer(r *http.Request) {
	r.Header.Add("Authorization", "Bearer "+b.bearer)
}

func (b *bannerBearRepo) getGenerationBannerURL(chart []float64, rate *domain.CurrencyRate) (string, error) {
	reqBody, err := b.getBannerRequestBody(chart, rate)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest("POST",
		b.apiEndpoint,
		bytes.NewBuffer(reqBody),
	)
	if err != nil {
		return "", err
	}
	b.addBannerBearer(req)
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

func (b *bannerBearRepo) getBannerRequestBody(chart []float64, rate *domain.CurrencyRate) ([]byte, error) {
	return json.Marshal(&bannerBearRequest{
		Template:    b.templateId,
		Transparent: false,
		WebHookUrl:  nil,
		Metadata:    nil,
		Modifications: []bannerBearModifications{
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
				Text: fmt.Sprintf("1 %s = %v %s",
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

func (b *bannerBearRepo) waitAndExtractBannerURL(imageURL string) (string, error) {
	var (
		jpgUrl string
		err    error
	)

	for jpgUrl == "" {
		time.Sleep(time.Second)
		jpgUrl, err = b.extractBannerURL(imageURL)
		if err != nil {
			log.Println(err.Error())
			return "", err
		}
	}

	return jpgUrl, nil
}

func (b *bannerBearRepo) extractBannerURL(url string) (string, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}

	b.addBannerBearer(req)
	jpg := struct {
		URL string `json:"image_url_jpg"`
	}{}
	err = utils.DoHttpAndParseBody(req, &jpg)
	if err != nil {
		return "", err
	}

	return jpg.URL, nil
}
