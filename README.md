![Go Gopher](http://upload.wikimedia.org/wikipedia/commons/2/23/Golang.png?uselang=en-gb)

Touter
==

Touter should create a little Golang UDP client which finds git repos (given a base directory) and reports branch names, shas and stuff.

Usage
--

```bash
Usage of ./bin/touter:
  -d=2: Depth to walk through repo_root (Shorthand)
  -depth=2: Depth to walk through repo_root
  -f="/etc/touter.ini": Config file for touter (Shorthand)
  -file="/etc/touter.ini": Config file for touter
  -p=2002: Port to connect on (Shorthand)
  -port=2002: Port to connect on
  -pr="rails": Profile from config file to use (Shorthand)
  -profile="rails": Profile from config file to use
  -r="/path": Repo root to start from (Shorthand)
  -repo_root="/path": Repo root to start from
  -s="localhost": Server to send to (Shorthand)
  -server="localhost": Server to send to
```

Rationale
--

Given a large enough infrastructure of provisioned machines where provisioned machines may number `1..n`, and no centralised register exists of these machines, it becomes difficult to manage and track versions of software deployed. Given the same infrastructure *with* a centralised register and where provisioned machines may number `n..n^n` (where n is a number of machines which makes a worker logging in, querying and parsing too time-intensive) this soon becomes too much work.

Should one, however, build a central server to consume updates and provision a small client on each provisioned machine one can split the work and let the collection server worry about storing the data.

Configuration
--

We provide a config file to control certain aspects of the app. It exists as a [gcfg](https://code.google.com/p/gcfg/) file, which is heavily based on the style of `.ini` file `git-config` uses. It declares 'profiles' with directories to exclude from search. The default config which accompanies this file is for rails styled capistrano deploys (We don't give a damn about checked out code in the `project/releases` dir). It looks like:

```ini
[profile "rails"]
  description = "A very simple profile to exclude the releases dir on rails apps"    ; Gets output to logs on start; just a prettiness thing
  exclude = "releases"                                                               ; These need to exclude any slashes
  exclude = "tmp"
```

Building
--

We provide a Makefile:

```bash
$ make
go build -a -v -o bin/touter src/touter.go
runtime
errors
sync/atomic
unicode
sync
io
unicode/utf8
math
bytes
syscall
strconv
time
reflect
os
sort
fmt
strings
path/filepath
encoding
io/ioutil
code.google.com/p/gcfg/token
code.google.com/p/gcfg/scanner
code.google.com/p/gcfg/types
encoding/binary
math/rand
encoding/base64
unicode/utf16
encoding/json
math/big
flag
code.google.com/p/gcfg
log
os/exec
runtime/cgo
net
command-line-arguments
```

Testing
--

My-fucking-god does `golang` testing suck. So I haven't bothered. Mind: this little script is so simple that to run it is to test it. That being said; there is a little ruby server configured which may do the job:

```bash
$ ruby src/test_server.rb
Received data about egg.zero-internet.org.uk
Sample repo: /Users/jspc/projects/App-perlbrew
```

Efficiencies
--

UDP is probably good, so lets use that. Then, should the server be busy (or even dead), we don't care.

Licence
--

The MIT License (MIT)

Copyright (c) 2014 James Condron

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
