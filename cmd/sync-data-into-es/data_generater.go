package main

type domainFinalData struct {
	Host   string `json:"host"`
	Zone   string `json:"zone"`
	Domain string `json:"domain"`
	Memo   string `json:"memo"`
}

func getFinalData() ([]*domainFinalData, error) {
	return []*domainFinalData{}, nil
}
