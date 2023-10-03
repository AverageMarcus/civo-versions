package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/civo/civogo"
)

var (
	client *civogo.Client
	port   string
)

type ReleaseTemplate struct {
	Releases []civogo.KubernetesVersion `json:"releases"`
}

func init() {
	var err error

	apiKey, ok := os.LookupEnv("API_KEY")
	if !ok {
		panic("API_KEY env var must be provided")
	}

	regionCode, ok := os.LookupEnv("REGION")
	if !ok {
		regionCode = "LON1"
	}

	client, err = civogo.NewClient(apiKey, regionCode)
	if err != nil {
		panic(err)
	}

	if p, ok := os.LookupEnv("PORT"); !ok {
		port = ":8000"
	} else {
		port = fmt.Sprintf(":%s", p)
	}
}

func main() {
	http.HandleFunc("/", handler("", ""))
	http.HandleFunc("/k3s/", handler("k3s", ""))
	http.HandleFunc("/k3s/stable", handler("k3s", "stable"))
	http.HandleFunc("/k3s/development", handler("k3s", "development"))
	http.HandleFunc("/talos/", handler("talos", ""))
	http.HandleFunc("/talos/stable", handler("talos", "stable"))
	http.HandleFunc("/talos/development", handler("talos", "development"))

	http.ListenAndServe(port, nil)
}

func handler(clusterType, versionType string) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		versions, err := getVersions(clusterType, versionType)
		if err != nil {
			w.WriteHeader(500)
			return
		}
		b, err := json.Marshal(versions)
		if err != nil {
			w.WriteHeader(500)
			return
		}

		w.Header().Add("Content-Type", "application/json")
		w.Write(b)
	}
}

func getVersions(clusterType string, versionType string) (ReleaseTemplate, error) {
	releases := ReleaseTemplate{
		Releases: []civogo.KubernetesVersion{},
	}

	allVersions, err := client.ListAvailableKubernetesVersions()
	if err != nil {
		return releases, err
	}

	for _, v := range allVersions {
		if v.Type != "deprecated" &&
			(versionType == "" || v.Type == versionType) &&
			(clusterType == "" || v.ClusterType == clusterType) {
			releases.Releases = append(releases.Releases, v)
		}
	}

	return releases, nil
}
