#!/bin/bash
PROJECT_ROOT = $(pwd)

# Prepare chaincodes packages
echo "Preparing chaincodes packages!"

cd $PROJECT_ROOT/chaincode/accesslist
go mod tidy
go mod vendor

cd $PROJECT_ROOT/chaincode/ehr
go mod tidy
go mod vendor

cp -r $PROJECT_ROOT/chaincode/assets/al_asset $PROJECT_ROOT/chaincode/accesslist/vendor
cp -r $PROJECT_ROOT/chaincode/assets/ehr_asset $PROJECT_ROOT/chaincode/ehr/vendor

cp -r chaincode/accesslist minifabric/chaincode
cp -r chaincode/ehr minifabric/chaincode

cd minifabric

