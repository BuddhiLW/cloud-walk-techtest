#!/bin/bash

cd cmd/qrep && go build -o qrep && chmod +x qrep && mv ./qrep ../../ && cd ../../
