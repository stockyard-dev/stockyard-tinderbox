package store
import ("database/sql";"fmt";"os";"path/filepath";"time";_ "modernc.org/sqlite")
type DB struct{db *sql.DB}
type Build struct{
	ID string `json:"id"`
	Name string `json:"name"`
	Branch string `json:"branch"`
	Commit string `json:"commit_hash"`
	Status string `json:"status"`
	Duration int `json:"duration"`
	Logs string `json:"logs"`
	CreatedAt string `json:"created_at"`
}
func Open(d string)(*DB,error){if err:=os.MkdirAll(d,0755);err!=nil{return nil,err};db,err:=sql.Open("sqlite",filepath.Join(d,"tinderbox.db")+"?_journal_mode=WAL&_busy_timeout=5000");if err!=nil{return nil,err}
db.Exec(`CREATE TABLE IF NOT EXISTS builds(id TEXT PRIMARY KEY,name TEXT NOT NULL,branch TEXT DEFAULT 'main',commit_hash TEXT DEFAULT '',status TEXT DEFAULT 'pending',duration INTEGER DEFAULT 0,logs TEXT DEFAULT '',created_at TEXT DEFAULT(datetime('now')))`)
return &DB{db:db},nil}
func(d *DB)Close()error{return d.db.Close()}
func genID()string{return fmt.Sprintf("%d",time.Now().UnixNano())}
func now()string{return time.Now().UTC().Format(time.RFC3339)}
func(d *DB)Create(e *Build)error{e.ID=genID();e.CreatedAt=now();_,err:=d.db.Exec(`INSERT INTO builds(id,name,branch,commit_hash,status,duration,logs,created_at)VALUES(?,?,?,?,?,?,?,?)`,e.ID,e.Name,e.Branch,e.Commit,e.Status,e.Duration,e.Logs,e.CreatedAt);return err}
func(d *DB)Get(id string)*Build{var e Build;if d.db.QueryRow(`SELECT id,name,branch,commit_hash,status,duration,logs,created_at FROM builds WHERE id=?`,id).Scan(&e.ID,&e.Name,&e.Branch,&e.Commit,&e.Status,&e.Duration,&e.Logs,&e.CreatedAt)!=nil{return nil};return &e}
func(d *DB)List()[]Build{rows,_:=d.db.Query(`SELECT id,name,branch,commit_hash,status,duration,logs,created_at FROM builds ORDER BY created_at DESC`);if rows==nil{return nil};defer rows.Close();var o []Build;for rows.Next(){var e Build;rows.Scan(&e.ID,&e.Name,&e.Branch,&e.Commit,&e.Status,&e.Duration,&e.Logs,&e.CreatedAt);o=append(o,e)};return o}
func(d *DB)Delete(id string)error{_,err:=d.db.Exec(`DELETE FROM builds WHERE id=?`,id);return err}
func(d *DB)Count()int{var n int;d.db.QueryRow(`SELECT COUNT(*) FROM builds`).Scan(&n);return n}
