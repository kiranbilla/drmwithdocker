
# opendrm
An open source implementation of **industry-grade** DRM(Digital Rights Management) or Key System. The common Key Systems we know are PlayReady of Microsoft, Widevine of Google and FairPlay of Apple. And china has ChinaDRM Key System specification.

This Key System is intended to work with [ISO Common Encryption(CENC) Protection Scheme](https://www.iso.org/obp/ui/#iso:std:iso-iec:23001:-7:ed-2:v1:en) and [W3C EME(Encrypted Media Extensions)](https://w3c.github.io/encrypted-media/). And it provides a Common Key System implementation and ChinaDRM implementation. Maybe we can privode more DRMs implementation like Widevine, PlayReady, etc. 

Anyone is welcome to join me if you are interested or specialized in any of PlayReady, Widevine, FairPlay or others.

# License
This software is under license [GPLv3](https://github.com/willkk/opendrm/blob/master/LICENSE).

# Standalone

$ go mod init opendrm #Initialize the Go module.

$ go mod tidy #Install dependencies.

$ go build #build the server.

$ go run main.go #run main_server.

$ go test ./...  # for test the project.

# Note: 
1.To generate a PEM certificate, use the command below.

# OpenSSL key generation:
$ openssl pkcs8 -topk8 -inform PEM -outform PEM -in /home/ubuntu/goproject/opendrm/test/rsa_private_key.pem -out rsa_private_key_pkcs8.pem -nocrypt

2. Make sure all paths need to be changed based on the environment.

 # Containerization

  Dockerfile run commands

     docker build -t opendrm 
     docker run -itd -p8090:8090 opendrm:latest
