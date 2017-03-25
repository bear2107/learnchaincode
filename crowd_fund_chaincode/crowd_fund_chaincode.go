package main

import (
        "errors"
        "fmt"
        "strconv"
        "encoding/json"
        "github.com/hyperledger/fabric/core/chaincode/shim"
)

// CrowdFundChaincode implementation
type CrowdFundChaincode struct {
}
type Info struct {

qrcode string
count string


}
//
// Init creates the state variable with name "account" and stores the value
// from the incoming request into this variable. We now have a key/value pair
// for account --> accountValue.
//
func (t *CrowdFundChaincode) Init(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
        // State variable "account"
        // The value stored inside the state variable "account"
        
        // Any error to be reported back to the client
        var err error

        if len(args) != 2 {
                return nil, errors.New("Incorrect number of arguments. Expecting 2.")
        }

        information := Info{}
        informationbyte, err := json.Marshal(information)
     if err!=nil {
                        return nil, err
                }
                err=stub.PutState("default",informationbyte)
         if err!=nil {
                        return nil, err
                }



        return nil, nil
}

//
// Invoke retrieves the state variable "account" and increases it by the ammount
// specified in the incoming request. Then it stores the new value back, thus
// updating the ledger.
//
func (t *CrowdFundChaincode) Invoke(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
        // State variable "account"
 var account string
        // The value stored inside the state variable "account"
//      var accountValue int
        // The ammount by which to increase the state variable
//      var increaseBy int
        // Any error to be reported back to the client
        var err error

        var value int

        if len(args) != 3 {
                return nil, errors.New("Incorrect number of arguments. Expecting 2.")
        }
                value, err = strconv.Atoi(args[2])
        if err != nil {
                return nil, errors.New("Invalid count value, expecting a integer value.")
        }
        // Read in the name of the state variable to be updated
        account = args[0]
        newinfo := Info{
	qrcode: args[1],
        count: args[2]}
        newinfobyte,err:=json.Marshal(newinfo)
        if err!=nil {
                return nil,err
        }

        err = stub.PutState(account, newinfobyte)
        if err != nil {
                return nil, err
        }

        return nil, nil
}

//
// Query retrieves the state variable "account" and returns its current value
// in the response.
//
func (t *CrowdFundChaincode) Query(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
  if function != "query" {
                return nil, errors.New("Invalid query function name. Expecting \"query\".")
        }

        // State variable "account"
        var account string
        // Any error to be reported back to the client
        var err error

        if len(args) != 1 {
                return nil, errors.New("Incorrect number of arguments. Expecting name of the state variable to query.")
        }

        // Read in the name of the state variable to be returned
        account = args[0]
        information:=Info{}
        // Get the current value of the state variable
        accountValueBytes, err := stub.GetState(account)
        if err != nil {
                jsonResp := "{\"Error\":\"Failed to get state for " + account + "\"}"
   		 return nil, errors.New(jsonResp)
        }
        if accountValueBytes == nil {
                jsonResp := "{\"Error\":\"Nil amount for " + account + "\"}"
                return nil, errors.New(jsonResp)
        }
        errUnmarshal:=json.Unmarshal(accountValueBytes,&information)
        if errUnmarshal!=nil {
                return nil,errUnmarshal
        }
        jsonResp := "{\"Name\":\"" + account + "\",\"qrcode\":\"" + information.qrcode + "\",\"count\":\"" + information.count + "\"}"
        fmt.Printf("Query Response:%s\n", jsonResp)
        return accountValueBytes, nil
}

func main() {
        err := shim.Start(new(CrowdFundChaincode))

        if err != nil {
                fmt.Printf("Error starting CrowdFundChaincode: %s", err)
        }
}



