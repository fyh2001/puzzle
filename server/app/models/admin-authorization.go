package models

type AdminAuthorization struct {
	OptUrl string `json:"optUrl"`
	Secret string `json:"secret"`
}

type AdminAuthorizationResp struct {
	OptUrl string `json:"optUrl"`
	Token  string `json:"token"`
}
