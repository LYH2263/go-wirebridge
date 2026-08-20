package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"example.com/wirebridge"
	"example.com/wirebridge/internal/api"
	"example.com/wirebridge/internal/handler"
	"example.com/wirebridge/internal/opcode"
)

func main() {
	addr := flag.String("addr", ":8101", "HTTP 监听地址")
	web := flag.String("web", "web", "静态管理页目录")
	persist := flag.String("persist", "", "路由快照 JSON 路径（可选）")
	bypass := flag.String("bypass-log", "", "旁路试帧日志路径（可选）")
	maxFrame := flag.Int("max-frame", 1<<20, "最大帧长")
	flag.Parse()

	opts := []wirebridge.Option{
		wirebridge.WithMaxFrame(*maxFrame),
		wirebridge.WithName("wired"),
	}
	if *persist != "" {
		opts = append(opts, wirebridge.WithPersistPath(*persist))
	}
	if *bypass != "" {
		opts = append(opts, wirebridge.WithBypassLog(*bypass))
	}

	b := wirebridge.New(opts...)
	defer b.Close()
	_ = b.Register(wirebridge.Opcode(opcode.Ping), "ping", handler.Echo(), "builtin")
	_ = b.Register(wirebridge.Opcode(opcode.Echo), "echo", handler.Echo(), "builtin")
	if *persist != "" {
		if err := b.LoadPersist(); err != nil {
			log.Printf("load persist: %v", err)
		}
	}

	srv := api.New(b, api.Options{WebDir: *web, AllowCORS: true})
	httpSrv := &http.Server{
		Addr:              *addr,
		Handler:           srv,
		ReadHeaderTimeout: 5 * time.Second,
	}
	go func() {
		log.Printf("wired 监听 %s", *addr)
		if err := httpSrv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal(err)
		}
	}()

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt, syscall.SIGTERM)
	<-ch
	_ = httpSrv.Close()
}
