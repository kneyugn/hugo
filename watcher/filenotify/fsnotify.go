// Package filenotify is adapted from https://github.com/moby/moby/tree/master/pkg/filenotify, Apache-2.0 License.
// Hopefully this can be replaced with an external package sometime in the future, see https://github.com/fsnotify/fsnotify/issues/9
package filenotify

import "github.com/fsnotify/fsnotify"

// fsNotifyWatcher wraps the fsnotify package to satisfy the FileNotifier interface
type fsNotifyWatcher struct {
	*fsnotify.Watcher
}

// Events returns the fsnotify event channel receiver
func (w *fsNotifyWatcher) Events() <-chan fsnotify.Event {
	return nil
}

// Errors returns the fsnotify error channel receiver
func (w *fsNotifyWatcher) Errors() <-chan error {
	return nil
}
