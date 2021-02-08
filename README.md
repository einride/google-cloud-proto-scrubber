# API Descriptor Scrubber

[![PkgGoDev][pkg-badge]][pkg]
[![GoReportCard][report-badge]][report]
[![Codecov][codecov-badge]][codecov]

[pkg-badge]: https://pkg.go.dev/badge/go.einride.tech/api-descriptor-scrubber
[pkg]: https://pkg.go.dev/go.einride.tech/api-descriptor-scrubber
[report-badge]: https://goreportcard.com/badge/go.einride.tech/api-descriptor-scrubber
[report]: https://goreportcard.com/report/go.einride.tech/api-descriptor-scrubber
[codecov-badge]: https://codecov.io/gh/einride/api-descriptor-scrubber-go/branch/master/graph/badge.svg
[codecov]: https://codecov.io/gh/einride/api-descriptor-scrubber-go

Scrub gRPC API descriptors of unsupported annotations before uploading
them to [Google Cloud Endpoints][google-cloud-endpoints] and [Google Cloud
API Gateway][google-cloud-api-gateway].

[google-cloud-endpoints]: https://cloud.google.com/endpoints
[google-cloud-api-gateway]: https://cloud.google.com/api-gateway

# Usage

```bash
$ go run go.einride.tech/api-descriptor-scrubber -f descriptor.pb
```
