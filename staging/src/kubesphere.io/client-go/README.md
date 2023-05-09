# ImagineKube client-go Project

The ImagineKube client-go Project is a rest-client of go libraries for communicating with the ImagineKube API Server.

# How to use it

1. Import client-go packages:
```golang
import (
    "k8s.io/client-go/rest"
	"imaginekube.com/client-go/client"
	"imaginekube.com/client-go/client/generic"
)
```
2. Create a generic client instance:
```golang
    var client client.Client
	config := &rest.Config{
		Host:     "127.0.0.1:9090",
		Username: "admin",
		Password: "P@88w0rd",
	}
	client = generic.NewForConfigOrDie(config, client.Options{Scheme: f.Scheme})
```
> generic.NewForConfigOrDie returns a client.Client that reads and writes from/to an ImagineKube API server. 

> It's only compatible with Kubernetes-like API objects.

3. ImagineKube API server provided a proxy to Kubernetes API Server. The client can read and write those Kubernetes native objects with the client directly.

```golang
	deploy := &appsv1.Deployment{}
	client.Get(context.TODO(), client.ObjectKey{Namespace: "imaginekube-system", Name: "ks-apiserver"}, deploy)
```

4. URLOptions and WorkspaceOptions can be provided to read and write Kubernetes likely Object that provided by ImagineKube API.
```golang
	ns := &corev1.Namespace{
		ObjectMeta: metav1.ObjectMeta{
			Name: "ks-test",
			Labels: map[string]string{
				constants.WorkspaceLabelKey: "Workspace",
			},
		},
	}

	opts := &client.URLOptions{
		Group:   "tenant.imaginekube.com",
		Version: "v1alpha2",
	}

	err := f.GenericClient(f.BaseName).Create(context.TODO(), ns, opts, &client.WorkspaceOptions{Name: "Workspace"})
```

The ImagineKube API Architecture can be found at https://imaginekube.com/docs/reference/api-docs/

# Where does it come from?

client-go is synced from https://github.com/imaginekube/imaginekube/blob/master/staging/src/imaginekube.com/client-go. Code changes are made in that location, merged into `imaginekube.com/client-go` and later synced here.