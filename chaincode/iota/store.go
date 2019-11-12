package iota

import (
	"encoding/json"
	"fmt"
 	badger "github.com/dgraph-io/badger"
)

type readfunc func(value []byte) error

func Store(data string, table string) {
	db, err := badger.Open(badger.DefaultOptions("./iota/store"))
	if err != nil {
		fmt.Println("error", err)
	}
	defer db.Close()

	Update(db, table, data)
}

func Read(table string) string {
	db, err := badger.Open(badger.DefaultOptions("./iota/store"))
	if err != nil {
		fmt.Println("error", err)
	}
	defer db.Close()

	var result string
	if err := ReadFromStorage(db, table, func(value []byte) error {
		result = string(value)
		return nil
	}); err != nil {
		fmt.Println("error", err)
	}
	return result
}

func Update(db *badger.DB, key string, value string) error {
  err := db.Update(func(txn *badger.Txn) error {
    err := txn.Set([]byte(key), []byte(value))
    return err
  })

  if err != nil {
    return err
  }
  return nil
}

func ReadFromStorage(db *badger.DB, key string, readFunc readfunc) error {
	return db.View(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte(key))
		if err == badger.ErrKeyNotFound {
			return readFunc(nil)
		}
		if err != nil {
			return err
		}
		
    	value, err := item.ValueCopy(nil)

		if err != nil {
			return err
		}
		return readFunc(value)
	})
}





func StoreJSON() {
	db, err := badger.Open(badger.DefaultOptions("./iota/store"))
	if err != nil {
		fmt.Println("error", err)
	}
	defer db.Close()

	var writevalue = ConvertToString("John", 33, "engineer", "travel", "music")
	// const writevalue = 777
	// Update(db, "answer", strconv.Itoa(writevalue))
	Update(db, "mamstate", writevalue)

	var result interface{}
	// var result string
	if err := ReadFromStorage(db, "mamstate", func(value []byte) error {
		result = ConvertToJSON(value)
		// result = string(value)
		return nil
	}); err != nil {
		fmt.Println("error", err)
	}
	// if result != "" {
	if result != nil {
		fmt.Println("result")
		fmt.Println(result)
	}

}

func ConvertToString(name string, age int, work string, hobby1 string, hobby2 string) string {
    // ********************* Marshal *********************
    u := map[string]interface{}{}
    u["name"] = name
    u["age"] = age
    u["work"] = work
    u["hobbies"] = []string{hobby1, hobby2}

    b, err := json.Marshal(u)
    if err != nil {
        panic(err)
    }
    fmt.Println(string(b))
	return string(b)
}

func ConvertToJSON(value []byte) interface{} {
    // ********************* Unmarshal *********************
    var a interface{}
    var err = json.Unmarshal(value, &a)
    if err != nil {
		fmt.Println("error:", err)
		return nil
    }
	return a
}