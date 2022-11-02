package main

import (
	"flag"

	"github.com/golang/glog"
)

func main() {

	glog.MaxSize = 1024 * 1024 * 256 // 256m自动分割
	flag.Parse()

	defer glog.Flush()

	glog.InfoDepth(2, "This is info message: %v", 123)
	glog.Infof("This is info message: %v", 123)

	glog.Warning("This is warning message")
	glog.Warningf("This is warning message: %v", 123)

	glog.Error("This is error message")
	glog.Errorf("This is error message: %v", 123)

	// glog.Fatal("This is fatal message")
	// glog.Fatalf("This is fatal message: %v", 123)

	glog.V(1).Info("LEVEL 1 message")
	glog.V(2).Info("LEVEL 2 message")
	glog.V(3).Info("LEVEL 3 message")
	glog.V(4).Info("LEVEL 4 message")
	glog.V(5).Info("LEVEL 5 message")
	glog.V(6).Info("LEVEL 6 message")
}
