package gracefulexit

import (
	"os"
	"os/signal"
	"sync"
	"syscall"
)

type Handler func()

type GracefulExist struct {
	pb sync.Map
}

// Capture capture program exit signal
// Suitable for k8s pod logout、docker container stop、program exit and etc.
func (r *GracefulExist) Capture() {
	sig := make(chan os.Signal)

	signal.Notify(sig, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM,
		syscall.SIGQUIT)

	for s := range sig {
		switch s {
		case syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT:
			r.pb.Range(func(key, value interface{}) bool {
				handler, ok := value.(Handler)
				if ok {
					handler()
				}
				return true
			})
			return
		}
	}
}

// RegistryHandle register handle
// the handle will be automatically executed when the program exits
func (r *GracefulExist) RegistryHandle(name string, handler Handler) {
	if _, ok := r.pb.Load(name); ok {
		return
	}
	r.pb.Store(name, handler)
}

// DestroyHandle destroy handle
// after destroy, the handle will no longer be executed when the program exits
func (r *GracefulExist) DestroyHandle(name string) {
	if _, ok := r.pb.Load("name"); ok {
		r.pb.Delete(name)
	}
}

// NewGracefulExit init new instance
func NewGracefulExit() *GracefulExist {
	return &GracefulExist{pb: sync.Map{}}
}
