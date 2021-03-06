[![Build Status](https://travis-ci.org/techjacker/diffence.svg?branch=master)](https://travis-ci.org/techjacker/diffence)
[![Go Report Card](https://goreportcard.com/badge/github.com/techjacker/diffence)](https://goreportcard.com/report/github.com/techjacker/diffence)

# diffence
- Checks a git diff for passwords/secret keys accidentally committed
- Golang 1.7+

-----------------------------------------------------------
### Check the entire history of current branch for passwords/keys committed
```$ git log -p | diffence```

#### Example
```
$ git log -p --full-diff | diffence

------------------
Violation 1
Commit: 4cc087a1b4731d1017844cc86323df43068b0409
File: web/src/db/seed.sql
Reason: "SQL dump file"

------------------
Violation 2
Commit: 142e6019248c0d53a5240242ed1a75c0cc110a0b
File: config/passwords.ini
Reason: "Contains word: password"
```


-----------------------------------------------------------
### Add false positives to `.secignore`

```
$ cat .secignore
file/that/is/not/really/a/secret/but/looks/like/one/to/diffence
these/pems/are/ok/*.pem
```

[See example in this repo](./.secignore).

-----------------------------------------------------------
## Install

### Binary
[Download](../../releases) the latest stable release.

### CLI
```
$ go get -u github.com/techjacker/diffence/cmd/diffence
```

### Library
```
$ go get -u github.com/techjacker/diffence
```

-----------------------------------------------------------
## CLI tool

### Example Usage
```
$ touch key.pem

$ git add -N key.pem

$ git diff --stat HEAD
gds HEAD
 key.pem | 0
 1 file changed, 0 insertions(+), 0 deletions(-)

$ git diff HEAD |diffence
File key.pem violates 1 rules:

Caption: Potential cryptographic private key
Description: <nil>
Part: extension
Pattern: pem
Type: match


```

-----------------------------------------------------------
## Rules
- [x] Analyse fPaths with [gitrob rules](https://github.com/michenriksen/gitrob#signature-keys)
- [ ] Analyse added lines - need to find/create ruleset that can analyse file contents
- [ ] Add option to use your own rules again file path/contents


-----------------------------------------------------------
## Tests
```
$ go test ./...
```

-----------------------------------------------------------
## Local Development

#### Build & Run Locally
```shell
$ go install -race ./cmd/diffence
```
OR
```shell
$ go build -race ./cmd/diffence
```

#### Check for race conditions
```shell
$ go run -race ./cmd/diffence/main.go
```

