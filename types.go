package main

type Sockets struct {
	Sockets map[string]Socket `json:"items"`
}

type Socket struct {
	ID       string `json:"socket_id"`
	Hostname string `json:"host_name"`
	PortTCP  int    `json:"port_tcp"`
	PathHTTP string `json:"path_http"`
	Muted    bool   `json:"muted"`
}

type Domains struct {
	Domains []Domain `json:"domains"`
}

type Domain struct {
	ID         string `json:"domain_id"`
	FQDN       string `json:"domain_fqdn"`
	Owner      string `json:"domain_owner"`
	Expiration string `json:"expiration_date"`
	Registrar  string `json:"registrar_name"`
	CfZone     string `json:"cf_zone_id"`
}

type Depots struct {
	Depot Depot `json:"items"`
}

type Depot struct {
	Owner string      `json:"owner_name"`
	Items []DepotItem `json:"depot_items"`
}

type DepotItem struct {
	ID       int    `json:"id"`
	Desc     string `json:"desc"`
	Misc     string `json:"misc"`
	Location string `json:"depot"`
}

type Nodes struct {
	Nodes []Node `json:"hosts"`
}

type Node struct {
	NameShort string   `json:"hostname_short"`
	NameFQDN  string   `json:"hostname_fqdn"`
	IPAddress []string `json:"ip_address"`
}

type Users struct {
	Users map[string]User `json:"items"`
}

type User struct {
	Name     string `json:"name"`
	FullName string `json:"full_name"`
	Active   bool   `json:"active"`
}

type News struct {
	News []NewsItem `json:"items"`
}

type NewsItem struct {
	Title   string `json:"title"`
	Perex   string `json:"perex"`
	Link    string `json:"link"`
	Server  string `json:"server"`
	PubDate string `json:"pub_date"`
}
