/*
Copyright 2023 The ImagineKube authors.

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

package v1alpha2

import (
	"net/http"

	restful "github.com/emicklei/go-restful"
	restfulspec "github.com/emicklei/go-restful-openapi"
	"k8s.io/apimachinery/pkg/runtime/schema"

	"imaginekube.com/imaginekube/pkg/apiserver/runtime"
	"imaginekube.com/imaginekube/pkg/constants"
)

const GroupName = "network.imaginekube.com"

var GroupVersion = schema.GroupVersion{Group: GroupName, Version: "v1alpha2"}

func AddToContainer(c *restful.Container, weaveScopeHost string) error {
	webservice := runtime.NewWebService(GroupVersion)
	h := handler{weaveScopeHost: weaveScopeHost}

	webservice.Route(webservice.GET("/namespaces/{namespace}/topology").
		To(h.getNamespaceTopology).
		Metadata(restfulspec.KeyOpenAPITags, []string{constants.NetworkTopologyTag}).
		Doc("Get the topology with specifying a namespace").
		Param(webservice.PathParameter("namespace", "name of the namespace").Required(true)).
		Returns(http.StatusOK, "ok", TopologyResponse{}).
		Writes(TopologyResponse{})).
		Produces(restful.MIME_JSON)

	webservice.Route(webservice.GET("/namespaces/{namespace}/topology/{node_id}").
		To(h.getNamespaceNodeTopology).
		Metadata(restfulspec.KeyOpenAPITags, []string{constants.NetworkTopologyTag}).
		Doc("Get the topology with specifying a node id in the whole topology and specifying a namespace").
		Param(webservice.PathParameter("namespace", "name of the namespace").Required(true)).
		Param(webservice.PathParameter("node_id", "id of the node in the whole topology").Required(true)).
		Returns(http.StatusOK, "ok", NodeResponse{}).
		Writes(NodeResponse{})).
		Produces(restful.MIME_JSON)

	c.Add(webservice)

	return nil
}
