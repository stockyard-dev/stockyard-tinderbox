package server
import("encoding/json";"net/http";"os/exec";"strconv";"time";"github.com/stockyard-dev/stockyard-tinderbox/internal/store")
func(s *Server)handleList(w http.ResponseWriter,r *http.Request){list,_:=s.db.List();if list==nil{list=[]store.Pipeline{}};writeJSON(w,200,list)}
func(s *Server)handleCreate(w http.ResponseWriter,r *http.Request){var p store.Pipeline;json.NewDecoder(r.Body).Decode(&p);if p.Name==""{writeError(w,400,"name required");return};s.db.Create(&p);writeJSON(w,201,p)}
func(s *Server)handleDelete(w http.ResponseWriter,r *http.Request){id,_:=strconv.ParseInt(r.PathValue("id"),10,64);s.db.Delete(id);writeJSON(w,200,map[string]string{"status":"deleted"})}
func(s *Server)handleRuns(w http.ResponseWriter,r *http.Request){id,_:=strconv.ParseInt(r.PathValue("id"),10,64);list,_:=s.db.ListRuns(id);if list==nil{list=[]store.Run{}};writeJSON(w,200,list)}
func(s *Server)handleTrigger(w http.ResponseWriter,r *http.Request){id,_:=strconv.ParseInt(r.PathValue("id"),10,64);pipes,_:=s.db.List();var script string;for _,p:=range pipes{if p.ID==id{script=p.Script;break}};run:=&store.Run{PipelineID:id,StartedAt:time.Now()};start:=time.Now();var out[]byte;var err error;if script!=""{cmd:=exec.Command("sh","-c",script);out,err=cmd.CombinedOutput()};run.DurationMs=time.Since(start).Milliseconds();run.Output=string(out);if err!=nil{run.Status="failed"}else{run.Status="success"};s.db.RecordRun(run);writeJSON(w,200,run)}
func(s *Server)handleOverview(w http.ResponseWriter,r *http.Request){m,_:=s.db.Stats();writeJSON(w,200,m)}
