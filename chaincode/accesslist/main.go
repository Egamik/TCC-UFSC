package main

import (
	"github.com/hyperledger/fabric-chaincode-go/shim"
	peer "github.com/hyperledger/fabric-protos-go/peer"
)

type AccessList struct{}

func (t *AccessList) Init(stub shim.ChaincodeStubInterface) peer.Response {
	return shim.Success(nil)
}

func (t *AccessList) Invoke(stub shim.ChaincodeStubInterface) peer.Response {
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

func (t *AccessList) addIdentity(stub shim.ChaincodeStubInterface, id string) bool {
	return true
}

func (t *AccessList) removeIdentity(stub shim.ChaincodeStubInterface, id string) bool {
	return true
}

func (t *AccessList) isIdentityApproved(stub shim.ChaincodeStubInterface, id string) bool {
	return true
}

func (t *AccessList) getIdentityList(stub shim.ChaincodeStubInterface, id string) bool {
	return true
}
