package ormbench

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/donnie4w/gdao"
	"github.com/donnie4w/gdao/base"
	"github.com/donnie4w/gdao/gdaoCache"
	"github.com/donnie4w/gdao/gdaoMapper"
	"github.com/donnie4w/gdaodemo/dao"
	"github.com/donnie4w/simplelog/logging"
	"github.com/donnie4w/test/ormbench/ent"
	hstestEnt "github.com/donnie4w/test/ormbench/ent/hstest"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
	"strconv"
	"testing"
	"time"
)

var p1, p2, p3 = `?`, `?`, `?`
var sqlstr = "select id,age,rowname,updatetime,body,floa,level from hstest where id between " + p1 + " and " + p2 + "  limit " + p3
var args = []any{1, 100, 10}

var postgreql_dns = "host=localhost user=postgres password=123456 dbname=hstest port=5432 sslmode=disable"
var mysql_dns = "root:123@tcp(localhost:3306)/hstest?charset=utf8mb4&parseTime=True&loc=Local"
var sqlite_dns = "hstest.db"
var pgx_dns = "postgresql://postgres:123456@localhost/hstest?sslmode=disable"

var db_mysql, _ = sql.Open("mysql", mysql_dns)
var db_sqlite, _ = sql.Open("sqlite3", sqlite_dns)
var db_postgres, _ = sql.Open("postgres", postgreql_dns)
var db_pgx, _ = sql.Open("pgx", pgx_dns)

var gormmysql, _ = gorm.Open(mysql.Open(mysql_dns), &gorm.Config{})
var gormsqlite, _ = gorm.Open(sqlite.Open(sqlite_dns), &gorm.Config{})
var gormpostgreql, _ = gorm.Open(postgres.Open(postgreql_dns), &gorm.Config{})

var sqlx_mysql, _ = sqlx.Connect("mysql", mysql_dns)
var sqlx_sqlite, _ = sqlx.Connect("sqlite3", sqlite_dns)
var sqlx_postgres, _ = sqlx.Connect("postgres", postgreql_dns)

var ent_mysql, _ = ent.Open("mysql", mysql_dns)
var ent_sqlite, _ = ent.Open("sqlite3", mysql_dns)
var ent_postgres, _ = ent.Open("postgres", postgreql_dns)

var db *sql.DB
var dbtype base.DBType

var gormdb *gorm.DB
var sqlxdb *sqlx.DB

var entdb *ent.Client

func choseDB(i int) {
	switch i {
	case 1: // MYSQL
		db, dbtype = db_mysql, gdao.MYSQL
		gormdb = gormmysql
		sqlxdb = sqlx_mysql
		entdb = ent_mysql
	case 2: //SQLITE
		db, dbtype = db_sqlite, gdao.SQLITE
		gormdb = gormsqlite
		sqlxdb = sqlx_sqlite
		entdb = ent_sqlite
	case 3: // POSTGRESQL
		db, dbtype = db_postgres, gdao.POSTGRESQL
		gormdb = gormpostgreql
		sqlxdb = sqlx_postgres
		p1, p2, p3 = `$1`, `$2`, `$3`
		sqlstr = "select id,age,rowname,updatetime,body,floa,level from hstest where id between " + p1 + " and " + p2 + "  limit " + p3
		entdb = ent_postgres
	case 4: // pgx
		db, dbtype = db_pgx, gdao.POSTGRESQL
		gormdb = gormpostgreql
		sqlxdb = sqlx_postgres
		p1, p2, p3 = `$1`, `$2`, `$3`
		sqlstr = "select id,age,rowname,updatetime,body,floa,level from hstest where id between " + p1 + " and " + p2 + "  limit " + p3
		entdb = ent_postgres
	default:
		panic("no such db type :" + strconv.Itoa(i))
	}
}

func init() {
	choseDB(1)
	gdao.Init(db, dbtype)
	db = gdao.GetDefaultDBHandle().GetDB()
	logging.Debug("check db")
	dbs := gdao.ExecuteQueryBeans(sqlstr, args...)
	logging.Debug("check sql:", dbs.Len())
	logging.Debug("build:", gdaoMapper.Builder("mapper.xml"))
}

func query(sqlstr string, args ...any) {
	rows, err := db.Query(sqlstr, args...)
	if err != nil {
		log.Fatalf("Error executing query: %v", err)
	}
	defer rows.Close()
	names, err := rows.Columns()
	if err != nil {
		log.Fatalf("Error getting column names: %v", err)
	}
	columnValues := make([]interface{}, len(names))
	pointerToColumnValues := make([]any, len(columnValues))
	for i := range columnValues {
		pointerToColumnValues[i] = &columnValues[i]
	}
	for rows.Next() {
		if err := rows.Scan(pointerToColumnValues...); err != nil {
			log.Fatalf("Error scanning row: %v", err)
		}
	}
	if err := rows.Err(); err != nil {
		log.Fatalf("Error iterating over rows: %v", err)
	}
}

func BenchmarkSerial_Native(b *testing.B) {
	for i := 0; i < b.N; i++ {
		query(sqlstr, args...)
	}
}

func BenchmarkSerial_gdao_struct(b *testing.B) {
	for i := 0; i < b.N; i++ {
		hs := dao.NewHstest()
		hs.Where(hs.ID.Between(1, 100)).Limit2(0, 10)
		if _, err := hs.Selects(); err != nil {
			logging.Error(err)
		}
	}
}

func BenchmarkSerial_gdao_sql(b *testing.B) {
	for i := 0; i < b.N; i++ {
		gdao.ExecuteQueryList[dao.Hstest](sqlstr, args...)
	}
}

func BenchmarkSerial_gdao_databean(b *testing.B) {
	for i := 0; i < b.N; i++ {
		gdao.ExecuteQueryBeans(sqlstr, args...)
	}
}

func BenchmarkSerial_gdao_mapper_struct(b *testing.B) {
	for i := 0; i < b.N; i++ {
		gdaoMapper.Selects[dao.Hstest]("user.selectHstestBench", args...)
	}
}

func BenchmarkSerial_gdao_mapper_databean(b *testing.B) {
	for i := 0; i < b.N; i++ {
		gdaoMapper.SelectBeans("user.selectHstestBench", args...)
	}
}

func Benchmark_gdao_cache_struct(b *testing.B) {
	gdaoCache.BindTableNames("hstest")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		hs := dao.NewHstest()
		hs.Where(hs.ID.Between(1, 100)).Limit2(0, 10)
		if _, err := hs.Selects(); err != nil {
			fmt.Println(err)
			break
		}
	}
}

func Benchmark_gdao_cache_mapper(b *testing.B) {
	gdaoCache.BindMapper("user")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if beans := gdaoMapper.SelectBeans("user.selectHstestBench", args...); beans.GetError() != nil {
			logging.Error(beans.GetError())
		}
	}
}

func Benchmark_gdao_cache_mapper_sql(b *testing.B) {
	gdaoCache.BindMapper("user")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		gdaoMapper.Selects[dao.Hstest]("user.selectHstestBench", args)
	}
}

func Benchmark_gdao_struct_reflect(b *testing.B) {
	type Hs1 struct {
		Id         int64
		Age        int64
		Rowname    string
		Value      string
		Updatetime time.Time
		Body       []byte
		Floa       []byte
		Level      int64
	}
	for i := 0; i < b.N; i++ {
		gdao.ExecuteQueryList[Hs1](sqlstr, args...)
	}
}

func Benchmark_gdao_sql_reflect(b *testing.B) {
	type Hs1 struct {
		Id         int64
		Age        int64
		Rowname    string
		Value      string
		Updatetime time.Time
		Body       []byte
		Floa       []byte
		Level      int64
	}
	for i := 0; i < b.N; i++ {
		var hs1s []Hs1
		gdao.ExecuteQueryBeans(sqlstr, args...).Scan(&hs1s)
	}
}

func BenchmarkSerial_gorm_struct(b *testing.B) {
	gormdb.AutoMigrate(&hstest{})
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var hs1s []hstest
		gormdb.Where("id BETWEEN "+p1+" AND "+p2, 1, 100).Limit(10).Find(&hs1s)
	}
}

func BenchmarkSerial_gorm_sql(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var hs1 []hstest
		gormdb.Raw(sqlstr, args...).Scan(&hs1)
	}
}

type hstest struct {
	gorm.Model
	Id         int64
	Age        int64
	Rowname    string
	Value      string
	Updatetime time.Time
	Body       []byte
	Floa       float64
	Level      int64
}

func (hstest) TableName() string {
	return "hstest"
}

func BenchmarkSerial_sqlx_sql(b *testing.B) {
	type Hs1 struct {
		Id         int64
		Age        int64
		Rowname    string
		Value      string
		Updatetime time.Time
		Body       []byte
		Floa       []byte
		Level      int64
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var Hss []Hs1
		sqlxdb.Select(&Hss, sqlstr, args...)
	}
}

func BenchmarkSerial_ent(b *testing.B) {
	//client, err := ent.Open("mysql", mysql_dns)
	//if err != nil {
	//	log.Fatalf("failed opening connection to sqlite: %v", err)
	//}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := entdb.Hstest.
			Query().Where(hstestEnt.IDGTE(1), hstestEnt.IDLTE(100)).
			Limit(10).All(context.TODO())
		if err != nil {
			b.Fatalf("failed querying hstest record: %v", err)
			break
		}
	}
}
