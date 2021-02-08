# Google Cloud Proto Scrubber

[![PkgGoDev][pkg-badge]][pkg]
[![GoReportCard][report-badge]][report]
[![Codecov][codecov-badge]][codecov]

[pkg-badge]: https://pkg.go.dev/badge/go.einride.tech/google-cloud-proto-scrubber
[pkg]: https://pkg.go.dev/go.einride.tech/google-cloud-proto-scrubber
[report-badge]: https://goreportcard.com/badge/go.einride.tech/google-cloud-proto-scrubber
[report]: https://goreportcard.com/report/go.einride.tech/google-cloud-proto-scrubber
[codecov-badge]: https://codecov.io/gh/einride/google-cloud-proto-scrubber-go/branch/master/graph/badge.svg
[codecov]: https://codecov.io/gh/einride/google-cloud-proto-scrubber

Scrub protobuf/gRPC API descriptors of unsupported annotations before uploading
them to [Google Cloud Endpoints][google-cloud-endpoints] and [Google Cloud
API Gateway][google-cloud-api-gateway].

[google-cloud-endpoints]: https://cloud.google.com/endpoints
[google-cloud-api-gateway]: https://cloud.google.com/api-gateway

# Usage

```bash
$ go run go.einride.tech/google-cloud-proto-scrubber -f descriptor.pb
```
