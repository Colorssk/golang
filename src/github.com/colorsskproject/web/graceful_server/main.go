package main

import (
	"context"
	"errors"
	"flag"
	"log"
	"net"
	"net/http"
	"os"
	"os/exec"
	"os/signal"
	"syscall"
	"time"
)

var (
	server   *http.Server
	listener net.Listener
	graceful = flag.Bool("graceful", false, "listen on fd open 3 (internal use only)")
)

func handler(w http.ResponseWriter, r *http.Request) {
	time.Sleep(20 * time.Second)
	w.Write([]byte("hello world233333!!!!"))
}

func main() {
	flag.Parse()

	http.HandleFunc("/hello", handler)
	server = &http.Server{Addr: ":9999"}

	var err error
	if *graceful {
		log.Print("main: Listening to existing file descriptor 3.")
		// cmd.ExtraFiles: If non-nil, entry i becomes file descriptor 3+i.
		// when we put socket FD at the first entry, it will always be 3(0+3)
		f := os.NewFile(3, "")              // 0:键盘输出，1:输出2：错误输出，之后都是参数[0],[1]
		listener, err = net.FileListener(f) // 实现子进程监听父进程（9999端口）
	} else {
		log.Print("main: Listening on a new file descriptor.")
		listener, err = net.Listen("tcp", server.Addr) // 不也子进程，正常监听端口
	}

	if err != nil {
		log.Fatalf("listener error: %v", err)
	}

	go func() {
		// server.Shutdown() stops Serve() immediately, thus server.Serve() should not be in main goroutine
		err = server.Serve(listener) // 这里就可以接受用户发送过来的请求
		log.Printf("server.Serve err: %v\n", err)
	}()
	signalHandler()
	log.Printf("signal end")
}

func reload() error {
	tl, ok := listener.(*net.TCPListener)
	if !ok {
		return errors.New("listener is not tcp listener")
	}

	f, err := tl.File()
	if err != nil {
		return err
	}

	args := []string{"-graceful"}
	cmd := exec.Command(os.Args[0], args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	// put socket FD at the first entry
	cmd.ExtraFiles = []*os.File{f}
	return cmd.Start()
}

func signalHandler() {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM, syscall.SIGUSR2)
	for {
		sig := <-ch
		log.Printf("signal: %v", sig)

		// timeout context for shutdown
		ctx, _ := context.WithTimeout(context.Background(), 20*time.Second)
		switch sig {
		case syscall.SIGINT, syscall.SIGTERM: // 入过程序要退出
			// stop
			log.Printf("stop")
			signal.Stop(ch)
			server.Shutdown(ctx) // 保证当前请求处理完之后再退出
			log.Printf("graceful shutdown")
			return
		case syscall.SIGUSR2:
			// reload
			log.Printf("reload")
			err := reload() // 启动子进程
			if err != nil {
				log.Fatalf("graceful restart error: %v", err)
			}
			server.Shutdown(ctx)
			log.Printf("graceful reload")
			return
		}
	}
}
