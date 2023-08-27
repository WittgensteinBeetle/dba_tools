package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"sort"
	"strings"
)

//Logic:  Build a string map so that each key value pair has the following structure:
//table_name.begin = create statement
//table_name.field = definition
//table_name.PRIMARY = definition
//table_name.index = definition
//table_name.end = end create table statement

type BaseVindex struct {
	Sharded string   `json:"sharded"`
	Vindex  HashName `json:"vindexes"`
	//Tables  BaseTable `json:"tables"`
}

type HashName struct {
	HType HashType `json:"hash_idx"`
}

type HashType struct {
	Type string `json:"type"`
}

type MainTables struct {
	TableName string   `json:"tableName"`
	Auto      SeqTable `json:"auto_increment"`
	Vindex    Vindex   `json:"column_vindex"`
}

type SeqTable struct {
	Column        string `json:"column"`
	SequenceTable string `json:"sequence"`
}

type Vindex struct {
	Column string `json:"column"`
	Name   string `json:"name"`
}

func main() {

	readFile, err := os.Open("/Users/josephtotin/Documents/program_input_files/test_dump_schema.sql")

	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	var create_tables = map[string]string{}

	//table names, begin scanning the file only update map when a CREATE TABLE
	//entry is found
	for fileScanner.Scan() {
		line := fileScanner.Text()

		if strings.Contains(line, "CREATE TABLE") {
			fmt.Println(line)

			//Cut the string by ` and store this in the map, removing the trailing (
			_, entries, _ := strings.Cut(line, "`")
			table_name := strings.ReplaceAll(entries, "` (", "")
			create_tables[table_name+".begin"] = line

			//Proceed with the next line for each field types

			//create a field slice to just get the column name
			for fileScanner.Scan() {
				nextLine := fileScanner.Text()
				fmt.Println(nextLine)
				field_name := strings.Fields(nextLine)

				//Remove backticks for ease of use
				full_name := table_name + "." + strings.ReplaceAll(field_name[0], "`", "")

				//The indexes begin with KEY, increase the array name so it gives the index name
				if strings.Contains(full_name, "KEY") {
					full_name = table_name + "." + "index." + strings.ReplaceAll(field_name[1], "`", "")
				}

				//Replace the db.table.) key with db.table.end , which contains the Engine
				if strings.Contains(full_name, ")") {
					full_name = table_name + "." + strings.ReplaceAll(field_name[0], ")", "end")
				}

				create_tables[full_name] = nextLine

				if strings.Contains(nextLine, "ENGINE=") {
					break
				}

			}
		}

	}

	readFile.Close()

	//First Part of Vschema file, primary vindex algorith defined
	Sharded := BaseVindex{
		Sharded: "true",
		Vindex: HashName{
			HType: HashType{"hash"}},
	}

	vschema_begin, _ := json.MarshalIndent(Sharded, "", "\t")
	fmt.Println(string(vschema_begin))

	//making a slice with enough capacity to hold all the keys
	keys := make([]string, 0, len(create_tables))

	//range over the keys and append them to the slice
	for k := range create_tables {
		keys = append(keys, k)
	}

	//look for auto increment keys to build the sequence vschema
	sort.Strings(keys)
	for _, k := range keys {

		table_name := strings.Split(k, ".")[0]
		column_name := strings.Split(k, ".")[1]

		if strings.Contains(create_tables[k], "AUTO") && !strings.Contains(k, "end") {
			fmt.Println("Found an AUTO_INCREMENT field at: ", k, "  ", create_tables[k])
			//populating the structure
			//edit note, for the vindex I need to add a search aspect and pull the hash name from
			//the base index struc

			MainT := MainTables{
				TableName: table_name,
				Auto:      SeqTable{column_name, table_name + "_seq"},
				Vindex:    Vindex{column_name, "hash_idx"},
			}

			//encoding
			json_sharded, _ := json.MarshalIndent(MainT, "", "\t")
			fmt.Println(string(json_sharded))

		}
	}

}
