package db

func (mdb *McflyDB) Init() {
	mdb.RunMigrate("down")
	mdb.RunMigrate("up")
}
