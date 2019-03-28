# uwping

[![GitHub release](https://img.shields.io/github/release/che0/uwping.svg?style=flat-square)](https://github.com/che0/uwping/releases)
[![Travis](https://img.shields.io/travis/che0/uwping.svg?style=flat-square)](https://travis-ci.org/che0/uwping)

Simple tool to check response code of GET requests to uwsgi socket (by uwsgi protocol)

## Usage

```bash
Usage: ./uwping [options] uwsgi://host:port/path

Parameters:
  -expected_status int
        expected_status (default 200)
  -host string
        HTTP_HOST
  -modifier1 int
        modifier1
  -remote string
        remote addr (default "127.0.0.1")
```

## Examples

```bash
uwget uwsgi://127.0.0.1:3031/health
uwget -expected_status 404 uwsgi://127.0.0.1:3031/not_exist

```
