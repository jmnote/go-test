package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	err := loadKubeconfig()
	if err != nil {
		log.Fatalf("loadKubeconfig err: %s", err)
	}
	fmt.Println("loadKubeconfig ok")
}

func loadKubeconfig() error {
	kubeconfig := os.Getenv("KUBECONFIG")
	if kubeconfig == "" {
		kubeconfig = "exists.yaml"
	}
	_, err := ioutil.ReadFile(kubeconfig)
	if err != nil {
		return fmt.Errorf("ReadFile err: %s", err)
	}
	return nil
}
