package models

import(
	"time"
	"encoding/json"
)

type Metric struct {
	ID        int `json:"id"`
	DriverId uint `gorm:"index" json:"driver_id,string" sql:"type:integer REFERENCES drivers(id)"`
	MetricName string `gorm:"not null" json:"metric_name"`
	Value string `gorm:"not null" json:"value"`
	Lat float32 `gorm:"not null" json:"lat"`
	Lon float32 `gorm:"not null" json:"lon"`
	Timestamp time.Time `gorm:"type:timestamp;not null" json:"timestamp"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}

func (obj *Metric) UnmarshalJSON(data []byte) error {
	type Alias Metric
	aux := &struct {
		Timestamp int64 `json:"timestamp"`
		*Alias
	}{
		Alias: (*Alias)(obj),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	obj.Timestamp = time.Unix(int64(aux.Timestamp), 0)
	return nil
}
