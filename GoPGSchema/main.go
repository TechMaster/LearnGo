package main


func main() {

	ConnectDb()

	InitSchema()

	InserData()
	defer Db.Close()
}
