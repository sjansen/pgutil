package tasks

type Exec struct {
	Args []string `json:"exec"`
}

type SQL struct {
	SQL string `json:"sql"`
}
