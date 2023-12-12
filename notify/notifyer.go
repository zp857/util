package notify

type Notifyer interface {
	Send(msg Msg) (resp string, err error)
}

type Msg struct {
	Title string
	Body  string
	Group string
}

type Config struct {
	Bark string `json:"bark"`
}
