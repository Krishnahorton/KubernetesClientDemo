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
	flag.String("apiserver","","api server uri string")
	flag.String("Token","","valid token for service account")
	flag.Parse()
	config, err := clientcmd.BuildConfigFromFlags("https://192.168.99.106:8443", "")
	// TOKEN=$(kubectl get secrets \
	//    -o jsonpath='{.items[?(@.type=="kubernetes.io/service-account-token")].data.token}' \
	//    | base64 --decode)
	config.BearerToken = "eyJhbGciOiJSUzI1NiIsImtpZCI6ImRVUXI1c1J1ZXFsbHZvTXU4R0NnTllBb2VDcGVuNUoyZzVUS0xBTDYwLWMifQ.eyJpc3MiOiJrdWJlcm5ldGVzL3NlcnZpY2VhY2NvdW50Iiwia3ViZXJuZXRlcy5pby9zZXJ2aWNlYWNjb3VudC9uYW1lc3BhY2UiOiJkZWZhdWx0Iiwia3ViZXJuZXRlcy5pby9zZXJ2aWNlYWNjb3VudC9zZWNyZXQubmFtZSI6ImRlZmF1bHQtdG9rZW4tNWNnYjgiLCJrdWJlcm5ldGVzLmlvL3NlcnZpY2VhY2NvdW50L3NlcnZpY2UtYWNjb3VudC5uYW1lIjoiZGVmYXVsdCIsImt1YmVybmV0ZXMuaW8vc2VydmljZWFjY291bnQvc2VydmljZS1hY2NvdW50LnVpZCI6IjgyMmVmNjU0LWNlNzktNGU1YS1iZWE1LTQyYjBhMmFiNTZmMyIsInN1YiI6InN5c3RlbTpzZXJ2aWNlYWNjb3VudDpkZWZhdWx0OmRlZmF1bHQifQ.psJfP9WdODkzIJ4uvSIGWhkq3bOcDtohgFULdtqZrpG3d5CCRg-nh5LWXkqwH70f8UQKmEk5xE5xwIkKKr5G5wqriuS4V9a_62gAE9Us6VKvR0qZx2P3Xc4xfsh2y9NwFaS0guwrrGSTg607EoweIjAIr1Ae2JNHv8bs8HmzupvrfkO3EGR107mv1USzrNLBGY2T5RY3krbhLQopgYwhEEyRrDoHhTyy_4rAE0kIagP65GbwxHlW7b8txxrxAvy8WDBJ5YAlAssZzPa3FBNk9RYGP1NnyDuQKsAabQt66FseahcWtXZ0LUgyMwyGM0WHZbjmlCFW_qyEtpe8LF7zfA"
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
