package entry

type Entry struct {
	Id        int   `json:"id"`
	Timestamp int64 `json:"timestamp"`
	Value     int   `json:"value"`
}
