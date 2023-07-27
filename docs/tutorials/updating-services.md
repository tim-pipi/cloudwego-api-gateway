# Updating Services

## Updating Hertz

To update the code after changes in the IDL:

```shell
hz update -idl ../idl/[YOUR_IDL_FILE].thrift
```

**Updating Behaviour**:

- No Custom Path:
  - Appends any new code to the **existing file**.
    - If you rename a method, the old method's code remains in the file.
  - Easier to handle
  - Might create duplicated code
- Custom Path
  - Guaranteed "clean code"
  - Reimplement handler logic each time
  - Confusing to keep track of directories after a while

Update the logic in `handler.go`.

### Service Registration and Discovery

Service registration and discovery is done using [etcd](https://etcd.io/docs/v3.5/)
and the [`registry-etcd`](https://github.com/kitex-contrib/registry-etcd) library.
