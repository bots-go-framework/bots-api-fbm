#!/usr/bin/env bash
rm send_api_ffjson.go
rm webhook_api_ffjson.go
ffjson -force-regenerate send_api.go
ffjson -force-regenerate webhook_api.go