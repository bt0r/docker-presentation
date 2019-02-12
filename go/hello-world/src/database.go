package main
import (
  	"github.com/go-pg/pg"
)

func main() {
    db := pg.Connect(&pg.Options{
        User:     "postgres",
        Password: "root",
        Database: "postgres",
        Addr: "db:5432",
    })
    defer db.Close()

    var n int
    db.QueryOne(pg.Scan(&n), "SELECT 1")
    switch(n){
       case 1:
        println("YES CONNECTED")
        break;
        default:
         println("NOT YET")
         break;
    }
}
