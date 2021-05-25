// Copyright 2019 The Kubernetes Authors.
// SPDX-License-Identifier: Apache-2.0

package filters

import (
	"fmt"

	"sigs.k8s.io/kustomize/kyaml/kio"
	"sigs.k8s.io/kustomize/kyaml/kio/kioutil"
	"sigs.k8s.io/kustomize/kyaml/yaml"
	"sigs.k8s.io/kustomize/kyaml/yaml/merge3"
)

const (
	mergeSourceAnnotation = "config.kubernetes.io/merge-source"
	mergeSourceOriginal   = "original"
	mergeSourceUpdated    = "updated"
	mergeSourceDest       = "dest"
)

// ResourceMatcher interface is used to match two resources based on IsSameResource implementation
// This is the way to group same logical resources in upstream, local and origin for merge
// The default way to group them is using GVKNN similar to how kubernetes server identifies resources
// Users of this library might have their own interpretation of grouping similar resources
// for e.g. if consumer adds a name-prefix to local resource, it should not be treated as new resource
// for updates etc.
// Hence, the callers of this library may pass different implementation for IsSameResource
type ResourceMatcher interface {
	IsSameResource(node1, node2 *yaml.RNode) bool
}

// Merge3 performs a 3-way merge on the original, updated, and destination packages.
type Merge3 struct {
	OriginalPath   string
	UpdatedPath    string
	DestPath       string
	MatchFilesGlob []string
	Matcher        ResourceMatcher
}

func (m Merge3) Merge() error {
	// Read the destination package.  The ReadWriter will take take of deleting files
	// for removed resources.
	var inputs []kio.Reader
	dest := &kio.LocalPackageReadWriter{
		PackagePath:    m.DestPath,
		MatchFilesGlob: m.MatchFilesGlob,
		SetAnnotations: map[string]string{mergeSourceAnnotation: mergeSourceDest},
	}
	inputs = append(inputs, dest)

	// Read the original package
	inputs = append(inputs, kio.LocalPackageReader{
		PackagePath:    m.OriginalPath,
		MatchFilesGlob: m.MatchFilesGlob,
		SetAnnotations: map[string]string{mergeSourceAnnotation: mergeSourceOriginal},
	})

	// Read the updated package
	inputs = append(inputs, kio.LocalPackageReader{
		PackagePath:    m.UpdatedPath,
		MatchFilesGlob: m.MatchFilesGlob,
		SetAnnotations: map[string]string{mergeSourceAnnotation: mergeSourceUpdated},
	})

	return kio.Pipeline{
		Inputs:  inputs,
		Filters: []kio.Filter{m},
		Outputs: []kio.Writer{dest},
	}.Execute()
}

// Filter combines Resources with the same GVK + N + NS into tuples, and then merges them
func (m Merge3) Filter(nodes []*yaml.RNode) ([]*yaml.RNode, error) {
	// index the nodes by their identity
	matcher := m.Matcher
	if matcher == nil {
		matcher = &DefaultGVKNNMatcher{MergeOnPath: true}
	}
	tl := tuples{matcher: matcher}
	for i := range nodes {
		if err := tl.add(nodes[i]); err != nil {
			return nil, err
		}
	}

	// iterate over the inputs, merging as needed
	var output []*yaml.RNode
	for i := range tl.list {
		t := tl.list[i]
		switch {
		case t.original == nil && t.updated == nil && t.dest != nil:
			// added locally -- keep dest
			output = append(output, t.dest)
		case t.updated != nil && t.dest == nil:
			// added in the update -- add update
			output = append(output, t.updated)
		case t.original != nil && t.updated == nil:
			// deleted in the update
		// don't include the resource in the output
		case t.original != nil && t.dest == nil:
			// deleted locally
			// don't include the resource in the output
		default:
			// dest and updated are non-nil -- merge them
			node, err := t.merge()
			if err != nil {
				return nil, err
			}
			if node != nil {
				output = append(output, node)
			}
		}
	}
	return output, nil
}

// tuples combines nodes with the same GVK + N + NS
type tuples struct {
	list []*tuple

	// matcher matches the resources for merge
	matcher ResourceMatcher
}

// DefaultGVKNNMatcher holds the default matching of resources implementation based on
// Group, Version, Kind, Name and Namespace of the resource
type DefaultGVKNNMatcher struct {
	// MergeOnPath will use the relative filepath as part of the merge key.
	// This may be necessary if the directory contains multiple copies of
	// the same resource, or resources patches.
	MergeOnPath bool
}

// IsSameResource returns true if metadata of node1 and metadata of node2 belongs to same logical resource
func (dm *DefaultGVKNNMatcher) IsSameResource(node1, node2 *yaml.RNode) bool {
	if node1 == nil || node2 == nil {
		return false
	}

	meta1, err := node1.GetMeta()
	if err != nil {
		return false
	}

	meta2, err := node2.GetMeta()
	if err != nil {
		return false
	}

	if meta1.Name != meta2.Name {
		return false
	}
	if meta1.Namespace != meta2.Namespace {
		return false
	}
	if meta1.APIVersion != meta2.APIVersion {
		return false
	}
	if meta1.Kind != meta2.Kind {
		return false
	}
	if dm.MergeOnPath {
		// directories may contain multiple copies of a resource with the same
		// name, namespace, apiVersion and kind -- e.g. kustomize patches, or
		// multiple environments
		// mergeOnPath configures the merge logic to use the path as part of the
		// resource key
		if meta1.Annotations[kioutil.PathAnnotation] != meta2.Annotations[kioutil.PathAnnotation] {
			return false
		}
	}
	return true
}

// add adds a node to the list, combining it with an existing matching Resource if found
func (ts *tuples) add(node *yaml.RNode) error {
	for i := range ts.list {
		t := ts.list[i]
		if ts.matcher.IsSameResource(addedNode(t), node) {
			return t.add(node)
		}
	}
	t := &tuple{}
	if err := t.add(node); err != nil {
		return err
	}
	ts.list = append(ts.list, t)
	return nil
}

// addedNode returns one on the existing added nodes in the tuple
func addedNode(t *tuple) *yaml.RNode {
	if t.updated != nil {
		return t.updated
	}
	if t.original != nil {
		return t.original
	}
	return t.dest
}

// tuple wraps an original, updated, and dest tuple for a given Resource
type tuple struct {
	original *yaml.RNode
	updated  *yaml.RNode
	dest     *yaml.RNode
}

// add sets the corresponding tuple field for the node
func (t *tuple) add(node *yaml.RNode) error {
	meta, err := node.GetMeta()
	if err != nil {
		return err
	}
	switch meta.Annotations[mergeSourceAnnotation] {
	case mergeSourceDest:
		if t.dest != nil {
			return duplicateError("local", meta.Annotations[kioutil.PathAnnotation])
		}
		t.dest = node
	case mergeSourceOriginal:
		if t.original != nil {
			return duplicateError("original upstream", meta.Annotations[kioutil.PathAnnotation])
		}
		t.original = node
	case mergeSourceUpdated:
		if t.updated != nil {
			return duplicateError("updated upstream", meta.Annotations[kioutil.PathAnnotation])
		}
		t.updated = node
	default:
		return fmt.Errorf("no source annotation for Resource")
	}
	return nil
}

// merge performs a 3-way merge on the tuple
func (t *tuple) merge() (*yaml.RNode, error) {
	return merge3.Merge(t.dest, t.original, t.updated)
}

// duplicateError returns duplicate resources error
func duplicateError(source, filePath string) error {
	return fmt.Errorf(`found duplicate %q resources in file %q, please refer to "update" documentation for the fix`, source, filePath)
}
