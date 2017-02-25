package test

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"runtime"
	"path/filepath"
	_ "github.com/eugene0707/gettek/routers"
	"github.com/astaxie/beego"
	. "github.com/smartystreets/goconvey/convey"
	"bytes"
	"github.com/eugene0707/gettek/common"
	"github.com/eugene0707/gettek/models"
	"time"
)

func init() {
	_, file, _, _ := runtime.Caller(1)
	apppath, _ := filepath.Abs(filepath.Dir(filepath.Join(file, ".." + string(filepath.Separator))))
	beego.TestBeegoInit(apppath)
}

func setupTestCase(t *testing.T) func(t *testing.T) {
	common.CurrentDB().AutoMigrate(&models.Driver{}, &models.Metric{})

	driver1 := models.Driver{Name: "John Doe", LicenseNumber: "00-000-01"}
	common.CurrentDB().Create(&driver1)

	metric1 := models.Metric{MetricName: "network.reception_strength", Value: "24", Lon:34.8358982, Lat:32.1075039,
		Timestamp: time.Unix(1352289164, 0), DriverId: uint(driver1.ID)}
	common.CurrentDB().Create(&metric1)

	metric2 := models.Metric{MetricName: "network.api_failed.meter_report", Value: "1", Lon:0.009253333333333334, Lat:51.55745666666667,
		Timestamp: time.Unix(1352488148, 0), DriverId: uint(driver1.ID)}
	common.CurrentDB().Create(&metric2)

	driver2 := models.Driver{Name: "Adam Smith", LicenseNumber: "00-000-02"}
	common.CurrentDB().Create(&driver2)

	return func(t *testing.T) {
		common.CurrentDB().Exec("DROP TABLE metrics;")
		common.CurrentDB().Exec("DROP TABLE drivers;")
	}
}

func TestGetDrivers(t *testing.T) {
	teardownTestCase := setupTestCase(t)
	defer teardownTestCase(t)

	r, _ := http.NewRequest("GET", "/api/v1/drivers", nil)
	w := httptest.NewRecorder()
	beego.BeeApp.Handlers.ServeHTTP(w, r)

	beego.Trace("testing", "TestGetDrivers", "Code[%d]\n%s", w.Code, w.Body.String())

	Convey("Subject: Test GetALL Drivers\n", t, func() {
		Convey("Status Code Should Be 200", func() {
			So(w.Code, ShouldEqual, 200)
		})
		Convey("The Result Should Not Be Empty", func() {
			So(w.Body.Len(), ShouldBeGreaterThan, 0)
		})
	})
}

func TestGetOneDriver(t *testing.T) {
	teardownTestCase := setupTestCase(t)
	defer teardownTestCase(t)

	r, _ := http.NewRequest("GET", "/api/v1/drivers/1", nil)
	w := httptest.NewRecorder()
	beego.BeeApp.Handlers.ServeHTTP(w, r)

	beego.Trace("testing", "TestGetOneDrivers", "Code[%d]\n%s", w.Code, w.Body.String())

	Convey("Subject: Test GetOne Driver\n", t, func() {
		Convey("Status Code Should Be 200", func() {
			So(w.Code, ShouldEqual, 200)
		})
		Convey("The Result Should Not Be Empty", func() {
			So(w.Body.Len(), ShouldBeGreaterThan, 0)
		})
	})
}

func TestPostDrivers(t *testing.T) {
	teardownTestCase := setupTestCase(t)
	defer teardownTestCase(t)

	var jsonRequest = []byte(`{"name":"My name", "license_number":"00-111-22"}`)
	r, _ := http.NewRequest("POST", "/api/v1/drivers",  bytes.NewBuffer(jsonRequest))
	w := httptest.NewRecorder()
	beego.BeeApp.Handlers.ServeHTTP(w, r)

	beego.Trace("testing", "TestPostDrivers", "Code[%d]\n%s", w.Code, w.Body.String())

	Convey("Subject: Test Post Driver\n", t, func() {
		Convey("Status Code Should Be 201", func() {
			So(w.Code, ShouldEqual, 201)
		})
		Convey("The Result Should Not Be Empty", func() {
			So(w.Body.Len(), ShouldBeGreaterThan, 0)
		})
	})
}

func TestPutDriver(t *testing.T) {
	teardownTestCase := setupTestCase(t)
	defer teardownTestCase(t)

	var jsonRequest = []byte(`{"name":"New driver", "license_number":"01-234-56"}`)
	r, _ := http.NewRequest("PUT", "/api/v1/drivers/1",  bytes.NewBuffer(jsonRequest))
	w := httptest.NewRecorder()
	beego.BeeApp.Handlers.ServeHTTP(w, r)

	beego.Trace("testing", "TestPutDriver", "Code[%d]\n%s", w.Code, w.Body.String())

	Convey("Subject: Test Put Driver\n", t, func() {
		Convey("Status Code Should Be 200", func() {
			So(w.Code, ShouldEqual, 200)
		})
		Convey("The Result Should Not Be Empty", func() {
			So(w.Body.Len(), ShouldBeGreaterThan, 0)
		})
	})
}

func TestDeleteDriver(t *testing.T) {
	teardownTestCase := setupTestCase(t)
	defer teardownTestCase(t)

	r, _ := http.NewRequest("DELETE", "/api/v1/drivers/2",  nil)
	w := httptest.NewRecorder()
	beego.BeeApp.Handlers.ServeHTTP(w, r)

	beego.Trace("testing", "TestDeleteDriver", "Code[%d]\n%s", w.Code, w.Body.String())

	Convey("Subject: Test Delete Driver\n", t, func() {
		Convey("Status Code Should Be 204", func() {
			So(w.Code, ShouldEqual, 204)
		})
		Convey("The Result Should Not Be Empty", func() {
			So(w.Body.Len(), ShouldBeGreaterThan, 0)
		})
	})
}

func TestGetMetrics(t *testing.T) {
	teardownTestCase := setupTestCase(t)
	defer teardownTestCase(t)

	r, _ := http.NewRequest("GET", "/api/v1/metrics", nil)
	w := httptest.NewRecorder()
	beego.BeeApp.Handlers.ServeHTTP(w, r)

	beego.Trace("testing", "TestGetMetrics", "Code[%d]\n%s", w.Code, w.Body.String())

	Convey("Subject: Test GetALL Metrics\n", t, func() {
		Convey("Status Code Should Be 200", func() {
			So(w.Code, ShouldEqual, 200)
		})
		Convey("The Result Should Not Be Empty", func() {
			So(w.Body.Len(), ShouldBeGreaterThan, 0)
		})
	})
}

func TestGetOneMetric(t *testing.T) {
	teardownTestCase := setupTestCase(t)
	defer teardownTestCase(t)

	r, _ := http.NewRequest("GET", "/api/v1/metrics/1", nil)
	w := httptest.NewRecorder()
	beego.BeeApp.Handlers.ServeHTTP(w, r)

	beego.Trace("testing", "TestGetOneMetric", "Code[%d]\n%s", w.Code, w.Body.String())

	Convey("Subject: Test GetOne Metric\n", t, func() {
		Convey("Status Code Should Be 200", func() {
			So(w.Code, ShouldEqual, 200)
		})
		Convey("The Result Should Not Be Empty", func() {
			So(w.Body.Len(), ShouldBeGreaterThan, 0)
		})
	})
}

func TestPostMetrics(t *testing.T) {
	teardownTestCase := setupTestCase(t)
	defer teardownTestCase(t)

	var jsonRequest = []byte(`{"metric_name":"network.reception_strength","value":"24","lon":34.8358982,"timestamp":1352289160,"lat":32.1075039,"driver_id":"1"}`)
	r, _ := http.NewRequest("POST", "/api/v1/metrics",  bytes.NewBuffer(jsonRequest))
	w := httptest.NewRecorder()
	beego.BeeApp.Handlers.ServeHTTP(w, r)

	beego.Trace("testing", "TestPostMetrics", "Code[%d]\n%s", w.Code, w.Body.String())

	Convey("Subject: Test Post Metric\n", t, func() {
		Convey("Status Code Should Be 201", func() {
			So(w.Code, ShouldEqual, 201)
		})
		Convey("The Result Should Not Be Empty", func() {
			So(w.Body.Len(), ShouldBeGreaterThan, 0)
		})
	})
}

func TestPutMetric(t *testing.T) {
	teardownTestCase := setupTestCase(t)
	defer teardownTestCase(t)

	var jsonRequest = []byte(`{"metric_name":"network.reception_strength","value":"24","lon":34.8358982,"timestamp":1352289160,"lat":32.1075039,"driver_id":"1"}`)
	r, _ := http.NewRequest("PUT", "/api/v1/metrics/2",  bytes.NewBuffer(jsonRequest))
	w := httptest.NewRecorder()
	beego.BeeApp.Handlers.ServeHTTP(w, r)

	beego.Trace("testing", "TestPutMetric", "Code[%d]\n%s", w.Code, w.Body.String())

	Convey("Subject: Test Put Metric\n", t, func() {
		Convey("Status Code Should Be 200", func() {
			So(w.Code, ShouldEqual, 200)
		})
		Convey("The Result Should Not Be Empty", func() {
			So(w.Body.Len(), ShouldBeGreaterThan, 0)
		})
	})
}

func TestDeleteMetric(t *testing.T) {
	teardownTestCase := setupTestCase(t)
	defer teardownTestCase(t)

	r, _ := http.NewRequest("DELETE", "/api/v1/metrics/1",  nil)
	w := httptest.NewRecorder()
	beego.BeeApp.Handlers.ServeHTTP(w, r)

	beego.Trace("testing", "TestDeleteMetric", "Code[%d]\n%s", w.Code, w.Body.String())

	Convey("Subject: Test Delete Metric\n", t, func() {
		Convey("Status Code Should Be 204", func() {
			So(w.Code, ShouldEqual, 204)
		})
		Convey("The Result Should Not Be Empty", func() {
			So(w.Body.Len(), ShouldBeGreaterThan, 0)
		})
	})
}

