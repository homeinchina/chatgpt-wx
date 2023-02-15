package bootstrap

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/homeinchina/chatgpt-wx/handler/api"
)

func StartApi() {
	mux := http.NewServeMux()
	mux.Handle("/", &apiHandler{})

	server := &http.Server{
		Addr:    ":8000",
		Handler: mux,
	}

	// 创建系统信号接收器
	done := make(chan os.Signal)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-done

		if err := server.Shutdown(context.Background()); err != nil {
			log.Fatal("Shutdown server:", err)
		}
	}()

	log.Println("Starting HTTP server...")
	err := server.ListenAndServe()
	if err != nil {
		if err == http.ErrServerClosed {
			log.Print("Server closed under request")
		} else {
			log.Fatal("Server closed unexpected")
		}
	}
}

type Ret struct {
	Code  int
	Msg   string
	Reply *string
}
type apiHandler struct{}

func (*apiHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	// q := r.FormValue("q")
	// q := r.PostFormValue("q")
	decoder := json.NewDecoder(r.Body)
	var params map[string]string
	decoder.Decode(&params)
	q := params["q"]
	if q == "" {
		w.Write([]byte("提问信息不能为空！"))
		return
	}
	ret := new(Ret)
	ret.Code = 0
	ret.Msg = "success"
	ret.Reply = api.Handle(q)
	ret_json, _ := json.Marshal(ret)
	w.Write(ret_json)
}
