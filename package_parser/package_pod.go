package package_parser

import (
	"context"
	"fmt"
	"net/url"
	"reflect"
	"strings"

	stderrors "errors"

	hcldec "github.com/hashicorp/hcl2/hcldec"
	"github.com/podinate/podinate/kube_client"
	"github.com/sirupsen/logrus"
	"github.com/zclconf/go-cty/cty"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	networkingv1 "k8s.io/api/networking/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
	"k8s.io/client-go/kubernetes"
)

const (
	KubernetesKindStatefulSet       = "StatefulSet"
	KubernetesAPIVersionStatefulSet = "apps/v1"
	KubernetesKindService           = "Service"
	KubernetesAPIVersionService     = "v1"
	KubernetesKindIngress           = "Ingress"
	KubernetesAPIVersionIngress     = "networking.k8s.io/v1"
)

type Pod struct {
	ID string
	// Name        string   `cty:"name"`
	Namespace   *string  `cty:"namespace"`
	Image       string   `cty:"image"`
	Tag         *string  `cty:"tag"`
	Command     []string `cty:"command"`
	Arguments   []string `cty:"arguments"`
	Environment map[string]struct {
		Value  string `cty:"value"`
		Secret *bool  `cty:"secret"`
	} `cty:"environment"`
	Services map[string]struct {
		Port       int     `cty:"port"`
		TargetPort *int    `cty:"target_port"`
		Protocol   *string `cty:"protocol"`
		URL        *string `cty:"url"`
	} `cty:"service"`

	// Removed for now, because these are separate Kube resources
	Volume map[string]struct {
		Size      int     `cty:"size"`
		MountPath string  `cty:"mount_path"`
		Class     *string `cty:"class"`
	} `cty:"volume"`
	// SharedVolume *[]struct {
	// 	VolumeID string `cty:"volume_id"`
	// 	Path     string `cty:"path"`
	// } `cty:"shared_volume"`
}

// GetHCLSpect returns the HCL spec of the pod type
var podHCLSpec = &hcldec.BlockMapSpec{
	TypeName:   "pod",
	LabelNames: []string{"id"},
	Nested: &hcldec.ObjectSpec{
		// "name": &hcldec.AttrSpec{
		// 	Name:     "name",
		// 	Type:     cty.String,
		// 	Required: true,
		// },
		"image": &hcldec.AttrSpec{
			Name:     "image",
			Type:     cty.String,
			Required: true,
		},
		"tag": &hcldec.AttrSpec{
			Name:     "tag",
			Type:     cty.String,
			Required: false,
		},
		"command": &hcldec.AttrSpec{
			Name:     "command",
			Type:     cty.List(cty.String),
			Required: false,
		},
		"arguments": &hcldec.AttrSpec{
			Name:     "arguments",
			Type:     cty.List(cty.String),
			Required: false,
		},
		"namespace": &hcldec.AttrSpec{
			Name:     "namespace",
			Type:     cty.String,
			Required: false,
		},
		"environment": &hcldec.BlockMapSpec{
			TypeName:   "environment",
			LabelNames: []string{"key"},
			Nested: &hcldec.ObjectSpec{
				"value": &hcldec.AttrSpec{
					Name:     "value",
					Type:     cty.String,
					Required: true,
				},
				"secret": &hcldec.AttrSpec{
					Name:     "secret",
					Type:     cty.Bool,
					Required: false,
				},
			},
		},
		"service": &hcldec.BlockMapSpec{
			TypeName:   "service",
			LabelNames: []string{"name"},
			Nested: &hcldec.ObjectSpec{
				// The listening port of the service
				"port": &hcldec.AttrSpec{
					Name:     "port",
					Type:     cty.Number,
					Required: true,
				},
				// The service on the container that the port is forwarded to
				"target_port": &hcldec.AttrSpec{
					Name:     "target_port",
					Type:     cty.Number,
					Required: false,
				},
				// Protocol of the service. Defaults to TCP
				// Can be set to UDP
				// If set to "http" and url is specified, an Ingress will be created
				"protocol": &hcldec.AttrSpec{
					Name:     "protocol",
					Type:     cty.String,
					Required: false,
				},
				// The url to make the service available at (make the service externally available)
				// If protocol is set to "http" or "https", an Ingress will be created
				// If a domain name followed by a path, the service will be available at that path on the ingress
				"url": &hcldec.AttrSpec{
					Name:     "url",
					Type:     cty.String,
					Required: false,
				},
			},
		},
		"volume": &hcldec.BlockMapSpec{
			TypeName:   "volume",
			LabelNames: []string{"name"},
			Nested: &hcldec.ObjectSpec{
				"size": &hcldec.AttrSpec{
					Name:     "size",
					Type:     cty.Number,
					Required: true,
				},
				"mount_path": &hcldec.AttrSpec{
					Name:     "mount_path",
					Type:     cty.String,
					Required: true,
				},
				"class": &hcldec.AttrSpec{
					Name:     "class",
					Type:     cty.String,
					Required: false,
				},
			},
		},
		// "shared_volume": &hcldec.BlockListSpec{
		// 	TypeName: "shared_volume",
		// 	MinItems: 0,
		// 	Nested: &hcldec.ObjectSpec{
		// 		"volume_id": &hcldec.AttrSpec{
		// 			Name:     "volume_id",
		// 			Type:     cty.String,
		// 			Required: true,
		// 		},
		// 		"path": &hcldec.AttrSpec{
		// 			Name:     "path",
		// 			Type:     cty.String,
		// 			Required: true,
		// 		},
		// 	},
		// },
	},
}

// GetResources returns the resources needed for the Pod
// This function deals with the StatefulSet, then calls other functions to add services, shared volumes, etc.
func (p *Pod) GetResources(ctx context.Context, pkg *Package) (*ChangeType, []ResourceChange, error) {
	var out []ResourceChange
	podChangeType := ChangeTypeNoop

	var imageID = p.Image
	if p.Tag != nil {
		imageID = imageID + ":" + *p.Tag
	}

	ssSpec := &appsv1.StatefulSet{

		ObjectMeta: metav1.ObjectMeta{
			Name:      p.ID,
			Namespace: pkg.Namespace,
		},
		Spec: appsv1.StatefulSetSpec{
			Replicas: func(val int32) *int32 { return &val }(1),
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{
					"podinate.com/pod": p.ID,
				},
			},
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{
						"podinate.com/pod": p.ID,
					},
					Name: p.ID,
				},
				Spec: corev1.PodSpec{
					Containers: []corev1.Container{
						{
							Name:    p.ID,
							Image:   imageID,
							Command: p.Command,
							Args:    p.Arguments,
						},
					},
				},
			},
		},
	}

	ssSpec.Kind = KubernetesKindStatefulSet
	ssSpec.APIVersion = KubernetesAPIVersionStatefulSet

	// Add environment variables
	for k, v := range p.Environment {
		ssSpec.Spec.Template.Spec.Containers[0].Env = append(ssSpec.Spec.Template.Spec.Containers[0].Env, corev1.EnvVar{
			Name:  k,
			Value: v.Value,
		})
	}

	// Add service ports to the StatefulSet
	// This port is just an annotation to help admins know what the container does
	for k, v := range p.Services {
		//ssSpec.Spec.ServiceName = k
		port := corev1.ContainerPort{
			Name:          k,
			ContainerPort: int32(v.Port),
		}

		if v.TargetPort != nil {
			port.ContainerPort = int32(*v.TargetPort)
		}
		ssSpec.Spec.Template.Spec.Containers[0].Ports = append(ssSpec.Spec.Template.Spec.Containers[0].Ports, port)

	}

	// Add volumes to the StatefulSet
	for k, volume := range p.Volume {
		ssSpec.Spec.Template.Spec.Containers[0].VolumeMounts = append(ssSpec.Spec.Template.Spec.Containers[0].VolumeMounts, corev1.VolumeMount{
			Name:      k,
			MountPath: volume.MountPath,
		})
		// Add volume claim templates
		newPVC := corev1.PersistentVolumeClaim{
			ObjectMeta: metav1.ObjectMeta{
				Name: k,
				Annotations: map[string]string{
					"volumeType": "local",
				},
				Labels: map[string]string{
					"podinate.com/pod": p.ID,
				},
			},
			Spec: corev1.PersistentVolumeClaimSpec{
				AccessModes: []corev1.PersistentVolumeAccessMode{
					"ReadWriteOnce",
				},
				//StorageClassName: func(val string) *string { return &val }("local-path"),
				Resources: corev1.VolumeResourceRequirements{
					Requests: corev1.ResourceList{
						"storage": func(val string) resource.Quantity { return resource.MustParse(val) }(fmt.Sprintf("%dGi", volume.Size)),
					},
				},
			},
		}

		if volume.Class != nil {
			// TODO: Check if the SC exists
			newPVC.Spec.StorageClassName = volume.Class
		}

		ssSpec.Spec.VolumeClaimTemplates = append(ssSpec.Spec.VolumeClaimTemplates, newPVC)
	}

	// Add service changes
	serviceChangeType, svcChanges, err := p.GetServiceResourceChanges(ctx)
	if err != nil {
		logrus.WithContext(ctx).WithFields(logrus.Fields{
			"error":           err,
			"service_changes": svcChanges,
			"service_ct":      serviceChangeType,
		}).Error("Error getting service changes")
		return nil, nil, err
	}
	logrus.WithContext(ctx).WithFields(logrus.Fields{
		"service_changes": svcChanges,
		"service_ct":      serviceChangeType,
	}).Trace("Service changes")

	//changeType = *ct
	out = append(out, svcChanges...)

	// Check if the StatefulSet already exists
	kube, err := kube_client.Client()
	if err != nil {
		return nil, nil, err
	}

	ss, err := kube.AppsV1().StatefulSets(pkg.Namespace).Update(ctx, ssSpec, metav1.UpdateOptions{DryRun: []string{metav1.DryRunAll}})
	ss.Kind = KubernetesKindStatefulSet
	ss.APIVersion = KubernetesAPIVersionStatefulSet

	logrus.WithContext(ctx).WithFields(logrus.Fields{
		"statefulset": ss,
		"kind":        ss.GetObjectKind(),
		"error":       err,
	}).Debug("StatefulSet")

	// If the StatefulSet didn't exist when we called update, create it
	if errors.IsNotFound(err) {

		logrus.WithContext(ctx).WithFields(logrus.Fields{
			"statefulset": ss,
		}).Error("Tried to get statefulset, got not found error")

		// Resource doesn't exist, so create it
		podChangeType = ChangeTypeCreate

		// Dry Run creation to fill out default fields
		// Disabled for now, fails if the Namespace isn't already created
		// ss, err := kube.AppsV1().StatefulSets(pkg.Namespace).Create(ctx, ssSpec, metav1.CreateOptions{DryRun: []string{metav1.DryRunAll}})
		// if err != nil {
		// 	return nil, nil, err
		// }

		out = append(out, ResourceChange{
			ChangeType:      ChangeTypeCreate,
			CurrentResource: nil,
			DesiredResource: ssSpec,
		})
	} else if err != nil {
		logrus.WithContext(ctx).WithFields(logrus.Fields{
			"error": err,
		}).Error("Error getting StatefulSet")

		return nil, nil, err
	} else {
		existing, err := kube.AppsV1().StatefulSets(pkg.Namespace).Get(ctx, p.ID, metav1.GetOptions{})
		if err != nil {
			return nil, nil, err
		}
		existing.Kind = KubernetesKindStatefulSet
		existing.APIVersion = KubernetesAPIVersionStatefulSet

		logrus.WithContext(ctx).WithFields(logrus.Fields{
			"statefulset": ss,
		}).Debug("StatefulSet exists, deep comparing")
		// Deep compare the object to see if it needs updating
		if reflect.DeepEqual(ss.Spec, existing.Spec) {
			logrus.WithContext(ctx).WithFields(logrus.Fields{
				"statefulset": ss.Spec,
				"existing":    existing.Spec,
			}).Debug("StatefulSet is up to date")
			//changeType = ChangeTypeNoop
		} else {
			logrus.WithContext(ctx).WithFields(logrus.Fields{
				"statefulset": ss,
				"new-kind":    ss.GetObjectKind(),
				"old-kind":    existing.GetObjectKind(),
				"existing":    existing,
			}).Debug("StatefulSet needs updating")

			// Resource exists, but needs updating
			out = append(out, ResourceChange{
				ChangeType:      ChangeTypeUpdate,
				CurrentResource: existing,
				DesiredResource: ss,
			})
			podChangeType = ChangeTypeUpdate
		}
	}

	// Decide the overall change type
	changeType := ChangeTypeNoop
	if podChangeType == ChangeTypeCreate || podChangeType == ChangeTypeUpdate {
		changeType = podChangeType
	} else if (*serviceChangeType == ChangeTypeCreate || *serviceChangeType == ChangeTypeUpdate) && podChangeType == ChangeTypeNoop {
		changeType = ChangeTypeUpdate
	}

	return &changeType, out, nil
}

func (p *Pod) GetServiceResourceChanges(ctx context.Context) (*ChangeType, []ResourceChange, error) {
	var out []ResourceChange
	changeType := ChangeTypeNoop

	for name, service := range p.Services {
		svcSpec := &corev1.Service{
			ObjectMeta: metav1.ObjectMeta{
				Name:      name,
				Namespace: *p.Namespace,
			},
			Spec: corev1.ServiceSpec{
				Selector: map[string]string{
					"podinate.com/pod": p.ID,
				},
				Ports: []corev1.ServicePort{
					{
						Port:     int32(service.Port),
						Protocol: corev1.ProtocolTCP,
					},
				},
			},
		}

		svcSpec.Kind = KubernetesKindService
		svcSpec.APIVersion = KubernetesAPIVersionService

		if service.Protocol != nil && strings.ToLower(*service.Protocol) == "udp" {
			svcSpec.Spec.Ports[0].Protocol = corev1.ProtocolUDP
		}

		if service.TargetPort != nil {
			svcSpec.Spec.Ports[0].TargetPort = intstr.FromInt(*service.TargetPort)
		}

		logrus.WithContext(ctx).WithFields(logrus.Fields{
			"service": svcSpec,
		}).Trace("generated service spec")

		// Check if the Service already exists
		kube, err := kube_client.Client()
		if err != nil {
			return nil, nil, err
		}

		//s, err := kube.CoreV1().Services(*p.Namespace).Update(ctx, svc, metav1.UpdateOptions{DryRun: []string{metav1.DryRunAll}})
		s, err := kube.CoreV1().Services(*p.Namespace).Get(ctx, name, metav1.GetOptions{})
		s.Kind = KubernetesKindService
		s.APIVersion = KubernetesAPIVersionService
		if errors.IsNotFound(err) {
			// Resource doesn't exist, so create it

			// TODO: Try changing this Update to Create
			svc, err := kube.CoreV1().Services(*p.Namespace).Update(ctx, svcSpec, metav1.UpdateOptions{DryRun: []string{metav1.DryRunAll}})
			if errors.IsNotFound(err) {
				svc = svcSpec
			} else if err != nil {
				return nil, nil, err
			}
			svc.Kind = KubernetesKindService
			svc.APIVersion = KubernetesAPIVersionService

			logrus.WithContext(ctx).WithFields(logrus.Fields{
				"service": svc,
				"s":       s,
			}).Debug("create service")

			changeType = ChangeTypeCreate
			out = append(out, ResourceChange{
				ChangeType:      ChangeTypeCreate,
				CurrentResource: nil,
				DesiredResource: svcSpec,
			})
		} else if err != nil {
			logrus.WithFields(logrus.Fields{
				"error": err,
			}).Error("Error getting Service")

			return nil, nil, err
		} else {
			// Resource exists, but needs updating
			svc, err := kube.CoreV1().Services(*p.Namespace).Update(ctx, svcSpec, metav1.UpdateOptions{DryRun: []string{metav1.DryRunAll}})
			if err != nil {
				return nil, nil, err
			}
			svc.Kind = KubernetesKindService
			svc.APIVersion = KubernetesAPIVersionService

			logrus.WithContext(ctx).WithFields(logrus.Fields{
				"old service": s,
				"service":     svc,
			}).Debug("update service")
			if !reflect.DeepEqual(svc.Spec, s.Spec) {
				out = append(out, ResourceChange{
					ChangeType:      ChangeTypeUpdate,
					CurrentResource: s,
					DesiredResource: svc,
				})
				changeType = ChangeTypeUpdate
			} else {
				logrus.WithContext(ctx).WithFields(logrus.Fields{
					"service":          svc,
					"existing_service": s,
				}).Debug("Service is up to date")
			}
		}

		// Check if the service needs an Ingress
		ct, ingressChanges, err := p.GetServiceIngressChanges(ctx, name)
		if err != nil {
			logrus.WithContext(ctx).WithFields(logrus.Fields{
				"error": err,
			}).Error("Error getting Ingress changes")
			return nil, nil, err
		}
		logrus.WithContext(ctx).WithFields(logrus.Fields{
			"ingress_changes": ingressChanges,
			"changeType":      changeType,
			"ingress_ct":      ct,
		}).Trace("Ingress changes")

		changeType = *ct
		out = append(out, ingressChanges...)
	}

	logrus.WithContext(ctx).WithFields(logrus.Fields{
		"changeType":       changeType,
		"resource_changes": out,
	}).Debug("GetServiceResourceChanges")

	return &changeType, out, nil
}

func (pod *Pod) GetServiceIngressChanges(ctx context.Context, serviceName string) (*ChangeType, []ResourceChange, error) {
	var out []ResourceChange
	changeType := ChangeTypeNoop

	// Check if the service has a URL and protocol set
	if pod.Services[serviceName].Protocol == nil || // If the protocol is nil, it's TCP
		pod.Services[serviceName].URL == nil || // If the URL is nil, no ingress is needed
		(strings.ToLower(*pod.Services[serviceName].Protocol) != "http" && strings.ToLower(*pod.Services[serviceName].Protocol) != "https") { // If the protocol is not HTTP or HTTPS, no ingress is needed
		// No ingress needed
		return nil, nil, nil
	}

	u, err := url.ParseRequestURI(*pod.Services[serviceName].URL)
	if err != nil {
		logrus.WithContext(ctx).WithFields(logrus.Fields{
			"error": err,
			"url":   *pod.Services[serviceName].URL,
		}).Error("Error parsing URL")
		return nil, nil, err
	}

	// Allow user to specify "http://example.com" for "http://example.com/"
	path := u.Path
	if path == "" {
		path = "/"
	}

	// Get the Kube connection and try to Get an existing ingress
	client, err := kube_client.Client()
	if err != nil {
		return nil, nil, err
	}

	// For some reason the kube client will not automatically apply the default ingressclass
	// to objects that are updated, so we doing that manually
	ic, err := getDefaultIngressClass(ctx, client)
	if err != nil {
		logrus.WithContext(ctx).WithFields(logrus.Fields{
			"error": err,
		}).Error("Error getting default IngressClass")
		return nil, nil, err
	}

	ingressSpec := &networkingv1.Ingress{
		ObjectMeta: metav1.ObjectMeta{
			Name:      serviceName,
			Namespace: *pod.Namespace,
			Labels: map[string]string{
				"podinate.com/pod": pod.ID,
			},
		},
		Spec: networkingv1.IngressSpec{
			IngressClassName: ic,
			Rules: []networkingv1.IngressRule{
				{
					Host: u.Hostname(),
					IngressRuleValue: networkingv1.IngressRuleValue{
						HTTP: &networkingv1.HTTPIngressRuleValue{
							Paths: []networkingv1.HTTPIngressPath{
								{
									PathType: func(val networkingv1.PathType) *networkingv1.PathType { return &val }(networkingv1.PathTypePrefix),
									Path:     path,
									Backend: networkingv1.IngressBackend{
										Service: &networkingv1.IngressServiceBackend{
											Name: serviceName,
											Port: networkingv1.ServiceBackendPort{
												Number: int32(pod.Services[serviceName].Port),
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}
	ingressSpec.Kind = KubernetesKindIngress
	ingressSpec.APIVersion = KubernetesAPIVersionIngress

	existing, err := client.NetworkingV1().Ingresses(*pod.Namespace).Get(ctx, serviceName, metav1.GetOptions{})
	existing.Kind = KubernetesKindIngress
	existing.APIVersion = KubernetesAPIVersionIngress

	if errors.IsNotFound(err) {
		// Resource doesn't exist, so create it

		created, err := client.NetworkingV1().Ingresses(*pod.Namespace).Create(ctx, ingressSpec, metav1.CreateOptions{DryRun: []string{metav1.DryRunAll}})
		if errors.IsNotFound(err) {
			// We get this error when the Namespace isn't found, in that case it's okay because we can just apply the spec as is
			created = ingressSpec
		} else if err != nil {
			return nil, nil, err
		}
		created.Kind = KubernetesKindIngress
		created.APIVersion = KubernetesAPIVersionIngress

		changeType = ChangeTypeCreate
		out = append(out, ResourceChange{
			ChangeType:      ChangeTypeCreate,
			CurrentResource: nil,
			DesiredResource: created,
		})
	} else if err != nil {
		// Something else went wrong getting the ingress, we can't proceed
		logrus.WithContext(ctx).WithFields(logrus.Fields{
			"error":   err,
			"pod_id":  pod.ID,
			"service": serviceName,
		}).Error("Error getting Ingress")
		return nil, nil, err
	} else {
		// Resource exists, but needs updating
		updated, err := client.NetworkingV1().Ingresses(*pod.Namespace).Update(ctx, ingressSpec, metav1.UpdateOptions{DryRun: []string{metav1.DryRunAll}})
		if err != nil {
			return nil, nil, err
		}
		updated.Kind = KubernetesKindIngress
		updated.APIVersion = KubernetesAPIVersionIngress

		if !reflect.DeepEqual(updated.Spec, existing.Spec) {
			logrus.WithContext(ctx).WithFields(logrus.Fields{
				"existing": existing,
				"ingress":  updated,
			}).Debug("Ingress needs updating")
			out = append(out, ResourceChange{
				ChangeType:      ChangeTypeUpdate,
				CurrentResource: existing,
				DesiredResource: updated,
			})
			changeType = ChangeTypeUpdate
		}
	}

	return &changeType, out, nil

}

// getDefaultIngressClass gets the default IngressClass from the cluster
func getDefaultIngressClass(ctx context.Context, client *kubernetes.Clientset) (*string, error) {
	// Get the IngressClass from the cluster

	ingressClasses, err := client.NetworkingV1().IngressClasses().List(ctx, metav1.ListOptions{})
	if err != nil {
		return nil, err
	}

	logrus.WithContext(ctx).WithFields(logrus.Fields{
		"ingressClasses": ingressClasses,
	}).Debug("Got IngressClasses")

	// Find the default IngressClass
	for _, ic := range ingressClasses.Items {
		if ic.ObjectMeta.Annotations["ingressclass.kubernetes.io/is-default-class"] == "true" {
			return &ic.Name, nil
		}
	}

	return nil, stderrors.New("No default IngressClass found")
}
