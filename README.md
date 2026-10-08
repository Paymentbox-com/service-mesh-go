# service-mesh-go

This version implements the Service Mesh API Specification v0.4.0.

service-mesh-go is the Go contract for the
[Service Mesh API Specification](https://github.com/Paymentbox-com/service-mesh-api).
It is the module `github.com/Paymentbox-com/service-mesh-go`, with one package,
`mesh`. It holds what every transport and every caller must agree on,
and nothing that moves bytes. Transports are separate modules that depend on it
and implement `Client` and `Runtime`.

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
  `mesh.Message`, `mesh.Endpoint`, and `mesh.Subscriber`.
- Target Kinds are implemented as `mesh.KindRoute` and `mesh.KindTopic`.
- `Message.Payload` is a `[]byte`. `nil` and an empty slice are both an empty payload.
- `Target.Equal` compares segments and kind and ignores metadata.
- The handler signatures `mesh.EndpointHandler` and `mesh.SubscriberHandler`.
- `mesh.Config` and the configuration key the specification defines, `mesh.DeploymentGroupKey`,
  which is runtime configuration.
- The consumer group is the `ConsumerGroup` field of `mesh.Endpoint` and `mesh.Subscriber`.
  `mesh.ConsumerGroupNone` is the value for no group.
- The reserved metadata prefix `mesh.ReservedPrefix` (`Mesh-`).
- The metadata keys the specification defines: `mesh.HandlerErrorKey`, `mesh.TimeoutKey`,
  `mesh.DeadlineKey`, `mesh.DeliveryAttemptKey`, and `mesh.MessageIDKey`.
- The errors defined by the specification: `mesh.ErrKindMismatch`,
  `mesh.ErrInvalidTarget`, and `mesh.ErrNoDeploymentGroup`.
- The `mesh.Client` and `mesh.Runtime` interfaces.

`Client` and `Runtime` are interfaces. The specification names their methods,
and a transport satisfies the contract by implementing them.

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
just test       # tests only
```

A release is a tag. `just bump patch`, `just bump minor`, or `just bump major`
raises the version in `VERSION` and commits that file. After the commit is
pushed and passes CI, `just release` tags the commit with it, pushes the tag,
and has the Go module proxy fetch it.
