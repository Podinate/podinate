package apiclient

import (
	"context"

	"github.com/johncave/podinate/lib/api_client"
	"github.com/spf13/viper"
)

type Pod struct {
	Project *Project
	// The short name (slug/url) of the pod
	ID string `json:"id"`
	// The name of the pod
	Name        string           `json:"name"`
	Image       string           `json:"image"`
	Tag         string           `json:"tag"`
	Status      string           `json:"status"`
	CreatedAt   string           `json:"created_at"`
	ResourceID  string           `json:"resource_id"`
	Volumes     VolumeSlice      `json:"volumes"`
	Services    ServiceSlice     `json:"services"`
	Environment EnvironmentSlice `json:"environment"`
}

// GetPodByID returns a pod by ID from the given project
func (p *Project) GetPodByID(id string) (*Pod, error) {
	resp, _, err := C.PodApi.ProjectProjectIdPodPodIdGet(context.Background(), p.ID, id).Account(viper.GetString("account")).Execute()
	if err != nil {
		return nil, err
	}
	return getPodFromApi(p, resp), nil

}

// GetPods returns all pods from the given project
func (p *Project) GetPods() ([]*Pod, error) {
	resp, _, err := C.PodApi.ProjectProjectIdPodGet(context.Background(), p.ID).Account(viper.GetString("account")).Execute()
	if err != nil {
		return nil, err
	}

	var pods []*Pod

	for _, i := range resp.Items {
		po := i.Pod

		pods = append(pods, getPodFromApi(p, po))
	}

	return pods, nil
}

func getPodFromApi(p *Project, in *api_client.Pod) *Pod {
	//fmt.Println("in.Id", in.Id, "in", in, "created", in.CreatedAt)
	//fmt.Println("%+V\n", in)
	out := &Pod{
		ID:          *in.Id,
		Name:        in.Name,
		Image:       in.Image,
		Tag:         in.Tag,
		Status:      *in.Status,
		ResourceID:  *in.ResourceId,
		Project:     p,
		Volumes:     volumesFromAPI(in.Volumes),
		Services:    servicesFromAPI(in.Services),
		Environment: environmentVariablesFromAPI(in.Environment),
	}

	return out
}

// getLogs returns the logs for a pod
func (p *Pod) GetLogs(lines int, follow bool) (string, error) {
	resp, _, err := C.PodApi.ProjectProjectIdPodPodIdLogsGet(context.Background(), p.Project.ID, p.ID).Account(viper.GetString("account")).Lines(int32(lines)).Execute()
	return resp, err
}

// Delete deletes the pod
func (p *Pod) Delete() error {
	_, err := C.PodApi.ProjectProjectIdPodPodIdDelete(context.Background(), p.Project.ID, p.ID).Account(viper.GetString("account")).Execute()
	return err
}
