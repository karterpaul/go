package main
import (
	"database/sql"
        "fmt"
        "log"
        _"github.com/lib/pq"
)

func main(){

db, err := sql.Open("postgres", "user=er password=lpass host=localhost port=5432 dbname=db1 sslmode=disable")
if err != nil{
log.Fatal(err)
}

fmt.Println("db=%s",db);

err=db.Ping()
if err!=nil{
  log.Fatal("Error: Could not establish a connection with the database")
}


rows, err := db.Query("SELECT data1,data2,data3  FROM schema1.table order by data1")
if err!=nil{
 log.Println("erro in dbquery.selection failed")
}

for rows.Next() {
        var d1 int
        var d2 string
        var d3 string
        err = rows.Scan(&d1,&d2,&d3)
        if err!=nil{
           log.Println("erro in dbquery.selection failed")
        }
        fmt.Println("d1 | d2 | d3 ")
        fmt.Printf("%3v | %8v | %6v \n", d1,d2,d3)
       } 

}


