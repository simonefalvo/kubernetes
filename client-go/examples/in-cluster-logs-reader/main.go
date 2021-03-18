package main

import (
	"bufio"
	"context"
	"fmt"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"log"
	"os"
	"regexp"
	"time"
)


func main() {

	namespace, ok := os.LookupEnv("NAMESPACE")
	if !ok {
		log.Fatal("$NAMESPACE not set")
	}

	// creates the in-cluster config
	config, err := rest.InClusterConfig()
	if err != nil {
		log.Fatal(err.Error())
	}

	// creates the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatal(err.Error())
	}

	re := regexp.MustCompile("gateway")
	scaleLine := fmt.Sprintf(`\[Scale\] function=%s 0 => 1 successful`, "nodeinfo")
	scaleRe := regexp.MustCompile(scaleLine)

	podLogOpts := v1.PodLogOptions{Container: "gateway"}

	for {
		pods, err := clientset.CoreV1().Pods(namespace).List(context.TODO(), metav1.ListOptions{})
		if err != nil {
			log.Fatal(err.Error())
		}


		for _, pod := range pods.Items {
			podName := pod.Name
			if re.MatchString(podName) {
				req := clientset.CoreV1().Pods(namespace).GetLogs(podName, &podLogOpts)
				podLogs, err := req.Stream(context.TODO())
				if err != nil {
					log.Fatal(err.Error())
				}

				//buf := new(bytes.Buffer)
				//_, err = io.Copy(buf, podLogs)
				//if err != nil {
				//	log.Fatal("error in copy information from podLogs to buf")
				//}
				//str := buf.String()
				//fmt.Printf("%s logs:\n", podName)
				//fmt.Println(str)

				scanner := bufio.NewScanner(podLogs)
				for scanner.Scan() {
					line := scanner.Text()
					if scaleRe.MatchString(line) {
						fmt.Println(line) // Println will add back the final '\n'
					}
				}
				err = scanner.Err()
				if err != nil {
					log.Fatal(err)
				}

				err = podLogs.Close()
				if err != nil {
					log.Fatal(err.Error())
				}
			}
		}
		time.Sleep(20 * time.Second)
	}
}
