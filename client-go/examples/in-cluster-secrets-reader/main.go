package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

func main() {
	namespace := os.Getenv("NAMESPACE")
	if len(namespace) == 0 {
		log.Fatal("Env variable NAMESPACE not set\n")
	}
	secretName := os.Getenv("SECRET_NAME")
	if len(secretName) == 0 {
		log.Fatal("Env variable SECRET_NAME not set\n")
	}

	// creates the in-cluster config
	config, err := rest.InClusterConfig()
	if err != nil {
		panic(err.Error())
	}
	// creates the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	secrets := clientset.CoreV1().Secrets(namespace)
	secret, err := secrets.Get(context.TODO(), secretName, metav1.GetOptions{})
	if err != nil {
		panic(err.Error())
	}

	secretData := secret.Data
	for key, value := range secretData {
		// key is string, value is []byte
		fmt.Printf("    %s: %s\n", key, value)
	}

	user := string(secretData["basic-auth-user"])
	password := string(secretData["basic-auth-password"])
	fmt.Printf("basic-auth-user: %v\n", user)
	fmt.Printf("basic-auth-password: %v\n", password)



	url := "http://gateway.openfaas:8080/system/functions"
	method := "GET"

	client := &http.Client{
		Timeout: time.Second * 10,
	}
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		log.Fatal(err)
	}
	req.SetBasicAuth(user, password)
	response, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("response body:\n%s\n", body)
}
