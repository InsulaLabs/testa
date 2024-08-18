package testa

import (
	"log/slog"
	"math/rand"
	"os"
	"path/filepath"
	"testing"
)

const dbFileNameQT = "query_test.db"

type tc struct {
	key string
	val []byte
}

func gentc(ks int, vs int) tc {
	kt := make([]byte, ks)
	rand.Read(kt)
	vt := make([]byte, vs)
	rand.Read(vt)
	return tc{string(kt), vt}
}

func TestTvksDb(t *testing.T) {

	size := 200
	vals := make([]tc, size)
	for i := 0; i < size; i++ {
		vals[i] = gentc(30, 200)
	}

	lvl := new(slog.LevelVar)
	lvl.Set(slog.LevelDebug)

	logger := slog.New(slog.NewTextHandler(
		os.Stderr,
		&slog.HandlerOptions{Level: lvl}))

	slog.SetDefault(logger)

	tdb, err := Open(
		filepath.Join(t.TempDir(), dbFileNameQT))

	defer tdb.Close()

	if err != nil {
		t.Fatalf(`Error: %v`, err)
	}

	for _, c := range vals {
		err := tdb.Set(c.key, c.val)
		if err != nil {
			t.FailNow()
		}
	}
	for _, c := range vals {
		val, err := tdb.Get(c.key)
		if err != nil {
			t.FailNow()
		}
		if len(val) != len(c.val) || string(val) != string(c.val) {
			t.Fatalf(
				"unexpected val %s != %s", string(c.val), string(val))
		}
	}
}
