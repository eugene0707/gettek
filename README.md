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

go test test/apitest.go

```

####Runtime configuration

.env see .env.example
