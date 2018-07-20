# revproxyhashry

revproxyhashry generates bcrypt hashes of passwords to be used with the reverse proxy 
[revproxyry](https://bitbucket.org/parqueryopen/revproxyry).

## Related Projects

[Apache HTTP server project 2.4](https://httpd.apache.org/docs/2.4) 
includes a utility tool [htpasswd](https://httpd.apache.org/docs/2.4/programs/htpasswd.html) 
  to hash bcrypt passwords. 
  We preferred a simple interface (one reverse proxy, one password hashing utility) to htpasswd's approach 
  which allows you to leverage many different hashings (such as MD5 _etc_.). 

  You can use the following command to generate bcyrpt hashes with `htpasswd` (see this 
  [StackExchange answer](https://unix.stackexchange.com/a/419855)):

```bash
htpasswd -bnBC 10 "" "some-password"
```

## Installation

### Pre-compiled Binaries

We provide the following pre-compiled binaries of the revproxyhashry:

Version|Arch|Release
---|---|---
1.0.0|Linux x64|[revproxyhashry-1.0.0-linux-x64.tar.gz](https://bitbucket.org/parqueryopen/revproxyhashry/downloads/revproxyhashry-1.0.0-linux-x64.tar.gz)

To install the release, just unpack it somewhere, add `bin/` directory to 
your `PATH` and you are ready to go.

### Debian Packages

We also provide a Debian package:

Version|Arch|Release
---|---|---
1.0.0|amd64|[revproxyhashry_1.0.0_amd64.deb](https://bitbucket.org/parqueryopen/revproxyhashry/downloads/revproxyhashry_1.0.0_amd64.deb)

For example, to download the package and install it, call:

```bash
wget https://bitbucket.org/parqueryopen/revproxyhashry/downloads/revproxyhashry_1.0.0_amd64.deb
sudo dpkg -i revproxyhashry_1.0.0_amd64.deb
```

### Compile From Source

Assuming you have a working Go environment, you can install the _revproxyhashry_
from the source by running:

```bash
go get -U bitbucket.org/parqueryopen/revproxyhashry/revproxyhashry
```

## Usage

To generate the password hash, just invoke: 

```bash
revproxyhashry
```

It will prompt you to type in the password and generate the hash once you input it.

You can also provide the password on the command-line:

```bash
revproxyhashry "some-password"
``` 

Bcrypt allows you to set the cost of the hashing 
(see this 
[StackExchange question](https://security.stackexchange.com/questions/17207/recommended-of-rounds-for-bcrypt)). 
The default cost of revproxyhashry is set to 10. You can set the cost of the hash 
by specifying `-cost` command-line argument:

```bash
revproxyhashry -cost 14
```

You can also combine the password together with the cost in the command line:

```bash
revproxyhashry -cost 14 "some-password"
```

## Development

* Clone the repository beneath your `GOPATH`:

```bash
go get bitbucket.org/parqueryopen/revproxyhashry
```

* Change to the _revproxyry_ directory:

```bash
cd $GOPATH/src/bitbucket.org/parqueryopen/revproxyhashry
```

* If you want to build everything in the project:

```bash
go build ./...
```

* If you want to build and install everything to $GOPATH/bin:

```bash
go install ./...
```

* Create a pull request and send it for review `:)`

## Versioning

We follow [Semantic Versioning](http://semver.org/spec/v1.0.0.html). 
The version X.Y.Z indicates:

* X is the major version (backward-incompatible),
* Y is the minor version (backward-compatible) and
* Z is the patch version (backward-compatible bug fix).
