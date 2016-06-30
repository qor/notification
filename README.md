# QOR Notification

QOR Notification

## Usage

```go
// controller
// func Index(context *admin.Context) {
// }

notification.Get(context *qor.Context) []Notification

type Notification struct {
  Title string
  Body string
  serializable_meta.SerializableMeta
}

notification.RegisterChannel()
```
