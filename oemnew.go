package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
)

type OEM struct {
}

//this struct is for purchase order
type PoOrder struct {
	PoID                 string       `json:"poID"`
	OemID                string       `json:"oemID"`
	OemName              string       `json:"oemName"`
	OemAddress           string       `json:"oemAddress"`
	DealerID             string       `json:"dealerID"`
	DealerName           string       `json:"dealerName"`
	DealerAddress        string       `json:"dealerAddress"`
	Order                OrderDetails `json:"order"`
	PoAmount             string       `json:"poAmount"`
	PoCreationDate       string       `json:"poCreationDate"`
	ExpectedDeliveryDate string       `json:"expectedDeliveryDate"`
	Status               string       `json:"status"`
	FinancialAgreementNo string		  `json:"financialAgreementNo"`
}

//this struct is to create orderdetails
type OrderDetails struct {
	VehicleMake string `json:"make"`
	Model       string `json:"model"`
	Price       string `json:"price"`
	Quantity	string	`json:"quantity"`
}

//this struct is to create InvoiceDetails
type InvoiceDetails struct {
	InvoiceID     string	`json:"invoiceID"`
	PoID          string	`json:"poID"`												
	Vin           string	`json:"vin"`
	VehicleMake   string	`json:"make"`
	Model         string	`json:"model"`
	InvoiceDate   string	`json:"invoiceDate"`
	InvoiceAmount string	`json:"invoiceAmount"`
	Status        string	`json:"status"`
}

func main() {
	err := shim.Start(new(OEM))
	if err != nil {
		fmt.Printf("Error is starting the ChainCode: %s", err)
	}
}

func (t *OEM) Init(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	
	err := createTableOne(stub)

	if err != nil {
		return nil, fmt.Errorf("Error creating table one during init. %s", err)
	}
	
	e := createTableTwo(stub)

	if e != nil {
		return nil, fmt.Errorf("Error creating table one during init. %s", err)
	}
	
	return []byte("Success"), nil
}

// Create PurchaseOrder table one
func createTableOne(stub shim.ChaincodeStubInterface) error {

	var columnTableOne []*shim.ColumnDefinition
	columnOne := shim.ColumnDefinition{Name: "poID",
		Type: shim.ColumnDefinition_STRING, Key: true}
	columnTwo := shim.ColumnDefinition{Name: "oemID",
		Type: shim.ColumnDefinition_STRING, Key: false}
	columnThree := shim.ColumnDefinition{Name: "oemName",
		Type: shim.ColumnDefinition_STRING, Key: false}
	columnFour := shim.ColumnDefinition{Name: "oemAddress",
		Type: shim.ColumnDefinition_STRING, Key: false}
	columnFive := shim.ColumnDefinition{Name: "dealerID",
		Type: shim.ColumnDefinition_STRING, Key: false}
	columnSix := shim.ColumnDefinition{Name: "dealerName",
		Type: shim.ColumnDefinition_STRING, Key: false}
	columnSeven := shim.ColumnDefinition{Name: "dealerAddress",
		Type: shim.ColumnDefinition_STRING, Key: false}
	columnEight := shim.ColumnDefinition{Name: "vehicleMake",
		Type: shim.ColumnDefinition_STRING, Key: false}
	columnNine := shim.ColumnDefinition{Name: "model",
		Type: shim.ColumnDefinition_STRING, Key: false}
	columnten := shim.ColumnDefinition{Name: "price",
		Type: shim.ColumnDefinition_STRING, Key: false}
	columnEleven := shim.ColumnDefinition{Name: "poAmount",
		Type: shim.ColumnDefinition_STRING, Key: false}
	columnTwelve := shim.ColumnDefinition{Name: "poCreationDate",
		Type: shim.ColumnDefinition_STRING, Key: false}
	columnThirteen := shim.ColumnDefinition{Name: "expectedDeliveryDate",
		Type: shim.ColumnDefinition_STRING, Key: false}
	columnFourteen := shim.ColumnDefinition{Name: "status",
		Type: shim.ColumnDefinition_STRING, Key: false}
	columnFifteen := shim.ColumnDefinition{Name: "quantity",
		Type: shim.ColumnDefinition_STRING, Key: false}
	columnSixteen := shim.ColumnDefinition{Name: "financialAgreementNo",
		Type: shim.ColumnDefinition_STRING, Key: false}

	columnTableOne = append(columnTableOne, &columnOne)
	columnTableOne = append(columnTableOne, &columnTwo)
	columnTableOne = append(columnTableOne, &columnThree)
	columnTableOne = append(columnTableOne, &columnFour)
	columnTableOne = append(columnTableOne, &columnFive)
	columnTableOne = append(columnTableOne, &columnSix)
	columnTableOne = append(columnTableOne, &columnSeven)
	columnTableOne = append(columnTableOne, &columnEight)
	columnTableOne = append(columnTableOne, &columnNine)
	columnTableOne = append(columnTableOne, &columnten)
	columnTableOne = append(columnTableOne, &columnEleven)
	columnTableOne = append(columnTableOne, &columnTwelve)
	columnTableOne = append(columnTableOne, &columnThirteen)
	columnTableOne = append(columnTableOne, &columnFourteen)
	columnTableOne = append(columnTableOne, &columnFifteen)
	columnTableOne = append(columnTableOne, &columnSixteen)
	
	return stub.CreateTable("PurchaseOrder", columnTableOne)
}
//Create Invoice Table
func createTableTwo(stub shim.ChaincodeStubInterface) error {

	var columnTableOne []*shim.ColumnDefinition
	columnOne := shim.ColumnDefinition{Name: "invoiceID",  
		Type: shim.ColumnDefinition_STRING, Key: true}
	columnTwo := shim.ColumnDefinition{Name: "poID",
		Type: shim.ColumnDefinition_STRING, Key: false}  
	columnThree := shim.ColumnDefinition{Name: "vin",
		Type: shim.ColumnDefinition_STRING, Key: false}
	columnFour := shim.ColumnDefinition{Name: "make",
		Type: shim.ColumnDefinition_STRING, Key: false}
	columnFive := shim.ColumnDefinition{Name: "model",
		Type: shim.ColumnDefinition_STRING, Key: false}
	columnSix := shim.ColumnDefinition{Name: "invoiceDate",
		Type: shim.ColumnDefinition_STRING, Key: false}
	columnSeven := shim.ColumnDefinition{Name: "invoiceAmount",
		Type: shim.ColumnDefinition_STRING, Key: false}
	columnEight := shim.ColumnDefinition{Name: "status",
		Type: shim.ColumnDefinition_STRING, Key: false}
	
	columnTableOne = append(columnTableOne, &columnOne)
	columnTableOne = append(columnTableOne, &columnTwo)
	columnTableOne = append(columnTableOne, &columnThree)
	columnTableOne = append(columnTableOne, &columnFour)
	columnTableOne = append(columnTableOne, &columnFive)
	columnTableOne = append(columnTableOne, &columnSix)
	columnTableOne = append(columnTableOne, &columnSeven)
	columnTableOne = append(columnTableOne, &columnEight)
	
	
	return stub.CreateTable("InvoiceDetails", columnTableOne)
}
func (t *OEM) Invoke(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	switch function {

	case "insertPo":
		if len(args) < 1 {
			return nil, errors.New("insertTableOne failed. Must include 14 column values")
		}

		col1Val := args[0]
		col2Val := args[1]
		col3Val := args[2]
		col4Val := args[3]
		col5Val := args[4]
		col6Val := args[5]
		col7Val := args[6]
		col8Val := args[7]
		col9Val := args[8]
		col10Val := args[9]
		col11Val := args[10]
		col12Val := args[11]
		col13Val := args[12]
		col14Val := args[13]
		col15Val := args[14]
		col16Val := args[15]
		
		

		var columns []*shim.Column
		col1 := shim.Column{Value: &shim.Column_String_{String_: col1Val}}
		col2 := shim.Column{Value: &shim.Column_String_{String_: col2Val}}
		col3 := shim.Column{Value: &shim.Column_String_{String_: col3Val}}
		col4 := shim.Column{Value: &shim.Column_String_{String_: col4Val}}
		col5 := shim.Column{Value: &shim.Column_String_{String_: col5Val}}
		col6 := shim.Column{Value: &shim.Column_String_{String_: col6Val}}
		col7 := shim.Column{Value: &shim.Column_String_{String_: col7Val}}
		col8 := shim.Column{Value: &shim.Column_String_{String_: col8Val}}
		col9 := shim.Column{Value: &shim.Column_String_{String_: col9Val}}
		col10 := shim.Column{Value: &shim.Column_String_{String_: col10Val}}
		col11 := shim.Column{Value: &shim.Column_String_{String_: col11Val}}
		col12 := shim.Column{Value: &shim.Column_String_{String_: col12Val}}
		col13 := shim.Column{Value: &shim.Column_String_{String_: col13Val}}
		col14 := shim.Column{Value: &shim.Column_String_{String_: col14Val}}
		col15 := shim.Column{Value: &shim.Column_String_{String_: col15Val}}
		col16 := shim.Column{Value: &shim.Column_String_{String_: col16Val}}

		columns = append(columns, &col1)
		columns = append(columns, &col2)
		columns = append(columns, &col3)
		columns = append(columns, &col4)
		columns = append(columns, &col5)
		columns = append(columns, &col6)
		columns = append(columns, &col7)
		columns = append(columns, &col8)
		columns = append(columns, &col9)
		columns = append(columns, &col10)
		columns = append(columns, &col11)
		columns = append(columns, &col12)
		columns = append(columns, &col13)
		columns = append(columns, &col14)
		columns = append(columns, &col15)
		columns = append(columns, &col16)

		row := shim.Row{Columns: columns}
		ok, err := stub.InsertRow("PurchaseOrder", row)

		if err != nil {
			return nil, fmt.Errorf("insertTableOne operation failed. %s", err)
			panic(err)

		}
		if !ok {
			return []byte("Row with given key" + args[0] + " already exists"), errors.New("insertTableOne operation failed. Row with given key already exists")
		}
	case "updatePo":
		if len(args) < 1 {
			return nil, errors.New("updatePo failed. Must have atleast 1 column values")
		}

		col1Val := args[0]
		
		var columnsUp []shim.Column
		col1 := shim.Column{Value: &shim.Column_String_{String_: col1Val}}
		

		columnsUp = append(columnsUp, col1)
	

		row, err := stub.GetRow("PurchaseOrder", columnsUp)
		if err != nil {
			return nil, fmt.Errorf("getRowTableOne operation failed. %s", err)
			panic(err)
		}

		v := t.convert(row)

		var columns []*shim.Column
		cl1 := shim.Column{Value: &shim.Column_String_{String_: v.PoID}}
		cl2 := shim.Column{Value: &shim.Column_String_{String_: v.OemID}}
		cl3 := shim.Column{Value: &shim.Column_String_{String_: v.OemName}}
		cl4 := shim.Column{Value: &shim.Column_String_{String_: v.OemAddress}}
		cl5 := shim.Column{Value: &shim.Column_String_{String_: v.DealerID}}
		cl6 := shim.Column{Value: &shim.Column_String_{String_: v.DealerName}}
		cl7 := shim.Column{Value: &shim.Column_String_{String_: v.DealerAddress}}
		cl8 := shim.Column{Value: &shim.Column_String_{String_: v.Order.VehicleMake}}
		cl9 := shim.Column{Value: &shim.Column_String_{String_: v.Order.Model}}
		cl10 := shim.Column{Value: &shim.Column_String_{String_: v.Order.Price}}
		cl11 := shim.Column{Value: &shim.Column_String_{String_: v.PoAmount}}
		cl12 := shim.Column{Value: &shim.Column_String_{String_: v.PoCreationDate}}
		cl13 := shim.Column{Value: &shim.Column_String_{String_: v.ExpectedDeliveryDate}}
		cl14 := shim.Column{Value: &shim.Column_String_{String_: args[1]}}
		cl15 := shim.Column{Value: &shim.Column_String_{String_: v.Order.Quantity}}
		columns = append(columns, &cl1)
		columns = append(columns, &cl2)
		columns = append(columns, &cl3)
		columns = append(columns, &cl4)
		columns = append(columns, &cl5)
		columns = append(columns, &cl6)
		columns = append(columns, &cl7)
		columns = append(columns, &cl8)
		columns = append(columns, &cl9)
		columns = append(columns, &cl10)
		columns = append(columns, &cl11)
		columns = append(columns, &cl12)
		columns = append(columns, &cl13)
		columns = append(columns, &cl14)
		columns = append(columns, &cl15)
		if args[2]!= ""{
		cl16 := shim.Column{Value: &shim.Column_String_{String_: args[2]}}	
		columns = append(columns, &cl16)
		}else{
		cl16 := shim.Column{Value: &shim.Column_String_{String_: v.FinancialAgreementNo}}
		columns = append(columns, &cl16)
		}
		
		
		
		rw := shim.Row{Columns: columns}
		ok, err := stub.ReplaceRow("PurchaseOrder", rw)

		if err != nil {
			return nil, fmt.Errorf("getRowTableOne operation failed. %s", err)
			panic(err)
		}
		if !ok {
			return nil, errors.New("replaceRowTableOne operation failed. Row with given key does not exist")
		}

	
	case "insertInvoice":
		if len(args) < 1 {
			return nil, errors.New("insertTableOne failed. Must include 14 column values")
		}
		
		data:=[]string{args[1]}
		valAsbytes, err := t.Query(stub, "getPoByID", data)
		if err != nil {
		return nil, nil
		panic(err)	
		}
		
		
		u:=PoOrder{}
		json.Unmarshal(valAsbytes, &u)
		
		if u.DealerID == ""{
			return nil, nil
		} else{
		col1Val := args[0]
		col2Val := args[1]
		col3Val := args[2]
		col4Val := args[3]
		col5Val := args[4]
		col6Val := args[5]
		col7Val := args[6]
		col8Val := args[7]
	
		var columns []*shim.Column
		col1 := shim.Column{Value: &shim.Column_String_{String_: col1Val}}
		col2 := shim.Column{Value: &shim.Column_String_{String_: col2Val}}
		col3 := shim.Column{Value: &shim.Column_String_{String_: col3Val}}
		col4 := shim.Column{Value: &shim.Column_String_{String_: col4Val}}
		col5 := shim.Column{Value: &shim.Column_String_{String_: col5Val}}
		col6 := shim.Column{Value: &shim.Column_String_{String_: col6Val}}
		col7 := shim.Column{Value: &shim.Column_String_{String_: col7Val}}
		col8 := shim.Column{Value: &shim.Column_String_{String_: col8Val}}

		columns = append(columns, &col1)
		columns = append(columns, &col2)
		columns = append(columns, &col3)
		columns = append(columns, &col4)
		columns = append(columns, &col5)
		columns = append(columns, &col6)
		columns = append(columns, &col7)
		columns = append(columns, &col8)
		
		row := shim.Row{Columns: columns}
		ok, err := stub.InsertRow("InvoiceDetails", row)

		if err != nil {
			return nil, fmt.Errorf("insertTableOne operation failed. %s", err)
			panic(err)

		}
		if !ok {
			return []byte("Row with given key" + args[0] + " already exists"), errors.New("insertTableOne operation failed. Row with given key already exists")
		}
		}
		
		case  "updateInvoice":
		if len(args) < 1 {
			return nil, errors.New("updatePo failed. Must have atleast 1 column values")
		}

		col1Val := args[0]
		
		var columnsUp []shim.Column
		col1 := shim.Column{Value: &shim.Column_String_{String_: col1Val}}
		

		columnsUp = append(columnsUp, col1)
	

		row, err := stub.GetRow("InvoiceDetails", columnsUp)
		if err != nil {
			return nil, fmt.Errorf("getRowTableOne operation failed. %s", err)
			panic(err)
		}
		v := t.convertInvoice(row)
		fmt.Println("the status id"+v.Status+"the new status is"+args[1])
		var columns []*shim.Column
		cl1 := shim.Column{Value: &shim.Column_String_{String_: v.InvoiceID}}
		cl2 := shim.Column{Value: &shim.Column_String_{String_: v.PoID}}
		cl3 := shim.Column{Value: &shim.Column_String_{String_: v.Vin}}
		cl4 := shim.Column{Value: &shim.Column_String_{String_: v.VehicleMake}}
		cl5 := shim.Column{Value: &shim.Column_String_{String_: v.Model}}
		cl6 := shim.Column{Value: &shim.Column_String_{String_: v.InvoiceDate}}
		cl7 := shim.Column{Value: &shim.Column_String_{String_: v.InvoiceAmount}}
		cl8 := shim.Column{Value: &shim.Column_String_{String_: args[1]}}

		columns = append(columns, &cl1)
		columns = append(columns, &cl2)
		columns = append(columns, &cl3)
		columns = append(columns, &cl4)
		columns = append(columns, &cl5)
		columns = append(columns, &cl6)
		columns = append(columns, &cl7)
		columns = append(columns, &cl8)
		
		rw := shim.Row{Columns: columns}
		ok, err := stub.ReplaceRow("InvoiceDetails", rw)

		if err != nil {
			return nil, fmt.Errorf("replaceRowTableOne operation failed. %s", err)
			panic(err)
		}
		if !ok {
			return nil, errors.New("replaceRowTableOne operation failed. Row with given key does not exist")
		}


	default:
		return nil, errors.New("Unsupported operation")
	
		}
return nil,nil
}

func (t *OEM) Query(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	switch function {

	case "getPoByID":
		if len(args) < 1 {
			return nil, errors.New("getRowTableOne failed. Must include 1 key value")
		}
		col1Val := args[0]
		
		var columns []shim.Column
		col1 := shim.Column{Value: &shim.Column_String_{String_: col1Val}}
	

		columns = append(columns, col1)
	

		row, err := stub.GetRow("PurchaseOrder", columns)
		if err != nil {
			return []byte("No record found for id: " + col1Val), fmt.Errorf("getRowTableOne operation failed. %s", err)
			panic(err)
		}
		//fmt.Println(row.Columns[0].GetString_()+" :it is the value")
		//var v=PoOrder{}

		if &row == nil {

			return []byte("No record found for id: " + col1Val), errors.New("getRowTableOne operation failed.")
			panic("No record found for id")
		}

		if len(row.Columns) < 1 {

			return []byte("No record found for id: " + col1Val), errors.New("getRowTableOne operation failed.")
			panic("No record found for id")
		}

		v := t.convert(row)

		json_byte, err := json.Marshal(v)

		if err != nil {

			return nil, err
			panic(err)
		}
		return json_byte, nil

	case "getAllPo":
		if len(args) < 1 {
			return nil, errors.New("getRowsTableFour failed. Must include 1 key value")
		}

		var columns []shim.Column

		rowChannel, err := stub.GetRows("PurchaseOrder", columns)

		if err != nil {
			return nil, fmt.Errorf("getRowsTableFour operation failed. %s", err)
		}

		res2E := []*PoOrder{}

		for {
			select {

			case row, ok := <-rowChannel:

				if !ok {
					rowChannel = nil
				} else {

					u := new(PoOrder)
					u.PoID = row.Columns[0].GetString_()
					u.OemID = row.Columns[1].GetString_()
					u.OemName = row.Columns[2].GetString_()
					u.OemAddress = row.Columns[3].GetString_()
					u.DealerID = row.Columns[4].GetString_()
					u.DealerName = row.Columns[5].GetString_()
					u.DealerAddress = row.Columns[6].GetString_()
					v := OrderDetails{}
					v.VehicleMake = row.Columns[7].GetString_()
					v.Model = row.Columns[8].GetString_()
					v.Price = row.Columns[9].GetString_()
					v.Quantity=row.Columns[14].GetString_()
					u.Order = v
					u.PoAmount = row.Columns[10].GetString_()
					u.PoCreationDate = row.Columns[11].GetString_()
					u.ExpectedDeliveryDate = row.Columns[12].GetString_()
					u.Status = row.Columns[13].GetString_()
					u.FinancialAgreementNo=row.Columns[15].GetString_()
					res2E = append(res2E, u)
				}
			}
			if rowChannel == nil {
				break
			}
		}

		jsonRows, err := json.Marshal(res2E)

		if err != nil {
			return nil, fmt.Errorf("getRowsTableFour operation failed. Error marshaling JSON: %s", err)
		}

		return jsonRows, nil
	
	case "getAllPoByOemId":
		if len(args) < 1 {
			return nil, errors.New("getRowsTableFour failed. Must include 1 key value")
		}

		var columns []shim.Column

		rowChannel, err := stub.GetRows("PurchaseOrder", columns)

		if err != nil {
			return nil, fmt.Errorf("getRowsTableFour operation failed. %s", err)
		}

		res2E := []*PoOrder{}

		for {
			select {

			case row, ok := <-rowChannel:

				if !ok {
					rowChannel = nil
				} else {

					u := new(PoOrder)
					u.PoID = row.Columns[0].GetString_()
				
					u.OemName = row.Columns[2].GetString_()
					u.OemAddress = row.Columns[3].GetString_()
					u.DealerID = row.Columns[4].GetString_()
					u.DealerName = row.Columns[5].GetString_()
					u.DealerAddress = row.Columns[6].GetString_()
					v := OrderDetails{}
					v.VehicleMake = row.Columns[7].GetString_()
					v.Model = row.Columns[8].GetString_()
					v.Price = row.Columns[9].GetString_()
					v.Quantity=row.Columns[14].GetString_()
					u.Order = v
					u.PoAmount = row.Columns[10].GetString_()
					u.PoCreationDate = row.Columns[11].GetString_()
					u.ExpectedDeliveryDate = row.Columns[12].GetString_()
					u.Status = row.Columns[13].GetString_()
					u.FinancialAgreementNo=row.Columns[15].GetString_()
					if row.Columns[1].GetString_() == args[0]{
						u.OemID = row.Columns[1].GetString_()
					res2E = append(res2E, u)
					}
				}
			}
			if rowChannel == nil {
				break
			}
		}

		jsonRows, err := json.Marshal(res2E)

		if err != nil {
			return nil, fmt.Errorf("getRowsTableFour operation failed. Error marshaling JSON: %s", err)
		}
		
		return jsonRows, nil
		
		case "getAllPoByDealerId":
		if len(args) < 1 {
			return nil, errors.New("getRowsTableFour failed. Must include 1 key value")
		}

		var columns []shim.Column

		rowChannel, err := stub.GetRows("PurchaseOrder", columns)

		if err != nil {
			return nil, fmt.Errorf("getRowsTableFour operation failed. %s", err)
		}

		res2E := []*PoOrder{}

		for {
			select {

			case row, ok := <-rowChannel:

				if !ok {
					rowChannel = nil
				} else {

					u := new(PoOrder)
					u.PoID = row.Columns[0].GetString_()
					u.OemID = row.Columns[1].GetString_()
					u.OemName = row.Columns[2].GetString_()
					u.OemAddress = row.Columns[3].GetString_()
					u.DealerName = row.Columns[5].GetString_()
					u.DealerAddress = row.Columns[6].GetString_()
					v := OrderDetails{}
					v.VehicleMake = row.Columns[7].GetString_()
					v.Model = row.Columns[8].GetString_()
					v.Price = row.Columns[9].GetString_()
					v.Quantity=row.Columns[14].GetString_()
					u.Order = v
					u.PoAmount = row.Columns[10].GetString_()
					u.PoCreationDate = row.Columns[11].GetString_()
					u.ExpectedDeliveryDate = row.Columns[12].GetString_()
					u.Status = row.Columns[13].GetString_()
					u.FinancialAgreementNo=row.Columns[15].GetString_()
					if row.Columns[4].GetString_() == args[0]{
						u.DealerID = row.Columns[4].GetString_()
					res2E = append(res2E, u)
					}
					
				}
			}
			if rowChannel == nil {
				break
			}
		}

		jsonRows, err := json.Marshal(res2E)

		if err != nil {
			return nil, fmt.Errorf("getRowsTableFour operation failed. Error marshaling JSON: %s", err)
		}

		return jsonRows, nil
		
		case "getInvoiceById":
		if len(args) < 1 {
			return nil, errors.New("getRowTableOne failed. Must include 1 key value")
		}
		col1Val := args[0]
		
		var columns []shim.Column
		col1 := shim.Column{Value: &shim.Column_String_{String_: col1Val}}
	

		columns = append(columns, col1)
	

		row, err := stub.GetRow("InvoiceDetails", columns)
		if err != nil {
			return []byte("No record found for id: " + col1Val), fmt.Errorf("getRowTableOne operation failed. %s", err)
			panic(err)
		}
		if len(row.Columns) < 1 {

			return []byte("No record found for id: " + col1Val), errors.New("getRowTableOne operation failed.")
			panic("No record found for id")
		}

		v := t.convertInvoice(row)
		
		json_byte, err := json.Marshal(v)

		if err != nil {

			return nil, err
			panic(err)
		}
		return json_byte, nil
		
		case "getAllInvoiceByPoID":
		if len(args) < 1 {
			return nil, errors.New("getRowsTableFour failed. Must include 1 key value")
		}

		var columns []shim.Column

		rowChannel, err := stub.GetRows("InvoiceDetails", columns)

		if err != nil {
			return nil, fmt.Errorf("getRowsTableFour operation failed. %s", err)
		}

		res2E := []*InvoiceDetails{}

		for {
			select {

			case row, ok := <-rowChannel:

				if !ok {
					rowChannel = nil
				} else {

					u := new(InvoiceDetails)
					u.InvoiceID = row.Columns[0].GetString_()
					u.Vin = row.Columns[2].GetString_()
					u.VehicleMake = row.Columns[3].GetString_()
					u.Model = row.Columns[4].GetString_()
					u.InvoiceDate = row.Columns[5].GetString_()
					u.InvoiceAmount = row.Columns[6].GetString_()
					u.Status = row.Columns[7].GetString_()
					if row.Columns[1].GetString_() == args[0]{
						u.PoID = row.Columns[1].GetString_()
					res2E = append(res2E, u)
					}
					
				}
			}
			if rowChannel == nil {
				break
			}
		}

		jsonRows, err := json.Marshal(res2E)

		if err != nil {
			return nil, fmt.Errorf("getRowsTableFour operation failed. Error marshaling JSON: %s", err)
		}

		return jsonRows, nil
		
		case "getAllInvoiceByStatus":
		if len(args) < 1 {
			return nil, errors.New("getRowsTableFour failed. Must include 1 key value")
		}

		var columns []shim.Column

		rowChannel, err := stub.GetRows("InvoiceDetails", columns)

		if err != nil {
			return nil, fmt.Errorf("getRowsTableFour operation failed. %s", err)
		}

		res2E := []*InvoiceDetails{}

		for {
			select {

			case row, ok := <-rowChannel:

				if !ok {
					rowChannel = nil
				} else {

					u := new(InvoiceDetails)
					u.InvoiceID = row.Columns[0].GetString_()
					u.PoID = row.Columns[1].GetString_()
					u.Vin = row.Columns[2].GetString_()
					u.VehicleMake = row.Columns[3].GetString_()
					u.Model = row.Columns[4].GetString_()
					u.InvoiceDate = row.Columns[5].GetString_()
					u.InvoiceAmount = row.Columns[6].GetString_()
					if row.Columns[7].GetString_() == args[0]{
						u.Status = row.Columns[7].GetString_()
					res2E = append(res2E, u)
					}
					
				}
			}
			if rowChannel == nil {
				break
			}
		}

		jsonRows, err := json.Marshal(res2E)

		if err != nil {
			return nil, fmt.Errorf("getRowsTableFour operation failed. Error marshaling JSON: %s", err)
		}

		return jsonRows, nil
		
	
	default:
		return nil, errors.New("Unsupported operation")
	}
}

//this function convert the Row from PoOrder table to PoOrder struct to return required Json
func (t *OEM) convert(row shim.Row) PoOrder {

	var u = PoOrder{}

	u.PoID = row.Columns[0].GetString_()
	u.OemID = row.Columns[1].GetString_()
	u.OemName = row.Columns[2].GetString_()
	u.OemAddress = row.Columns[3].GetString_()
	u.DealerID = row.Columns[4].GetString_()
	u.DealerName = row.Columns[5].GetString_()
	u.DealerAddress = row.Columns[6].GetString_()
	v := OrderDetails{}
	v.VehicleMake = row.Columns[7].GetString_()
	v.Model = row.Columns[8].GetString_()
	v.Price = row.Columns[9].GetString_()
	v.Quantity=row.Columns[14].GetString_()
	u.Order = v
	u.PoAmount = row.Columns[10].GetString_()
	u.PoCreationDate = row.Columns[11].GetString_()
	u.ExpectedDeliveryDate = row.Columns[12].GetString_()
	u.Status = row.Columns[13].GetString_()
	u.FinancialAgreementNo=row.Columns[15].GetString_()
	
	return u
}

//this function convert the Row from InvoiceDetails table to InvoiceDetails struct to return required Json
func (t *OEM) convertInvoice(row shim.Row) InvoiceDetails {

	var u = InvoiceDetails{}

	u.InvoiceID = row.Columns[0].GetString_()
	u.PoID = row.Columns[1].GetString_()
	u.Vin = row.Columns[2].GetString_()
	u.VehicleMake = row.Columns[3].GetString_()
	u.Model = row.Columns[4].GetString_()
	u.InvoiceDate = row.Columns[5].GetString_()
	u.InvoiceAmount = row.Columns[6].GetString_()
	u.Status = row.Columns[7].GetString_()
	
	fmt.Println("The value of status in covert method is"+u.Status)
	
	return u
}
