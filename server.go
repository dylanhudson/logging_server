package main

import (
    "fmt"
    "log"
    "net/http"
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
)

var db *sql.DB
var err error

func handler(w http.ResponseWriter, r *http.Request) {

    sensors, ok := r.URL.Query()["sensor_id"]
    if !ok || len(sensors[0]) < 1 {
       log.Println("Url Param 'sensor_id' is missing")
       return
    }
    fmt.Println(sensors)
    sensor_id := sensors[0]

    values, ok := r.URL.Query()["value"]

    if !ok || len(values[0]) < 1 {
       log.Println("Url Param 'value' is missing")
       return
    }
    value := values[0]

    insert_data(string(sensor_id), string(value))

    log.Println("Url Param 'sensor_id' is: " + string(sensor_id))

    fmt.Fprintf(w, "%s, %s", string(sensor_id), string(value))

} //end handler


func insert_data(sensor_id, value string) {

    stmt, err := db.Prepare("INSERT test SET datetime=NOW(), sensor_id=?, value=?")
    if err != nil {
      log.Fatal(err)
    }

    res, err := stmt.Exec(sensor_id, value)
    if err != nil {
      log.Fatal(err)
    }

    fmt.Println(res)

} //end insert data


func main() {

    db, err = sql.Open("mysql", "root@tcp(localhost:3306)/go_test")
    defer db.Close()
    if err != nil {
      log.Fatal(err)
    }

    http.HandleFunc("/", handler)

    //start server
    log.Fatal(http.ListenAndServe(":8080", nil))

} //end main
