package settings

type DatabaseSettings struct {
	Host   string
	Port   int
	User   string
	Dbname string
}

var (
	TestAppDatabase DatabaseSettings = DatabaseSettings{
		Host:   "localhost",
		Port:   5432,
		User:   "postgres",
		Dbname: "testkeycut",
	}
)
