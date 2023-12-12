# Google Cloud Proto Scrubber

[![PkgGoDev](https://pkg.go.dev/badge/go.einride.tech/google-cloud-proto-scrubber)](https://pkg.go.dev/go.einride.tech/google-cloud-proto-scrubber)
[![GoReportCard](https://goreportcard.com/badge/go.einride.tech/google-cloud-proto-scrubber)](https://goreportcard.com/report/go.einride.tech/google-cloud-proto-scrubber)
[![Codecov](https://codecov.io/gh/einride/google-cloud-proto-scrubber-go/branch/master/graph/badge.svg)](https://codecov.io/gh/einride/google-cloud-proto-scrubber)

Scrub protobuf/gRPC API descriptors of unsupported annotations before uploading
them to [Google Cloud Endpoints](https://cloud.google.com/endpoints) and
[Google Cloud API Gateway](https://cloud.google.com/api-gateway).

# Usage

```bash
$ go run go.einride.tech/google-cloud-proto-scrubber -f descriptor.pb
```
