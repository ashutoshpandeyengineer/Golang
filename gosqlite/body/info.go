package info
import( "database/sql"

)
type Name struct {                                                       //Declared structs .
     DB *sql.DB	
}
func ReadTable(){                                                      //fetching data from the table.

}
func CreateTable(db *sql.DB) *Name{                                  //creation of a table.
    statement,_ := db.Prepare("Create table people IF NOT EXISTS ( id INTEGER NOT NULL PRIMARY KEY ,   First name TEXT  ,  Last name TEXT)")
	statement.Exec()
	return &Name{DB: db,}
}
func AddColumn(db *sql.DB) *Name{
	statement,_:=db.Prepare("ALTER TABLE people ADD COLUMN contact int(10)")
    statement.Exec()
	return &Name{ DB:db,}
}
func DeleteTable(db *sql.DB) *Name{
	statement,_:=db.Prepare("Drop table people")
    statement.Exec()
	return &Name{ DB:db,}
}
func (name *Name) InsertItem(item Item) {                                                   //(name *Name)-->methods used .
	statement,_ := name.DB.Prepare("INSERT INTO people ( First name ,  Last name) values (?,?)")
	statement.Exec(item.Firstname,item.Lastname)
}
func (name *Name) UpdateItem(item Item) {                                                   //(name *Name)-->methods used .
	statement,_ := name.DB.Prepare("UPDATE people set Lastname=John where id=1")
	statement.Exec(item.Lastname)
}
func (name *Name) DeleteItem(item Item) {                                                   //(name *Name)-->methods used .
	statement,_ := name.DB.Prepare("DELETE from people where id")
	statement.Exec(item.ID,item.Firstname,item.Lastname)
}






