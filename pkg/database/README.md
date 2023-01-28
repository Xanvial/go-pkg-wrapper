### Sample usage

```
dbConn, err := database.NewGorm(database.Config{
    DatabaseType: database.DBTypePostgresql,
    Host:         "localhost",
    Username:     "postgres",
    Password:     "postgres",
    DatabaseName: "zero_setting",
})
if err != nil {
    log.Println("err:", err)
    return
}
// dbConn, err := database.NewSqlxDB(database.Config{
// 	DatabaseType: database.DBTypePostgresql,
// 	Host:         "localhost",
// 	Username:     "postgres",
// 	Password:     "postgres",
// 	DatabaseName: "zero_setting",
// })
// if err != nil {
// 	log.Println("err:", err)
// 	return
// }

type TestTable struct {
    ID          int    `json:"id" db:"id"`
    TestName    string `json:"name" db:"test_name"`
    Description string `json:"desc" db:"description"`
}

err = dbConn.Exec("INSERT INTO test_table (test_name, description) VALUES ($1, $2) RETURNING id",
    "input name", "sample data")
log.Println("err:", err)

retID, err := dbConn.ExecReturn("INSERT INTO test_table (test_name, description) VALUES ($1, $2) RETURNING id",
    "sample name", "sample data with last inserted id")
log.Println("err:", err)
log.Println("retID:", retID)

var output TestTable
err = dbConn.QueryRow("select * from test_table where id = $1", &output, 1)
log.Println("err:", err)
log.Println("output:", output)

var result []TestTable
err = dbConn.Query("select * from test_table where description ilike '%'||$1||'%'", &result, `data`)
log.Println("err:", err)
log.Println("result:", result)
for _, val := range result {
    log.Printf("val:%+v\n", val)
}
```