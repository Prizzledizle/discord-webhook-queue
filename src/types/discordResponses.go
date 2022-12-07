package types

type RatelimitResponse struct {
	Code        int     `json:"code"`
	Global      bool    `json:"global"`
	Message     string  `json:"message"`
	Retry_after float64 `json:"retry_after"`
}
