package models

type SessionStore struct {
	Size     int    `json:"Server"`
	NetWork  string `json:"NetWork"`
	Host     string `json:"Host"`
	Port     int    `json:"Port"`
	DB       int    `json:"DB"`
	Password string `json:"Password"`
}

type Server struct {
	Platform                string       `json:"Server"`
	Host                    string       `json:"Host"`
	Port                    int          `json:"Port"`
	Version                 string       `json:"Version"`
	SESSION_COOKIE_DOMAIN   string       `json:"SESSION_COOKIE_DOMAIN"`
	SESSION_COOKIE_SIGN_KEY string       `json:"SESSION_COOKIE_SIGN_KEY"`
	SESSION_TIMEOUT         int          `json:"SESSION_TIMEOUT"`
	MaxAge                  int          `json:"MaxAge"`
	Static                  string       `json:"Static"`
	MainPage                string       `json:"MainPage"`
	SessionStore            SessionStore `json:"SessionStore"`
}

type Etcd struct {
	Addrs string
}

type MicroConfig struct {
	Name    string
	Version string
	Etcd    Etcd
}

type MicroServices struct {
	MicroMongo MicroConfig
	MicroApi   MicroConfig
}

type BasicAuth struct {
	Accounts map[string]string `json:"Accounts"`
	Version  string            `json:"Version"`
	Prefix   string            `json:"Prefix"`
}

type Config struct {
	Server        Server `json:"Server"`
	MicroServices MicroServices
	BasicAuth     BasicAuth `json:"BasicAuth"`
}

type UserProfile struct {
	ID     string
	Name   string
	mobile string
}
