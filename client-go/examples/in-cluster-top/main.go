package main

import (
	"context"
	"fmt"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/rest"
	metrics "k8s.io/metrics/pkg/client/clientset/versioned"
	"log"
	"os"
	"time"
)

func main() {

	os.LookupEnv("NAMESPACE")

	// creates the in-cluster config
	config, err := rest.InClusterConfig()
	if err != nil {
		log.Fatal(err.Error())
	}
	metricsClient, err := metrics.NewForConfig(config)
	if err != nil {
		panic(err)
	}

	for {
		podMetrics, err := metricsClient.MetricsV1beta1().PodMetricses("openfaas-fn").
			List(context.TODO(), metav1.ListOptions{})
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		for _, podMetric := range podMetrics.Items {
			podContainers := podMetric.Containers
			for _, container := range podContainers {
				cpuQuantity, ok := container.Usage.Cpu().AsInt64()
				memQuantity, ok := container.Usage.Memory().AsInt64()
				if !ok {
					log.Fatal("Error while retrieving resources' usage\n")
				}
				msg := fmt.Sprintf("Pod Name %s \n Container Name: %s \n  CPU usage: %d \n  Memory usage: %d",
					podMetric.Name, container.Name, cpuQuantity, memQuantity)
				fmt.Println(msg)
			}
		}
		time.Sleep(10 * time.Second)
	}
}
