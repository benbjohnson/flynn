package types

type ShutdownInfo struct {
	LastIndex uint64 `json:"last_index"`
}

type RaftLeader struct {
	Host string `json:"host"`
}
