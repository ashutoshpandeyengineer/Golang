package main
import ("database/sql"
    _"github.com/mattn/go-sqlite3" 
//	 "Golang/gosqlite/"
	 "fmt"
)
func main() {
	db,_ := sql.Open("sqlite3","./base.db" )
//	statement,_ := db.Prepare("Create table people ( id INTEGER NOT NULL PRIMARY KEY ,   Frist name TEXT  ,  Last name TEXT)")
//	statement.Exec()
    name :=  info.InsertItem(db)
    name.Add(info.Item{Fristname : "Hy" , Lastname  : "Pandey" })
//   name.Add(info.Item{Fristname : "Hi" , Lastname  : "Pandey" })
//    items:=feed.Get()
//   fmt.Println(items)	
//  statement,_ := db.Prepare("CREATE TABLE IF NOT EXISTS newsfeed(id INTEGER PRIMARY KEY , fristname TEXT , lastname TEXT)")
//  statement.Exec()




}