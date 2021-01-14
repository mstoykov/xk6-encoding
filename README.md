# xk6-encoding

This is a [k6](https://github.com/loadimpact/k6) extension using the [xk6](https://github.com/k6io/xk6) system.

| :exclamation: This is a proof of concept, isn't supported by the k6 team, and may break in the future. USE AT YOUR OWN RISK! |
| ---------------------------------------------------------------------------------------------------------------------------- |

This a try at fixing some of the problems with the internal k6 encoding package


## Build

To build a `k6` binary with this extension, first ensure you have the prerequisites:

- [gvm](https://github.com/moovweb/gvm)
- [Git](https://git-scm.com/)

Then, install [xk6](https://github.com/k6io/xk6) and build your custom k6 binary with the Kafka extension:

1. Install `xk6`:
  ```shell
  $ go get -u github.com/k6io/xk6/cmd/xk6
  ```

2. Build the binary:
  ```shell
  $ xk6 build v0.29.0 --with github.com/mstoykov/xk6-encoding
  ```
