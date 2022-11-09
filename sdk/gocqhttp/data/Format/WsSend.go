package Format

type WsSend struct {
	Action string   `json:"action"`
	Params struct{} `json:"params"`
	Echo   string   `json:":echo"`
}
