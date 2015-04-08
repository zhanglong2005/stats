package entry

type EntryByTimestamp []Entry

func (v EntryByTimestamp) Len() int { return len(v) }
func (v EntryByTimestamp) Swap(i, j int) { v[i], v[j] = v[j], v[i] }
func (v EntryByTimestamp) Less(i, j int) bool { return v[i].Timestamp > v[j].Timestamp }


type EntryByTimestampDesc []Entry

func (v EntryByTimestampDesc) Len() int { return len(v) }
func (v EntryByTimestampDesc) Swap(i, j int) { v[i], v[j] = v[j], v[i] }
func (v EntryByTimestampDesc) Less(i, j int) bool { return v[i].Timestamp < v[j].Timestamp }


