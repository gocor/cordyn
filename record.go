package cordyn

// Record ...
type Record struct {
	PK            string `json:"pk"`
	SK            string `json:"sk"`
	PK2           string `json:"pk2"`
	SK2           string `json:"sk2"`
	RecordType    string `json:"typ"`
	SchemaVersion string `json:"v"`
	ID            string `json:"id"`
	CreateDate    string `json:"create_dt"`
	UpdateDate    string `json:"update_dt"`
}

// CompositeKey ...
func (r Record) CompositeKey() CompositeKey {
	return CompositeKey{
		PK: r.PK,
		SK: r.SK,
	}
}

// SecondaryCompositeKey ...
func (r Record) SecondaryCompositeKey() CompositeKey {
	return CompositeKey{
		PK: r.PK2,
		SK: r.SK2,
	}
}
