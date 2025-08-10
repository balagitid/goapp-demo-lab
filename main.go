package main

import (
	"fmt"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	// Get environment variables
	podName := os.Getenv("POD_NAME")
	podNamespace := os.Getenv("POD_NAMESPACE")
	nodeName := os.Getenv("NODE_NAME")
	clusterName := os.Getenv("CLUSTER_NAME")
	containerImage := os.Getenv("CONTAINER_IMAGE")
	containerTag := os.Getenv("CONTAINER_TAG")

	// Get Pod IP and Hostname
	podIP := os.Getenv("POD_IP")
	hostIP := os.Getenv("HOST_IP")
	hostname, _ := os.Hostname()

	// Simple HTML output
	fmt.Fprintf(w, `
	<html>
		<head><title>KubeApp Demo</title></head>
		<body style="font-family: Arial; background-color: #f2f2f2; padding: 20px;">
			<h1 style="color: #2E86C1;">KubeApp - Kubernetes Info</h1>
			<ul>
				<li><strong>Pod Name:</strong> %s</li>
				<li><strong>Namespace:</strong> %s</li>
				<li><strong>Node Name:</strong> %s</li>
				<li><strong>Cluster Name:</strong> %s</li>
				<li><strong>Pod IP:</strong> %s</li>
				<li><strong>Host IP:</strong> %s</li>
				<li><strong>Hostname:</strong> %s</li>
				<li><strong>Container Image:</strong> %s</li>
				<li><strong>Container Tag:</strong> %s</li>
			</ul>
		</body>
	</html>
	`, podName, podNamespace, nodeName, clusterName, podIP, hostIP, hostname, containerImage, containerTag)
}

func main() {
	http.HandleFunc("/", handler)
	port := "8080"
	fmt.Println("KubeApp is running on port", port)
	http.ListenAndServe(":"+port, nil)
}

