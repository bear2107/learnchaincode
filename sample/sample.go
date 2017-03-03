package main

import (
        "errors"
        "fmt"


        "github.com/hyperledger/fabric/core/chaincode/shim"
)

// CrowdFundChaincode implementation
type CrowdFundChaincode struct {
}

//
// Init creates the state variable with name "account" and stores the value
// from the incoming request into this variable. We now have a key/value pair
// for account --> accountValue.
//
func (t *CrowdFundChaincode) Init(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
        // State variable "account"
        var account string
 // The value stored inside the state variable "account"
        var accountValue  string
        // Any error to be reported back to the client
        var err error

        if len(args) != 2 {
                return nil, errors.New("Incorrect number of arguments. Expecting 2.")
        }

        // Initialize the state variable name
        account = args[0]
        // Initialize the state variable value
        accountValue  = args[1]
//      if err != nil {
//              return nil, errors.New("Expecting integer value for account initialization.")
//      }

//      fmt.Printf("accountValue = %d\n", accountValue)

        // Write the state to the ledger
         err = stub.PutState(account, []byte(accountValue))
        if err != nil {
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
        var accountValue string
        // The ammount by which to increase the state variable
        // Any error to be reported back to the client
        var err error

        if len(args) != 2 {
                return nil, errors.New("Incorrect number of arguments. Expecting 2.")

  // Read in the name of the state variable to be updated
        account = args[0]

        accountValue = args[1]

        // Get the current value of the state variable
//      accountValue, _ = strconv.Atoi(string(accountValueBytes))

        // Update the "account" state variable
//      increaseBy, err = strconv.Atoi(args[1])
//      if err != nil {
//              return nil, errors.New("Invalid transaction amount, expecting a integer value.")
//      }
//      accountValue = accountValue + increaseBy
//      fmt.Printf("accountValue = %d\n", accountValue)

        // Write the state back to the ledger
        err = stub.PutState(account, []byte(accountValue))
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

        jsonResp := "{\"Name\":\"" + account + "\",\"Amount\":\"" + string(accountValueBytes) + "\"}"
 fmt.Printf("Query Response:%s\n", jsonResp)
        return accountValueBytes, nil
}

func main() {
        err := shim.Start(new(CrowdFundChaincode))

        if err != nil {
                fmt.Printf("Error starting CrowdFundChaincode: %s", err)
        }
}

