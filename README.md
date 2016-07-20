# QOR Notification

QOR Notification (WIP)

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

## License

Released under the [MIT License](http://opensource.org/licenses/MIT).
