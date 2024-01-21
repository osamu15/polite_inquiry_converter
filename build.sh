#!/bin/bash
cd receive_jira_webhook
go build -buildvcs=false -o main
zip /usr/src/polite_inquiry_converter/receive_jira_webhook/function.zip main
