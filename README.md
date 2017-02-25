### Gett home assignment solution

**Deployed to heroku:** http://gettek.herokuapp.com

**JSON API docs:** http://gettek.herokuapp.com/api/v1/doc


####Load data from json files:

```

go run utils/load_data.go -type=drivers -file=data/drivers.json
go run utils/load_data.go -type=metrics -file=data/metrics.json

```

####Run API tests

```

go test test/api_test.go

```

####Runtime configuration

.env and test/.env
```

DATABASE_URL = postgres://dbuser:password@localhost:5432/gettek_dev
BEEGO_RUNMODE = dev

```
see .env.example