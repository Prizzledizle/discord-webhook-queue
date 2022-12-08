package types

type Settings struct {
	Webhooks []struct {
		Webhook string `json:"webhook"`
		Alias   string `json:"alias"`
	} `json:"webhooks"`
	Port int `json:"port"`
}
