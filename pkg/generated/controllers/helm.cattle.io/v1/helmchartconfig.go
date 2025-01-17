/*
Copyright The Kubernetes Authors.

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

// Code generated by main. DO NOT EDIT.

package v1

import (
	"context"
	"time"

	v1 "github.com/TinyStac/helm-controller/pkg/apis/helm.cattle.io/v1"
	"github.com/rancher/lasso/pkg/client"
	"github.com/rancher/lasso/pkg/controller"
	"github.com/rancher/wrangler/pkg/generic"
	"k8s.io/apimachinery/pkg/api/equality"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/tools/cache"
)

type HelmChartConfigHandler func(string, *v1.HelmChartConfig) (*v1.HelmChartConfig, error)

type HelmChartConfigController interface {
	generic.ControllerMeta
	HelmChartConfigClient

	OnChange(ctx context.Context, name string, sync HelmChartConfigHandler)
	OnRemove(ctx context.Context, name string, sync HelmChartConfigHandler)
	Enqueue(namespace, name string)
	EnqueueAfter(namespace, name string, duration time.Duration)

	Cache() HelmChartConfigCache
}

type HelmChartConfigClient interface {
	Create(*v1.HelmChartConfig) (*v1.HelmChartConfig, error)
	Update(*v1.HelmChartConfig) (*v1.HelmChartConfig, error)

	Delete(namespace, name string, options *metav1.DeleteOptions) error
	Get(namespace, name string, options metav1.GetOptions) (*v1.HelmChartConfig, error)
	List(namespace string, opts metav1.ListOptions) (*v1.HelmChartConfigList, error)
	Watch(namespace string, opts metav1.ListOptions) (watch.Interface, error)
	Patch(namespace, name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.HelmChartConfig, err error)
}

type HelmChartConfigCache interface {
	Get(namespace, name string) (*v1.HelmChartConfig, error)
	List(namespace string, selector labels.Selector) ([]*v1.HelmChartConfig, error)

	AddIndexer(indexName string, indexer HelmChartConfigIndexer)
	GetByIndex(indexName, key string) ([]*v1.HelmChartConfig, error)
}

type HelmChartConfigIndexer func(obj *v1.HelmChartConfig) ([]string, error)

type helmChartConfigController struct {
	controller    controller.SharedController
	client        *client.Client
	gvk           schema.GroupVersionKind
	groupResource schema.GroupResource
}

func NewHelmChartConfigController(gvk schema.GroupVersionKind, resource string, namespaced bool, controller controller.SharedControllerFactory) HelmChartConfigController {
	c := controller.ForResourceKind(gvk.GroupVersion().WithResource(resource), gvk.Kind, namespaced)
	return &helmChartConfigController{
		controller: c,
		client:     c.Client(),
		gvk:        gvk,
		groupResource: schema.GroupResource{
			Group:    gvk.Group,
			Resource: resource,
		},
	}
}

func FromHelmChartConfigHandlerToHandler(sync HelmChartConfigHandler) generic.Handler {
	return func(key string, obj runtime.Object) (ret runtime.Object, err error) {
		var v *v1.HelmChartConfig
		if obj == nil {
			v, err = sync(key, nil)
		} else {
			v, err = sync(key, obj.(*v1.HelmChartConfig))
		}
		if v == nil {
			return nil, err
		}
		return v, err
	}
}

func (c *helmChartConfigController) Updater() generic.Updater {
	return func(obj runtime.Object) (runtime.Object, error) {
		newObj, err := c.Update(obj.(*v1.HelmChartConfig))
		if newObj == nil {
			return nil, err
		}
		return newObj, err
	}
}

func UpdateHelmChartConfigDeepCopyOnChange(client HelmChartConfigClient, obj *v1.HelmChartConfig, handler func(obj *v1.HelmChartConfig) (*v1.HelmChartConfig, error)) (*v1.HelmChartConfig, error) {
	if obj == nil {
		return obj, nil
	}

	copyObj := obj.DeepCopy()
	newObj, err := handler(copyObj)
	if newObj != nil {
		copyObj = newObj
	}
	if obj.ResourceVersion == copyObj.ResourceVersion && !equality.Semantic.DeepEqual(obj, copyObj) {
		return client.Update(copyObj)
	}

	return copyObj, err
}

func (c *helmChartConfigController) AddGenericHandler(ctx context.Context, name string, handler generic.Handler) {
	c.controller.RegisterHandler(ctx, name, controller.SharedControllerHandlerFunc(handler))
}

func (c *helmChartConfigController) AddGenericRemoveHandler(ctx context.Context, name string, handler generic.Handler) {
	c.AddGenericHandler(ctx, name, generic.NewRemoveHandler(name, c.Updater(), handler))
}

func (c *helmChartConfigController) OnChange(ctx context.Context, name string, sync HelmChartConfigHandler) {
	c.AddGenericHandler(ctx, name, FromHelmChartConfigHandlerToHandler(sync))
}

func (c *helmChartConfigController) OnRemove(ctx context.Context, name string, sync HelmChartConfigHandler) {
	c.AddGenericHandler(ctx, name, generic.NewRemoveHandler(name, c.Updater(), FromHelmChartConfigHandlerToHandler(sync)))
}

func (c *helmChartConfigController) Enqueue(namespace, name string) {
	c.controller.Enqueue(namespace, name)
}

func (c *helmChartConfigController) EnqueueAfter(namespace, name string, duration time.Duration) {
	c.controller.EnqueueAfter(namespace, name, duration)
}

func (c *helmChartConfigController) Informer() cache.SharedIndexInformer {
	return c.controller.Informer()
}

func (c *helmChartConfigController) GroupVersionKind() schema.GroupVersionKind {
	return c.gvk
}

func (c *helmChartConfigController) Cache() HelmChartConfigCache {
	return &helmChartConfigCache{
		indexer:  c.Informer().GetIndexer(),
		resource: c.groupResource,
	}
}

func (c *helmChartConfigController) Create(obj *v1.HelmChartConfig) (*v1.HelmChartConfig, error) {
	result := &v1.HelmChartConfig{}
	return result, c.client.Create(context.TODO(), obj.Namespace, obj, result, metav1.CreateOptions{})
}

func (c *helmChartConfigController) Update(obj *v1.HelmChartConfig) (*v1.HelmChartConfig, error) {
	result := &v1.HelmChartConfig{}
	return result, c.client.Update(context.TODO(), obj.Namespace, obj, result, metav1.UpdateOptions{})
}

func (c *helmChartConfigController) Delete(namespace, name string, options *metav1.DeleteOptions) error {
	if options == nil {
		options = &metav1.DeleteOptions{}
	}
	return c.client.Delete(context.TODO(), namespace, name, *options)
}

func (c *helmChartConfigController) Get(namespace, name string, options metav1.GetOptions) (*v1.HelmChartConfig, error) {
	result := &v1.HelmChartConfig{}
	return result, c.client.Get(context.TODO(), namespace, name, result, options)
}

func (c *helmChartConfigController) List(namespace string, opts metav1.ListOptions) (*v1.HelmChartConfigList, error) {
	result := &v1.HelmChartConfigList{}
	return result, c.client.List(context.TODO(), namespace, result, opts)
}

func (c *helmChartConfigController) Watch(namespace string, opts metav1.ListOptions) (watch.Interface, error) {
	return c.client.Watch(context.TODO(), namespace, opts)
}

func (c *helmChartConfigController) Patch(namespace, name string, pt types.PatchType, data []byte, subresources ...string) (*v1.HelmChartConfig, error) {
	result := &v1.HelmChartConfig{}
	return result, c.client.Patch(context.TODO(), namespace, name, pt, data, result, metav1.PatchOptions{}, subresources...)
}

type helmChartConfigCache struct {
	indexer  cache.Indexer
	resource schema.GroupResource
}

func (c *helmChartConfigCache) Get(namespace, name string) (*v1.HelmChartConfig, error) {
	obj, exists, err := c.indexer.GetByKey(namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(c.resource, name)
	}
	return obj.(*v1.HelmChartConfig), nil
}

func (c *helmChartConfigCache) List(namespace string, selector labels.Selector) (ret []*v1.HelmChartConfig, err error) {

	err = cache.ListAllByNamespace(c.indexer, namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.HelmChartConfig))
	})

	return ret, err
}

func (c *helmChartConfigCache) AddIndexer(indexName string, indexer HelmChartConfigIndexer) {
	utilruntime.Must(c.indexer.AddIndexers(map[string]cache.IndexFunc{
		indexName: func(obj interface{}) (strings []string, e error) {
			return indexer(obj.(*v1.HelmChartConfig))
		},
	}))
}

func (c *helmChartConfigCache) GetByIndex(indexName, key string) (result []*v1.HelmChartConfig, err error) {
	objs, err := c.indexer.ByIndex(indexName, key)
	if err != nil {
		return nil, err
	}
	result = make([]*v1.HelmChartConfig, 0, len(objs))
	for _, obj := range objs {
		result = append(result, obj.(*v1.HelmChartConfig))
	}
	return result, nil
}
