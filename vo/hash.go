package vo

type CreateHashPost struct {
	Maxsize int `json:"maxsize" binding:required`
}

type RUDHashPost struct {
	Key int `json:"key" binding:required`
}
