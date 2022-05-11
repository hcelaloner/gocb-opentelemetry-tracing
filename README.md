# gocb-opentelemetry-tracing

This repo is actually a fork of [couchbase/gocb-opentelemetry](https://github.com/couchbase/gocb-opentelemetry). When I tried to used it, I was not able to use it because of the version conflicts in  OpenTelemetry's alpha Metrics API. That's why I removed components related with metrics, and this repo only contains tracing bits. I am planning to provide metrics in another module/repo in the future.  Many thanks to original project's contributors.

### Built With

* [Golang](https://go.dev/)
* [GoReleaser](https://goreleaser.com/)


### Usage

```go
import (
	"github.com/hcelaloner/gocb-opentelemetry-tracing"
)


cluster, err := gocb.Connect(server, gocb.ClusterOptions{
    Authenticator: gocb.PasswordAuthenticator{
        Username: user,
        Password: password,
    },
    Tracer: NewOpenTelemetryRequestTracer(otel.GetTracerProvider()),
})
if err != nil {
    panic(err)
}
defer cluster.Close(nil)

b := cluster.Bucket(bucket)
err = b.WaitUntilReady(5*time.Second, nil)
if err != nil {
    panic(err)
}

_, err = col.Upsert("someid", "someval", &gocb.UpsertOptions{
    ParentSpan: NewOpenTelemetryRequestSpanFromContext(ctx),
})
if err != nil {
    panic(err)
}
```