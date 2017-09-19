package fbm_api

//go:generate ffjson $GOFILE

type WhitelistedDomainsMessage struct {
	WhitelistedDomains []string `json:"whitelisted_domains"`
}
