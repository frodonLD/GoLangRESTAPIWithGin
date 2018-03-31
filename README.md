# GoLangRESTAPIWithGin

Golang Rest API based on [Gin framework](https://github.com/gin-gonic/gin)

## Use the app

```bash
docker build -t sample/gin-rest-api .
```

## Developper doc

By default gin always run in debug mode. Pass it to release mode:

* using env:   `export GIN_MODE=release`
* using code:  `gin.SetMode(gin.ReleaseMode`

You can find some examples by following this [link](https://github.com/gin-gonic/gin#api-examples)

### Validator

Gin uses [go-playground/validator.v8](https://github.com/go-playground/validator) for validation. See some [examples](https://github.com/gin-gonic/gin#model-binding-and-validation). Check the full docs on tags usage [here](https://godoc.org/gopkg.in/go-playground/validator.v8#hdr-Baked_In_Validators_and_Tags).

### Tests

Tests are obviously possible and we use [testify](https://godoc.org/github.com/stretchr/testify)

Run your tests this way:

```bash
go test ./...
```

Run your benchmark tests this way:

```bash
go test -bench=.
```

### Logging

Logging is handled by [Logrus](https://github.com/sirupsen/logrus) that help us doing it simply.

### Middleware

https://github.com/gin-gonic/gin#custom-middleware

#### Set up

According to the logging documentation, here is the minimal fields we need in our logs:

* timestamp
* correlation_id
* level
* component
* line_number
* message

The option that is be used in this projet is to log in **Json format** which is better than Text format.
Obviously, we choose this because the Json logging format [is supported by fluentd](https://docs.fluentd.org/v1.0/articles/parser_json).

#### Our logs and logrus

Writting logs in logrus is quiet simple. The pre-requiste for us is to init the logger to make him log in JSON.
Writing this code:

```golang
import logrus "github.com/sirupsen/logrus"
// Init your logrus with your formatter, level etc..
logrus.WithFields(logrus.Fields{
    "a_field": "a_value",
}).Warn("A warn log")
```

have this result:

```json
{"a_field":"a_value","level":"warning","msg":"A warn log","time":"2018-02-23T17:51:39+01:00"}
```

#### Good read about entries, hooks and  formatters

* Default entries => [https://github.com/sirupsen/logrus/blob/master/README.md#entries](https://github.com/sirupsen/logrus/blob/master/README.md#entries)
* filename and line => [https://github.com/sirupsen/logrus/issues/63](https://github.com/sirupsen/logrus/issues/63)
* logrus fluentd hook => [https://github.com/evalphobia/logrus_fluent](https://github.com/evalphobia/logrus_fluent)

## Saving external dependencies somewhere on Norauto's repos

In this Golang project, we use a lot of external dependencies and we need to be sure that we're not completely dependant of these exeternal github sources.
In the [Gopkg.toml file](Gopkg.toml) you can find these dependencies to have a look on which dependencie se need to get locally on some Norauto server.

## Other good read

* Logrus and gin => [https://github.com/stephenmuss/ginerus](https://github.com/stephenmuss/ginerus); [https://github.com/gin-gonic/contrib/blob/master/ginrus/ginrus.go](https://github.com/gin-gonic/contrib/blob/master/ginrus/ginrus.go); [https://github.com/gin-gonic/contrib/blob/master/ginrus/example/example.go](https://github.com/gin-gonic/contrib/blob/master/ginrus/example/example.go)
* Gin contrib repo [https://github.com/gin-gonic/contrib](https://github.com/gin-gonic/contrib)
* Others Gin examples [https://medium.com/@thedevsaddam/build-restful-api-service-in-golang-using-gin-gonic-framework-85b1a6e176f3](https://medium.com/@thedevsaddam/build-restful-api-service-in-golang-using-gin-gonic-framework-85b1a6e176f3); [Gin api with Cors](https://gist.github.com/345161974/b62ad14be916aa7d348a1e7f1c02d72f)
