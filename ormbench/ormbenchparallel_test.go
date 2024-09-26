package ormbench

import (
	"github.com/donnie4w/gdao"
	"github.com/donnie4w/gdao/gdaoMapper"
	"github.com/donnie4w/gdaodemo/dao"
	"github.com/donnie4w/simplelog/logging"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/jackc/pgx/v5/stdlib"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
	"testing"
	"time"
)

func BenchmarkParallel_Native(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			query(sqlstr, args...)
		}
	})
}

func BenchmarkParallel_gdao_struct(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			hs := dao.NewHstest()
			hs.Where(hs.ID.Between(1, 100)).Limit2(0, 10)
			if _, err := hs.Selects(); err != nil {
				logging.Error(err)
			}
		}
	})
}

func BenchmarkParallel_gdao_sql(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			gdao.ExecuteQueryList[dao.Hstest](sqlstr, args...)
		}
	})
}

func BenchmarkParallel_gdao_databean(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			gdao.ExecuteQueryBeans(sqlstr, args...)
		}
	})
}

func BenchmarkParallel_gdao_mapper_struct(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			gdaoMapper.Selects[dao.Hstest]("user.selectHstestBench", args...)
		}
	})
}

func BenchmarkParallel_gdao_mapper_databean(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			gdaoMapper.SelectBeans("user.selectHstestBench", args...)
		}
	})
}

func BenchmarkParallel_gorm_struct(b *testing.B) {
	gormdb.AutoMigrate(&hstest{})
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			var hs1s []hstest
			gormdb.Where("id BETWEEN "+p1+" AND "+p2, 1, 100).Limit(10).Find(&hs1s)
		}
	})

}

func BenchmarkParallel_gorm_sql(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			var hs1 []hstest
			gormdb.Raw(sqlstr, args...).Scan(&hs1)
		}
	})
}

func BenchmarkParallel_sqlx_sql(b *testing.B) {
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
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			var Hss []Hs1
			sqlxdb.Select(&Hss, sqlstr, args...)
		}
	})
}
