package al_contract

import (
	"al_asset"
	"encoding/json"

	chaincodeErrors "chaincodeErrors"

	"github.com/hyperledger/fabric-contract-api-go/v2/contractapi"
)

type AccessListContract struct {
	contractapi.Contract
}

func (s *AccessListContract) createAsset(ctx contractapi.TransactionContextInterface, ownerID string) error {
	funcName := "CreateAsset"
	exists, err := s.ownerExists(ctx, ownerID)
	if err != nil {
		return err
	}

	if !exists {
		return chaincodeErrors.NewAssetNotFoundError(funcName, "ownerID", nil)
	}

	asset := al_asset.AccessList{
		OwnerID:    ownerID,
		AllowedIDs: []string{},
	}

	assetJSON, err := json.Marshal(asset)

	if err != nil {
		return chaincodeErrors.NewMarshallingError(funcName, "Asset", err)
	}

	err = ctx.GetStub().PutState(ownerID, assetJSON)

	if err != nil {
		return chaincodeErrors.NewUpdateWorldStateError(funcName, err)
	}

	return nil
}

func (s *AccessListContract) addIdentity(ctx contractapi.TransactionContextInterface, ownerID string, proID string) error {
	funcName := "addIdentity"

	exists, err := s.ownerExists(ctx, ownerID)

	if !exists {
		return chaincodeErrors.NewGenericError(funcName, err)
	}

	alBytes, err := ctx.GetStub().GetState(ownerID)

	if err != nil {
		return chaincodeErrors.NewReadWorldStateError(funcName, err)
	}

	var accessList al_asset.AccessList

	err = json.Unmarshal(alBytes, &accessList)

	if err != nil {
		return nil
	}

	for _, id := range accessList.AllowedIDs {
		if id == proID {
			return nil
		}
	}

	accessList.AllowedIDs = append(accessList.AllowedIDs, proID)

	accessListJSON, err := json.Marshal(accessList)

	if err != nil {
		return chaincodeErrors.NewMarshallingError(funcName, "AccessList", err)
	}

	err = ctx.GetStub().PutState(ownerID, accessListJSON)

	if err != nil {
		return chaincodeErrors.NewUpdateWorldStateError(funcName, err)
	}

	return nil
}

func (s *AccessListContract) removeIdentity(ctx contractapi.TransactionContextInterface, ownerID string, proID string) error {
	funcName := "RemoveIdentity"
	exists, err := s.ownerExists(ctx, ownerID)

	if !exists {
		return chaincodeErrors.NewAssetNotFoundError(funcName, ownerID, err)
	}

	var accessList al_asset.AccessList
	accessListJSON, err := ctx.GetStub().GetState(ownerID)

	if err != nil {
		return chaincodeErrors.NewReadWorldStateError(funcName, err)
	}

	err = json.Unmarshal(accessListJSON, &accessList)

	if err != nil {
		return chaincodeErrors.NewMarshallingError(funcName, "AccessList", err)
	}

	for i, item := range accessList.AllowedIDs {

		if item == proID {
			accessList.AllowedIDs = append(accessList.AllowedIDs[:i], accessList.AllowedIDs[i+1:]...)
			updatedALJSON, err := json.Marshal(accessList)

			if err != nil {
				return chaincodeErrors.NewMarshallingError(funcName, "AccessList", err)
			}

			err = ctx.GetStub().PutState(ownerID, updatedALJSON)

			if err != nil {
				return chaincodeErrors.NewUpdateWorldStateError(funcName, err)
			}

			return nil
		}
	}

	return nil
}

func (s *AccessListContract) isIdentityApproved(ctx contractapi.TransactionContextInterface, ownerID string, proID string) (bool, error) {
	funcName := "IsIdentityApproved"

	exists, err := s.ownerExists(ctx, ownerID)

	if !exists {
		return false, nil
	}

	var accessList al_asset.AccessList
	accessListJSON, err := ctx.GetStub().GetState(ownerID)

	if err != nil {
		return false, chaincodeErrors.NewReadWorldStateError(funcName, err)
	}

	err = json.Unmarshal(accessListJSON, &accessList)

	if err != nil {
		return false, chaincodeErrors.NewMarshallingError(funcName, "AccessList", err)
	}

	for _, item := range accessList.AllowedIDs {
		if item == proID {
			return true, nil
		}
	}

	return false, nil
}

func (s *AccessListContract) getIdentityList(ctx contractapi.TransactionContextInterface, id string) ([]string, error) {
	exists, err := s.ownerExists(ctx, id)

	if err != nil {
		return nil, err
	}

	if !exists {
		return nil, nil
	}

	var accessList al_asset.AccessList
	accessListJSON, err := ctx.GetStub().GetState(id)

	if err != nil {
		return nil, nil
	}

	err = json.Unmarshal(accessListJSON, &accessList)

	if err != nil {
		return nil, nil
	}

	return accessList.AllowedIDs, nil
}

func (s *AccessListContract) ownerExists(ctx contractapi.TransactionContextInterface, id string) (bool, error) {
	funcName := "ownerExists"

	response, err := ctx.GetStub().GetState(id)

	if err != nil {
		return false, chaincodeErrors.NewReadWorldStateError(funcName, err)
	}

	return response != nil, nil
}
