# service-mesh-go

The Go contract for the
[Service Mesh API Specification](https://github.com/Paymentbox-com/service-mesh-api).
It is one package, `mesh`, with no dependencies, imported as
`github.com/Paymentbox-com/service-mesh-go/mesh`. Transports and protocol
layers build against it; nothing in it moves bytes. The
[gRPC Service Mesh API](https://github.com/Paymentbox-com/grpc-service-mesh-api)
is the protocol layer that generates code against this contract from protobuf
definitions, through its Go library
[grpc-service-mesh-go](https://github.com/Paymentbox-com/grpc-service-mesh-go).

The package fixes:

- the types `Target`, `ServiceMap`, `Message`, `Endpoint`, and `Subscriber`
- the handler signatures `EndpointHandler` and `SubscriberHandler`
- the `Client` and `Runtime` interfaces
- `Config` and the two keys the specification defines, `deployment_group`
  and `consumer_group`, with `ConsumerGroupNone`
- the three errors `ErrKindMismatch`, `ErrInvalidTarget`, and
  `ErrNoDeploymentGroup`

## Transports

A transport is a separate module that implements `mesh.Runtime` and
`mesh.Client` and exports its own `New` and `NewClient`. `NewClient` takes
the config and the transport's `ServiceMap`, which the client holds and
returns from `ServiceMap()`.

- NATS: [service-mesh-nats-go](https://github.com/Paymentbox-com/service-mesh-nats-go)

## Usage

Code written against this package works with any transport. A function that
takes a `mesh.Client` and a `mesh.Target` does not know or care which one it
was handed.

```go
import "github.com/Paymentbox-com/service-mesh-go/mesh"

func lookup(ctx context.Context, c mesh.Client, target mesh.Target, id string) ([]byte, error) {
    reply, err := c.Request(ctx, mesh.Message{
        Target:   target,
        Metadata: map[string]string{"Request-Id": id},
        Payload:  []byte(id),
    }, nil)
    if err != nil {
        return nil, err
    }
    return reply.Payload, nil
}
```

The caller constructs a runtime from whichever transport module it uses and
passes `rt.Client()`.

## Development

Tool versions are pinned in `mise.toml`; `mise install` installs them. `just`
lists the recipes. `just check` runs the same format, vet, test,
vulnerability, and lint steps as CI.

## Tests

```
go test -race ./...
```
