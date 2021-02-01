# Digital Certificate Manager (Go version)
## I. Installation
1. Clone this repo
2. Run make
3. Locate the binary in the bin directory
4. Add the binary to your PATH
5. Execute dcm for more information
## II. Commands Reference.
```shell
dcm [command]
```
Available Commands:<br />
<pre>
  certificate   The command line interface for certificate management
  help          Help about any command
  version       Print dcm command version
</pre>
### 1. Renew expire date of certification.
```shell
dcm certificate renew [flags]
```
<pre>
Flags:
  -d, --days int                    Number date expire of new certificate (default 1)
  -h, --help                        help for renew
      --old-cert string             Path to certificate of old certificate that need to renew
      --old-private-key string      Path to private key of old certificate
  -o, --output string               Path to output file of new certificate (default ".dcm/output/renew-cert.pem")
      --parent-cert string          Path to parent certificate (CA or ICA)
      --parent-private-key string   Path to private key of parent certificate (CA or ICA)
</pre>
#
Example command renew expire date of peer certificate:
```shell script
dcm certificate renew  --parent-cert=../data/intermediate-ca/signcerts/ica-cert.pem 
--parent-private-key=../data/intermediate-ca/keystore/ica-key.pem 
--old-cert=../data/peer/output/peer-cert.pem 
--old-private-key=../data/peer/keystore/peer-key.pem 
--output=../data/peer/output/renew-peer-cert.pem
--days=365
```

### 2. Inspect Certificate.
```shell script
dcm certificate inspect [flags]
```
<pre>
Flags:
  -c, --cert string   Path to your certificate
  -h, --help          help for inspect
</pre>
#
Example command inspect of peer certificate:
```shell script
dcm certificate inspect -c=../data/peer/output/peer-cert.pem
```

### 3. Check expire date of certificate.
```shell script
dcm certificate check [flags]
```
<pre>
Flags:
  -c, --cert-path string     Path to your certificate
  -f, --folder-cert string   Path to folder have certificates
  -h, --help                 help for check
</pre>
#### a. Single cert file.
Example command check expire date of peer certificate:
```shell script
dcm certificate check -c=../data/peer/output/renew-peer-cert.pem
```

#### b. Folder certificate.
```shell script
dcm certificate check -f=../data/peer/output
```