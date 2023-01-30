package types

type GoteleportActive struct {
	Username string   `json:"username"`
	Roles    []string `json:"roles"`
	Logins   []string `json:"logins"`
}

type GoteleportCMDStatus struct {
	Active GoteleportActive `json:"active"`
}

type TsshConnection struct {
	User string
	Host string
}
