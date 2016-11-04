package prefixed_db

import (
	"database/sql"
	"log"
	"strings"
	"sync"
	"testing"

	"github.com/go-sql-driver/mysql"
	"github.com/letsencrypt/boulder/test/vars"
)

func TestPrefixing(t *testing.T) {
	sql.Register("prefixedmysql", New("SET STATEMENT max_statement_time=0.1 FOR ", mysql.MySQLDriver{}))
	db, err := sql.Open("prefixedmysql", vars.DBConnSA)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			_, err := db.Exec("SELECT 1 FROM (SELECT SLEEP(?)) as subselect;", i)
			if err == nil || !strings.HasPrefix(err.Error(), "Error 1969:") {
				t.Error("Expected to get Error 1969 (timeout), got", err)
			}
			wg.Done()
		}()
	}
	wg.Wait()
}
