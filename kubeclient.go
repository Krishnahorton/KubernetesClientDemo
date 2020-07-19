package main

import (
	"context"
	"flag"
	"fmt"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"os"
)

func main() {
	apiServer := flag.String("apiserver","","api server uri string")
	token := flag.String("Token","","valid token for service account")
	flag.Parse()
	config, err := clientcmd.BuildConfigFromFlags(*apiServer, "")
	// TOKEN=$(kubectl get secrets \
	//    -o jsonpath='{.items[?(@.type=="kubernetes.io/service-account-token")].data.token}' \
	//    | base64 --decode)
	config.BearerToken = *token
	config.Insecure = true

	if err != nil {
		panic(err.Error())
	}

	clientset,err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}
	// retrieving kubernetes namespace
	namespaces,err := clientset.CoreV1().Namespaces().List(context.TODO(),metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("##### Printing namespaces and Pods ######")
	for _,namespace := range namespaces.Items {
		// list pods in each namespace
		pods,_ := clientset.CoreV1().Pods(namespace.Name).List(context.TODO(),metav1.ListOptions{})
		for _,pod := range pods.Items {
			fmt.Fprintf(os.Stdout, "namespace: %v pod: %v\n", namespace.Name, pod.Name)
		}
	}


}
