package models

import(
//	"encoding/base64"
	"time"
	"encoding/json"
	"github.com/eugene0707/gettek/common"
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

func LoadMetricsFromJSONArray(JSONArray string) (r int64, e error) {
	var err error
	var result int64

	tx := common.CurrentDB().Begin()

	if err = tx.Exec("CREATE TEMPORARY TABLE temp_json (data text) ON COMMIT DROP;").Error; err != nil { return 0, err }

	if err = tx.Exec("INSERT INTO temp_json (data) values (?);", JSONArray).Error; err != nil { return 0, err }

	if res := tx.Exec(`
		WITH jsdata AS (
		SELECT data->>'metric_name' AS metric_name,
		       data->>'value' AS value,
		       (data->>'lon')::numeric AS lon,
		       (data->>'lat')::numeric AS lat,
		       to_timestamp ((data->>'timestamp')::int) AS "timestamp",
		       (data->>'driver_id')::int AS driver_id,
		       current_timestamp AS created_at,
		       current_timestamp AS updated_at
		from   (
			   SELECT json_array_elements(data::json) AS data
			   FROM   temp_json
		       ) arr
		)
		INSERT INTO metrics (metric_name, value, lon, lat, "timestamp", driver_id, created_at, updated_at)
		SELECT * FROM jsdata WHERE driver_id IN (SELECT id FROM drivers);
	`); res.Error != nil {
		return 0, res.Error
	} else {
		result = res.RowsAffected
	}

	return result, tx.Commit().Error
}
