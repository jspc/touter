Touter
==

Touter should create a little Golang UDP client which finds git repos (given a base directory) and reports branch names, shas and stuff.

Usage
--

```bash
$ bin/touter
Usage of ./bin/touter:
  -f="/etc/touter.ini": Config file for touter (Shorthand)
  -file="/etc/touter.ini": Config file for touter
  -p=2002: Port to connect on (Shorthand)
  -port=2002: Port to connect on
  -s="localhost": Server to send to (Shorthand)
  -server="localhost": Server to send to
```

Rationale
--

Given a large enough infrastructure of provisioned machines where provisioned machines may number `1..n`, and no centralised register exists of these machines, it becomes difficult to manage and track versions of software deployed. Given the same infrastructure *with* a centralised register and where provisioned machines may number `n..n^n` (where n is a number of machines which makes a worker logging in, querying and parsing too time-intensive) this soon becomes too much work.

Should one, however, build a central server to consume updates and provision a small client on each provisioned machine one can split the work and let the collection server worry about storing the data.

Efficiencies
--

UDP is probably good, so lets use that. The, should the server be busy, we don't care.

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
