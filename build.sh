#!/bin/bash

cd cmd/cw-report && go build -o qrep && chmod +x qrep && mv ./qrep ../../ && cd ../../
