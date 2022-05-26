# gseo

[![build-test](https://github.com/xiexianbin/gseo/actions/workflows/workflow.yaml/badge.svg)](https://github.com/xiexianbin/gseo/actions/workflows/workflow.yaml)
[![GoDoc](https://godoc.org/github.com/xiexianbin/gseo?status.svg)](https://pkg.go.dev/github.com/xiexianbin/gseo)

a golang client to optimize [hugo](https://www.xiexianbin.cn/tags/hugo/) seo by Google Search Console. read [gseo spec](./docs/specification.md) for more information.

## install

- source

```
go install github.com/xiexianbin/gseo
```

- bin

```
curl -Lfs -o gseo https://github.com/xiexianbin/gseo/releases/latest/download/gseo-{linux|darwin|windows}
chmod +x gseo
./gseo
```

## Use

- google auth

```
cat ~/.gseo/client_secret.json
{"installed":{"client_id":"1017408311257-hq3j99vk9ludpoff862mnp52v36nv4gc.apps.googleusercontent.com","project_id":"adept-button-344010","auth_uri":"https://accounts.google.com/o/oauth2/auth","token_uri":"https://oauth2.googleapis.com/token","auth_provider_x509_cert_url":"https://www.googleapis.com/oauth2/v1/certs","client_secret":"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","redirect_uris":["urn:ietf:wg:oauth:2.0:oob","http://localhost"]}}
```

- init

```
gseo init
```

## dev

- debug

```
go test -v -run TestClient google/client_test.go
```

- local run

```
$ make go-build
  >  Building binary...

# default args
$ ./bin/gseo render --content /Users/xiexianbin/workspace/code/github.com/xiexianbin/note.seo/content --position 10 --ctr 0 --impressions 10 --clicks 0.1 --max 8 --dryrun true

# my args 1
$ ./bin/gseo render --content /Users/xiexianbin/workspace/code/github.com/xiexianbin/note.seo/content --position 10 --ctr 0 --impressions 100 --clicks 0.3 --max 8
https://www.xiexianbin.cn/windows/tools/2018-08-19-windows-tcping/
  - tcping windows
update markdownFilePath /Users/xiexianbin/workspace/code/github.com/xiexianbin/note.seo/content/windows/tools/2018-08-19-windows-tcping.md, keywords [tcping windows] done.

# my args 1
$ ./bin/gseo render --content /Users/xiexianbin/workspace/code/github.com/xiexianbin/note.seo/content --position 0 --ctr 0 --impressions 1 --clicks 0 --max 8 --dryrun true
```

## ref

- https://developers.google.com/webmaster-tools/about
- https://developers.google.com/webmaster-tools/v1/api_reference_index
- [Quickstart: Run a Search Console App in Python](https://developers.google.com/webmaster-tools/v1/quickstart/quickstart-python)
  - [Search Console APP Dashboard](https://console.cloud.google.com/apis/api/cloudsearch.googleapis.com/overview)
- [Search Console Testing Tools API (Experimental)](https://developers.google.com/webmaster-tools/search-console-api)
