/*
Copyright 2023 The ImagineKube Authors.

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

package openpitrix

import (
	"k8s.io/client-go/tools/cache"
	"k8s.io/klog/v2"

	"imaginekube.com/imaginekube/pkg/utils/clusterclient"

	"imaginekube.com/api/application/v1alpha1"

	"imaginekube.com/imaginekube/pkg/client/clientset/versioned"
	ks_informers "imaginekube.com/imaginekube/pkg/informers"
	"imaginekube.com/imaginekube/pkg/simple/client/s3"
	"imaginekube.com/imaginekube/pkg/utils/reposcache"
)

type Interface interface {
	AttachmentInterface
	ApplicationInterface
	RepoInterface
	ReleaseInterface
	CategoryInterface
}

type openpitrixOperator struct {
	AttachmentInterface
	ApplicationInterface
	RepoInterface
	ReleaseInterface
	CategoryInterface
}

func NewOpenpitrixOperator(ksInformers ks_informers.InformerFactory, ksClient versioned.Interface, s3Client s3.Interface, cc clusterclient.ClusterClients) Interface {
	klog.Infof("start helm repo informer")
	cachedReposData := reposcache.NewReposCache()
	helmReposInformer := ksInformers.ImagineKubeSharedInformerFactory().Application().V1alpha1().HelmRepos().Informer()
	helmReposInformer.AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc: func(obj interface{}) {
			r := obj.(*v1alpha1.HelmRepo)
			cachedReposData.AddRepo(r)
		},
		UpdateFunc: func(oldObj, newObj interface{}) {
			oldRepo := oldObj.(*v1alpha1.HelmRepo)
			newRepo := newObj.(*v1alpha1.HelmRepo)
			cachedReposData.UpdateRepo(oldRepo, newRepo)
		},
		DeleteFunc: func(obj interface{}) {
			r := obj.(*v1alpha1.HelmRepo)
			cachedReposData.DeleteRepo(r)
		},
	})

	ctgInformer := ksInformers.ImagineKubeSharedInformerFactory().Application().V1alpha1().HelmCategories().Informer()
	ctgInformer.AddIndexers(map[string]cache.IndexFunc{
		reposcache.CategoryIndexer: func(obj interface{}) ([]string, error) {
			ctg, _ := obj.(*v1alpha1.HelmCategory)
			return []string{ctg.Spec.Name}, nil
		},
	})
	indexer := ctgInformer.GetIndexer()

	cachedReposData.SetCategoryIndexer(indexer)

	return &openpitrixOperator{
		AttachmentInterface:  newAttachmentOperator(s3Client),
		ApplicationInterface: newApplicationOperator(cachedReposData, ksInformers.ImagineKubeSharedInformerFactory(), ksClient, s3Client),
		RepoInterface:        newRepoOperator(cachedReposData, ksInformers.ImagineKubeSharedInformerFactory(), ksClient),
		ReleaseInterface:     newReleaseOperator(cachedReposData, ksInformers.KubernetesSharedInformerFactory(), ksInformers.ImagineKubeSharedInformerFactory(), ksClient, cc),
		CategoryInterface:    newCategoryOperator(cachedReposData, ksInformers.ImagineKubeSharedInformerFactory(), ksClient),
	}
}
