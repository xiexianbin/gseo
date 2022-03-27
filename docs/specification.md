# Google Search Console keywords auto render to Hugo post file

## 1 - About This Document

This document is a specification for Google Search Console keywords 
auto render to Hugo post file. 

## 2 - Functions

`gseo` is A Client write by Golang. Func like:

- Call Google Search Console API, and cache the keywords result to `./.gseo/cache-1.json`
- read `./.gseo/cache-1.json` and render hugo post markdown files which the special keywords

`gseo` has the following functions:

- `gseo init ...` init google auth Credentials
- `gseo site` show all sites
- `gseo keyword ...` download hugo post keywords from Google Search Console API, and cache it in `./.gseo/` dir
- `gseo render ...` render hugo post markdown files
- `gseo version` show the client version
