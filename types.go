package main

type Sockets struct {
	Sockets map[string]Socket `json:"sockets"`
}

type Socket struct {
	ID       string `json:"socket_id"`
	Hostname string `json:"host_name"`
	PortTCP  int    `json:"port_tcp"`
	PathHTTP string `json:"path_http"`
	Muted    bool   `json:"muted"`
}

type Nodes struct {
	Nodes []Node `json:"hosts"`
}

type Node struct {
	NameShort string   `json:"hostname_short"`
	IPAddress []string `json:"ip_address"`
}

type Users struct {
	Users map[string]User `json:"users"`
}

type User struct {
	Name     string `json:"name"`
	FullName string `json:"full_name"`
	Active   bool   `json:"active"`
}

type News struct {
	News []NewsItem `json:"news"`
}

type NewsItem struct {
	Title   string `json:"title"`
	Perex   string `json:"perex"`
	Link    string `json:"link"`
	PubDate string `json:"pub_date"`
}
