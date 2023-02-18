# schnorrcli

Very simple CLI for doing Schnorr stuff on command line.

```
go install github.com/lindlof/schnorrcli@latest
key=$(openssl ecparam -name secp256k1 -genkey -outform der | base64)
schnorrcli pubkey $key
schnorrcli sign $key hello
schnorrcli signb64 $key $(echo -n hello | base64)
schnorrcli verify $(schnorrcli pubkey $key) $(schnorrcli sign $key hello) hello
```
