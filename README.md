# service-mesh-go

The Go contract for the
[Service Mesh API Specification](https://github.com/Paymentbox-com/service-mesh-api),
packaged as the module `github.com/Paymentbox-com/service-mesh-go` with one
package, `mesh`. It holds what every transport and every caller must agree on,
and nothing that moves bytes. Transports are separate modules that depend on it
and implement `Client` and `Runtime`.

The [gRPC Service Mesh API](https://github.com/Paymentbox-com/grpc-service-mesh-api) is a protocol layer that generates code against this contract from protobuf
definitions, through its Go library [grpc-service-mesh-go](https://github.com/Paymentbox-com/grpc-service-mesh-go). Other protocol layers may be implemented
to do the same.

## Install

```sh
go get github.com/Paymentbox-com/service-mesh-go
```

```go
import "github.com/Paymentbox-com/service-mesh-go/mesh"
```

Requires Go 1.26 or newer. The module has no dependencies.

## What it Implements

- The value types from the specification: `mesh.Target`, `mesh.ServiceMap`,
  `mesh.Message`, `mesh.Endpoint` and `mesh.Subscriber`.
- Target Kinds are implemented as `mesh.KindRoute` and `mesh.KindTopic`.
- `Message.Payload` is a `[]byte`; `nil` and an empty slice are both an empty payload.
- `Target.Equal` compares segments and kind and ignores metadata.
- The handler signatures `mesh.EndpointHandler` and `mesh.SubscriberHandler`.
- `mesh.Config` and the configuration keys the specification defines: `mesh.DeploymentGroupKey`,
  `mesh.ConsumerGroupKey`, and the value `mesh.ConsumerGroupNone`.
- The errors defined by the specification: `mesh.ErrKindMismatch`,
  `mesh.ErrInvalidTarget`, `mesh.ErrNoDeploymentGroup`.
- The `mesh.Client` and `mesh.Runtime` interfaces.

`Client` and `Runtime` are interfaces. The specification names their methods;
a transport satisfies the contract by implementing them.

## Transports

- NATS: [service-mesh-nats-go](https://github.com/Paymentbox-com/service-mesh-nats-go),
  package `nats`.

The Ruby counterpart of this module is
[service-mesh-ruby](https://github.com/Paymentbox-com/service-mesh-ruby).

## Usage

A transport that implements `Client` and `Runtime` according to the specification uses the
types defined here.

```go
import "github.com/Paymentbox-com/service-mesh-go/mesh"

var target = mesh.Target{Segments: []string{"accounts", "lookup"}, Kind: mesh.KindRoute}

func lookup(ctx context.Context, c mesh.Client, id string) ([]byte, error) {
    reply, err := c.Request(ctx, mesh.Message{Target: target, Payload: []byte(id)}, nil)
    if err != nil {
        return nil, err
    }
    return reply.Payload, nil
}
```

## Development

```
mise install
just check      # format, vet, test, vulnerability scan, lint
```

## Tests

```
just test
```
