#!/bin/bash
sudo chown -R ubuntu:ubuntu /tmp/go-to-sample
sudo supervisorctl stop go-to-sample
cp /tmp/go-to-sample/go-to-sample.bin /opt/go-to-sample/go-to-sample.bin
sudo supervisorctl start go-to-sample
