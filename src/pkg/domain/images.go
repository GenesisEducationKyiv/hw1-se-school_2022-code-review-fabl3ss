package domain

type BannerRequest struct {
	Template      string                `json:"template"`
	Transparent   bool                  `json:"transparent"`
	Modifications []BannerModifications `json:"modifications"`
	WebHookUrl    interface{}           `json:"webHook_url"`
	Metadata      interface{}           `json:"metadata"`
}

type BannerModifications struct {
	Name       string      `json:"name"`
	Background interface{} `json:"background"`
	Text       string      `json:"text,omitempty"`
	ChartData  []float64   `json:"chart_data,omitempty"`
}

type ImageRepository interface {
	GetCryptoBannerUrl(chart []float64, rate *CurrencyRate) (string, error)
}
