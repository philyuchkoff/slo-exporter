# Operating

## Debugging
If you need to dynamically change the log level of the application, you can use the `/logging` HTTP endpoint.
The endpoint is enabled by default and can be disabled with the `--enable-dynamic-logging=false` flag.
To set the log level use the `POST` method with URL parameter `level` of value `error`, `warning`, `info` or `debug`.

Example using `cURL`
```bash
# Use GET to get current log level.
$ curl -s http://0.0.0.0:8080/logging
current logging level is: debug

# Use POST to set the log level.
$ curl -XPOST -s http://0.0.0.0:8080/logging?level=info
logging level set to: info
```

#### Profiling
In case of issues with leaking resources for example, slo-exporter supports the
Go profiling using pprof on `/debug/pprof/` web interface path. For usage see the official [docs](https://golang.org/pkg/net/http/pprof/).

Profiling is **disabled by default** as it adds noticeable overhead. Enable it with the `--enable-pprof` flag:
```bash
$ ./slo_exporter --config-file config.yml --enable-pprof
```
Optionally, you can tune the profiling intensity with `--block-profile-rate` and
`--mutex-profile-fraction` (both default to `0`, meaning disabled). See the
[Go documentation](https://pkg.go.dev/runtime#SetBlockProfileRate) for the meaning of the values.


## Frequently asked questions

### How to add new normalization replacement rule?
Event normalization can be done using the `relabel` module, see [its documentation](modules/relabel.md).

### How to deal with malformed lines?
Before !87. If you are seeing too many malformed lines then you should inspect [tailer package](pkg/tailer/tailer.go) and seek for variable `lineParseRegexp`.
After !87, slo-exporter main config supports to specify custom regular expression in field `.module.tailer.loglineParseRegexp`.
