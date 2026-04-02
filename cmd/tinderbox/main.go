package main
import ("fmt";"log";"net/http";"os";"github.com/stockyard-dev/stockyard-tinderbox/internal/server";"github.com/stockyard-dev/stockyard-tinderbox/internal/store")
func main(){port:=os.Getenv("PORT");if port==""{port="9670"};dataDir:=os.Getenv("DATA_DIR");if dataDir==""{dataDir="./tinderbox-data"}
db,err:=store.Open(dataDir);if err!=nil{log.Fatalf("tinderbox: %v",err)};defer db.Close();srv:=server.New(db)
fmt.Printf("\n  Tinderbox — CI/CD trigger server\n  Dashboard:  http://localhost:%s/ui\n  API:        http://localhost:%s/api\n\n",port,port)
log.Printf("tinderbox: listening on :%s",port);log.Fatal(http.ListenAndServe(":"+port,srv))}
