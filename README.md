# msisdn-parser

Since the country list and list of providers and their range of subscriber numbers is incomplete, parser will return error, e.g. Invalid MSISDN, for all countries except Slovenia.

Requirements: Vagrant and Go

Developed and tested on Debian Squeeze.

### Usage

```
git clone https://github.com/jaka/msisdn-parser.git
cd msisdn-parser
vagrant up
vagrant ssh
go run /vagrant/src/client/client.go [msisdn]
```
or
```
git clone https://github.com/jaka/msisdn-parser.git
cd msisdn-parser
vagrant up
export GOPATH=`pwd`
export SERVER=192.168.100.100:12345
go run src/client/client.go [msisdn]
```

### Tests
```
git clone https://github.com/jaka/msisdn-parser.git
cd msisdn-parser
vagrant up
vagrant ssh
cd /vagrant/src/msisdn/
go test
```
