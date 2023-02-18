# schnorrcli

Very simple CLI for doing Schnorr stuff on command line.

```
go install github.com/lindlof/schnorrcli@latest
key=$(openssl ecparam -name secp256k1 -genkey -outform der | base64)
schnorrcli pubkey $key
schnorrcli sign $key hello
schnorrcli verify $(schnorrcli pubkey $key) $(schnorrcli sign $key hello) hello

echo -n hello > file
schnorrcli signb64 $key $(cat file | base64)
schnorrcli verifyb64 $(schnorrcli pubkey $key) $(schnorrcli sign $key hello) $(cat file | base64)
```
