/*
Copyright The Longhorn Authors.

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

// Code generated by lister-gen. DO NOT EDIT.

package v1beta1

import (
	longhornv1beta1 "github.com/longhorn/longhorn-manager/k8s/pkg/apis/longhorn/v1beta1"
	labels "k8s.io/apimachinery/pkg/labels"
	listers "k8s.io/client-go/listers"
	cache "k8s.io/client-go/tools/cache"
)

// NodeLister helps list Nodes.
// All objects returned here must be treated as read-only.
type NodeLister interface {
	// List lists all Nodes in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*longhornv1beta1.Node, err error)
	// Nodes returns an object that can list and get Nodes.
	Nodes(namespace string) NodeNamespaceLister
	NodeListerExpansion
}

// nodeLister implements the NodeLister interface.
type nodeLister struct {
	listers.ResourceIndexer[*longhornv1beta1.Node]
}

// NewNodeLister returns a new NodeLister.
func NewNodeLister(indexer cache.Indexer) NodeLister {
	return &nodeLister{listers.New[*longhornv1beta1.Node](indexer, longhornv1beta1.Resource("node"))}
}

// Nodes returns an object that can list and get Nodes.
func (s *nodeLister) Nodes(namespace string) NodeNamespaceLister {
	return nodeNamespaceLister{listers.NewNamespaced[*longhornv1beta1.Node](s.ResourceIndexer, namespace)}
}

// NodeNamespaceLister helps list and get Nodes.
// All objects returned here must be treated as read-only.
type NodeNamespaceLister interface {
	// List lists all Nodes in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*longhornv1beta1.Node, err error)
	// Get retrieves the Node from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*longhornv1beta1.Node, error)
	NodeNamespaceListerExpansion
}

// nodeNamespaceLister implements the NodeNamespaceLister
// interface.
type nodeNamespaceLister struct {
	listers.ResourceIndexer[*longhornv1beta1.Node]
}
