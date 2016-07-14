#!/bin/bash
# Uses curl to deploy to artifactory on successful run of test

curl -udjangoxv:APEBJ2ebyFBTBsQ1 -T ./gopkg "http://52.40.15.68/artifactory/gopkg/latest/gopkg"
