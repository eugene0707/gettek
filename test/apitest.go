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
)

func init() {
	_, file, _, _ := runtime.Caller(1)
	apppath, _ := filepath.Abs(filepath.Dir(filepath.Join(file, ".." + string(filepath.Separator))))
	beego.TestBeegoInit(apppath)
}

func TestGetDrivers(t *testing.T) {
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

func TestPostDrivers(t *testing.T) {
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

func TestGetMetrics(t *testing.T) {
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

func TestPostMetrics(t *testing.T) {
	var jsonRequest = []byte(`{"metric_name":"network.reception_strength","value":"24","lon":34.8358982,"timestamp":1352289160,"lat":32.1075039,"driver_id":"3"}`)
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

