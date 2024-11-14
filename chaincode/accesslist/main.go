package accesslist

import (
	"al_asset"
	"encoding/json"

	"github.com/hyperledger/fabric-chaincode-go/shim"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
	peer "github.com/hyperledger/fabric-protos-go/peer"
)

type AccessListContract struct {
	contractapi.Contract
}

func (t *AccessListContract) Init(stub shim.ChaincodeStubInterface) peer.Response {
	return shim.Success(nil)
}

func (t *AccessListContract) Invoke(stub shim.ChaincodeStubInterface) peer.Response {
	// Function name and arguments
	fn, args := stub.GetFunctionAndParameters()

	switch fn {
	case "AddIdentity":
		if len(args) != 1 {
			return peer.Response{
				Status:  500,
				Message: "Missing arguments",
				Payload: nil,
			}
		}
		return peer.Response{
			Status:  200,
			Payload: nil,
		}

	case "RemoveIdentity":
		if len(args) != 1 {
			return peer.Response{
				Status:  500,
				Message: "Missing arguments",
				Payload: nil,
			}
		}
		return peer.Response{
			Status:  200,
			Payload: nil,
		}

	case "IsIdentityApproved":
		if len(args) != 1 {
			return peer.Response{
				Status:  500,
				Message: "Missing arguments",
				Payload: nil,
			}
		}
		return peer.Response{
			Status:  200,
			Payload: nil,
		}

	case "GetIdentityList":
		if len(args) != 1 {
			return peer.Response{
				Status:  500,
				Message: "Missing arguments",
				Payload: nil,
			}
		}
		return peer.Response{
			Status:  200,
			Payload: nil,
		}

	default:
		return peer.Response{
			Status:  500,
			Message: "Invalid function name",
			Payload: nil,
		}
	}
}

func createAsset(stub shim.ChaincodeStubInterface, ownerID string) error {
	exists := ownerExists(stub, ownerID)

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

	return stub.PutState(ownerID, assetJSON)
}

func addIdentity(stub shim.ChaincodeStubInterface, ownerID string, proID string) bool {
	exists := ownerExists(stub, ownerID)

	if !exists {
		return false
	}

	alBytes, err := stub.GetState(ownerID)

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

	return stub.PutState(ownerID, accessListJSON) == nil
}

func removeIdentity(stub shim.ChaincodeStubInterface, ownerID string, proID string) bool {
	exists := ownerExists(stub, ownerID)

	if !exists {
		return false
	}

	var accessList al_asset.AccessList
	accessListJSON, err := stub.GetState(ownerID)

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

			err = stub.PutState(ownerID, updatedALJSON)

			if err != nil {
				return false
			}

			return true
		}
	}

	return true
}

func isIdentityApproved(stub shim.ChaincodeStubInterface, ownerID string, proID string) bool {
	exists := ownerExists(stub, ownerID)

	if !exists {
		return false
	}

	var accessList al_asset.AccessList
	accessListJSON, err := stub.GetState(ownerID)

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

func getIdentityList(stub shim.ChaincodeStubInterface, id string) ([]string, bool) {
	exists := ownerExists(stub, id)

	if !exists {
		return nil, false
	}

	var accessList al_asset.AccessList
	accessListJSON, err := stub.GetState(id)

	if err != nil {
		return nil, false
	}

	err = json.Unmarshal(accessListJSON, &accessList)

	if err != nil {
		return nil, false
	}

	return accessList.AllowedIDs, true
}

func ownerExists(stub shim.ChaincodeStubInterface, id string) bool {
	response, err := stub.GetState(id)

	if err != nil {
		return false
	}

	return response != nil
}
