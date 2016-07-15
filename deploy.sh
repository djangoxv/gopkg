#!/bin/bash
# Uses curl to deploy to artifactory on successful run of test
deploy_user="djangoxv"
deploy_pw="APEBJ2ebyFBTBsQ1"
deploy_url="http://52.40.15.68/artifactory/gopkg/latest/gopkg"

curl -u$deploy_user:$deploy_pw -T $1 $deploy_url 
