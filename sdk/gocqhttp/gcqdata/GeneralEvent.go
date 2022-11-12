package gcqdata

type GeneralEvent struct {
	Time     int    `json:"time"`
	SelfId   int    `json:"self_id"`
	PostType string `json:"post_type"`
}
