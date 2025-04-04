package al_contract

import (
	"al_asset"
	"encoding/json"

	"github.com/hyperledger/fabric-contract-api-go/v2/contractapi"
)

type AccessListContract struct {
	contractapi.Contract
}

func (s *AccessListContract) createAsset(ctx contractapi.TransactionContextInterface, ownerID string) error {
	exists := s.ownerExists(ctx, ownerID)

	if !exists {
		return nil
	}

	asset := al_asset.AccessList{
		OwnerID:    ownerID,
		AllowedIDs: []string{},
	}

	assetJSON, err := json.Marshal(asset)

	if err != nil {
		return err
	}

	return ctx.GetStub().PutState(ownerID, assetJSON)
}

func (s *AccessListContract) addIdentity(ctx contractapi.TransactionContextInterface, ownerID string, proID string) bool {
	exists := s.ownerExists(ctx, ownerID)

	if !exists {
		return false
	}

	alBytes, err := ctx.GetStub().GetState(ownerID)

	if err != nil {
		return false
	}

	var accessList al_asset.AccessList

	err = json.Unmarshal(alBytes, &accessList)

	if err != nil {
		return false
	}

	for _, id := range accessList.AllowedIDs {
		if id == proID {
			return true
		}
	}

	accessList.AllowedIDs = append(accessList.AllowedIDs, proID)

	accessListJSON, err := json.Marshal(accessList)

	if err != nil {
		return false
	}

	return ctx.GetStub().PutState(ownerID, accessListJSON) == nil
}

func (s *AccessListContract) removeIdentity(ctx contractapi.TransactionContextInterface, ownerID string, proID string) bool {
	exists := s.ownerExists(ctx, ownerID)

	if !exists {
		return false
	}

	var accessList al_asset.AccessList
	accessListJSON, err := ctx.GetStub().GetState(ownerID)

	if err != nil {
		return false
	}

	err = json.Unmarshal(accessListJSON, &accessList)

	if err != nil {
		return false
	}

	for i, item := range accessList.AllowedIDs {
		if item == proID {
			accessList.AllowedIDs = append(accessList.AllowedIDs[:i], accessList.AllowedIDs[i+1:]...)
			updatedALJSON, err := json.Marshal(accessList)
			if err != nil {
				return false
			}

			err = ctx.GetStub().PutState(ownerID, updatedALJSON)

			if err != nil {
				return false
			}

			return true
		}
	}

	return true
}

func (s *AccessListContract) isIdentityApproved(ctx contractapi.TransactionContextInterface, ownerID string, proID string) bool {
	exists := s.ownerExists(ctx, ownerID)

	if !exists {
		return false
	}

	var accessList al_asset.AccessList
	accessListJSON, err := ctx.GetStub().GetState(ownerID)

	if err != nil {
		return false
	}

	err = json.Unmarshal(accessListJSON, &accessList)

	if err != nil {
		return false
	}

	for _, item := range accessList.AllowedIDs {
		if item == proID {
			return true
		}
	}

	return false
}

func (s *AccessListContract) getIdentityList(ctx contractapi.TransactionContextInterface, id string) ([]string, bool) {
	exists := s.ownerExists(ctx, id)

	if !exists {
		return nil, false
	}

	var accessList al_asset.AccessList
	accessListJSON, err := ctx.GetStub().GetState(id)

	if err != nil {
		return nil, false
	}

	err = json.Unmarshal(accessListJSON, &accessList)

	if err != nil {
		return nil, false
	}

	return accessList.AllowedIDs, true
}

func (s *AccessListContract) ownerExists(ctx contractapi.TransactionContextInterface, id string) bool {
	response, err := ctx.GetStub().GetState(id)

	if err != nil {
		return false
	}

	return response != nil
}
