// Copyright 2019 Istio Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package event

import (
	"fmt"

	"istio.io/istio/galley/pkg/config/collection"
	"istio.io/istio/galley/pkg/config/resource"
)

// Event represents a change that occurred against a resource in the source config system.
type Event struct {
	Kind Kind

	// The collection that this event is emanating from.
	Source collection.Name

	// A single entry, in case the event is Added, Updated or Deleted.
	Entry *resource.Entry
}

// Clone creates a deep clone of the event.
func (e Event) Clone() Event {
	entry := e.Entry
	if entry != nil {
		entry = entry.Clone()
	}
	return Event{
		Entry:  entry,
		Source: e.Source,
		Kind:   e.Kind,
	}
}

// String implements Stringer.String
func (e Event) String() string {
	switch e.Kind {
	case Added, Updated, Deleted:
		return fmt.Sprintf("[Event](%s: %v/%v)", e.Kind.String(), e.Source, e.Entry.Metadata.Name)
	case FullSync:
		return fmt.Sprintf("[Event](%s: %v)", e.Kind.String(), e.Source)
	default:
		return fmt.Sprintf("[Event](%s)", e.Kind.String())
	}
}

// FullSyncFor creates a FullSync event for the given source.
func FullSyncFor(source collection.Name) Event {
	return Event{Kind: FullSync, Source: source}
}