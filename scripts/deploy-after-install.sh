#!/bin/bash
sudo chown -R ubuntu:ubuntu /tmp/go-to-sample
sudo supervisorctl stop go-to-sample
mv go-to-sample.bin /opt/go-to-sample/
sudo supervisorctl start go-to-sample
