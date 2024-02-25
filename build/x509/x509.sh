#!/bin/bash

#
# Copyright © 2024 honeysense.com All rights reserved.
# Author: sunrui
# Date: 2024-02-25 13:48:30
#

# Create the server CA certs.
openssl req -x509                               \
  -newkey rsa:4096                              \
  -nodes                                        \
  -days 3650                                    \
  -keyout ca_key.pem                            \
  -out ca_cert.pem                              \
  -subj /C=US/ST=CA/L=SVL/O=gRPC/CN=honeysense/ \
  -config ./openssl.cnf                         \
  -extensions ca                                \
  -sha256

# Generate a server cert.
openssl genrsa -out server_key.pem 4096
openssl req -new                                \
  -key server_key.pem                           \
  -out server_csr.pem                           \
  -subj /C=US/ST=CA/L=SVL/O=gRPC/CN=honeysense/ \
  -config ./openssl.cnf                         \
  -reqexts server
openssl x509 -req        \
  -in server_csr.pem     \
  -CAkey ca_key.pem      \
  -CA ca_cert.pem        \
  -days 3650             \
  -set_serial 1000       \
  -out server_cert.pem   \
  -extfile ./openssl.cnf \
  -extensions server   \
  -sha256
openssl verify -verbose -CAfile ca_cert.pem  server_cert.pem

rm *_csr.pem
