/*
Copyright 2023 ImagineKube Authors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

     http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1alpha3

import (
	"fmt"

	"github.com/emicklei/go-restful"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/client-go/kubernetes"
	"k8s.io/klog/v2"

	"imaginekube.com/imaginekube/pkg/api"
	"imaginekube.com/imaginekube/pkg/apiserver/authorization/authorizer"
	"imaginekube.com/imaginekube/pkg/apiserver/query"
	"imaginekube.com/imaginekube/pkg/apiserver/request"
	imaginekube "imaginekube.com/imaginekube/pkg/client/clientset/versioned"
	"imaginekube.com/imaginekube/pkg/informers"
	"imaginekube.com/imaginekube/pkg/models/iam/am"
	"imaginekube.com/imaginekube/pkg/models/iam/im"
	"imaginekube.com/imaginekube/pkg/models/openpitrix"
	resourcev1alpha3 "imaginekube.com/imaginekube/pkg/models/resources/v1alpha3/resource"
	"imaginekube.com/imaginekube/pkg/models/tenant"
	"imaginekube.com/imaginekube/pkg/simple/client/auditing"
	"imaginekube.com/imaginekube/pkg/simple/client/events"
	"imaginekube.com/imaginekube/pkg/simple/client/logging"
	meteringclient "imaginekube.com/imaginekube/pkg/simple/client/metering"
	monitoringclient "imaginekube.com/imaginekube/pkg/simple/client/monitoring"
)

type tenantHandler struct {
	tenant          tenant.Interface
	meteringOptions *meteringclient.Options
}

func newTenantHandler(factory informers.InformerFactory, k8sclient kubernetes.Interface, ksclient imaginekube.Interface,
	evtsClient events.Client, loggingClient logging.Client, auditingclient auditing.Client,
	am am.AccessManagementInterface, im im.IdentityManagementInterface, authorizer authorizer.Authorizer,
	monitoringclient monitoringclient.Interface, resourceGetter *resourcev1alpha3.ResourceGetter,
	meteringOptions *meteringclient.Options, opClient openpitrix.Interface) *tenantHandler {

	if meteringOptions == nil || meteringOptions.RetentionDay == "" {
		meteringOptions = &meteringclient.DefaultMeteringOption
	}

	return &tenantHandler{
		tenant:          tenant.New(factory, k8sclient, ksclient, evtsClient, loggingClient, auditingclient, am, im, authorizer, monitoringclient, resourceGetter, opClient),
		meteringOptions: meteringOptions,
	}
}

func (h *tenantHandler) ListWorkspaces(req *restful.Request, resp *restful.Response) {
	queryParam := query.ParseQueryParameter(req)
	user, ok := request.UserFrom(req.Request.Context())
	if !ok {
		err := fmt.Errorf("cannot obtain user info")
		klog.Errorln(err)
		api.HandleForbidden(resp, nil, err)
		return
	}

	result, err := h.tenant.ListWorkspaces(user, queryParam)
	if err != nil {
		api.HandleInternalError(resp, nil, err)
		return
	}

	resp.WriteEntity(result)
}

func (h *tenantHandler) GetWorkspace(request *restful.Request, response *restful.Response) {
	workspace, err := h.tenant.GetWorkspace(request.PathParameter("workspace"))
	if err != nil {
		klog.Error(err)
		if errors.IsNotFound(err) {
			api.HandleNotFound(response, request, err)
			return
		}
		api.HandleInternalError(response, request, err)
		return
	}

	response.WriteEntity(workspace)
}
