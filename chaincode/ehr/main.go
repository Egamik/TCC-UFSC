package main

import (
	"github.com/hyperledger/fabric-chaincode-go/shim"
	peer "github.com/hyperledger/fabric-protos-go/peer"
)

type EHealth struct{}

func (t *EHealth) Init(stub shim.ChaincodeStubInterface) peer.Response {
	return shim.Success(nil)
}

func (t *EHealth) Invoke(stub shim.ChaincodeStubInterface) peer.Response {
	// Function name and arguments
	fn, args := stub.GetFunctionAndParameters()

	switch fn {
	case "CreateRecord":
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

	case "RemoveRecord":
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

	case "AddPrescription":
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

	case "AddAppointment":
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
