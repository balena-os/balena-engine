package daemon // import "github.com/docker/docker/daemon"

import (
	"context"
	"strings"
	"time"

	"github.com/docker/docker/api/types/events"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/container"
	daemonevents "github.com/docker/docker/daemon/events"
	"github.com/docker/docker/libnetwork"
)

// LogContainerEvent generates an event related to a container with only the default attributes.
func (daemon *Daemon) LogContainerEvent(container *container.Container, action events.Action) {
	daemon.LogContainerEventWithAttributes(container, action, map[string]string{})
}

// LogContainerEventWithAttributes generates an event related to a container with specific given attributes.
func (daemon *Daemon) LogContainerEventWithAttributes(container *container.Container, action events.Action, attributes map[string]string) {
	copyAttributes(attributes, container.Config.Labels)
	if container.Config.Image != "" {
		attributes["image"] = container.Config.Image
	}
	attributes["name"] = strings.TrimLeft(container.Name, "/")
	daemon.EventsService.Log(action, events.ContainerEventType, events.Actor{
		ID:         container.ID,
		Attributes: attributes,
	})
}

// LogPluginEvent generates an event related to a plugin with only the default attributes.
func (daemon *Daemon) LogPluginEvent(pluginID, refName string, action events.Action) {
	daemon.EventsService.Log(action, events.PluginEventType, events.Actor{
		ID:         pluginID,
		Attributes: map[string]string{"name": refName},
	})
}

// LogVolumeEvent generates an event related to a volume.
func (daemon *Daemon) LogVolumeEvent(volumeID string, action events.Action, attributes map[string]string) {
	daemon.EventsService.Log(action, events.VolumeEventType, events.Actor{
		ID:         volumeID,
		Attributes: attributes,
	})
}

// LogNetworkEvent generates an event related to a network with only the default attributes.
func (daemon *Daemon) LogNetworkEvent(nw *libnetwork.Network, action events.Action) {
	daemon.LogNetworkEventWithAttributes(nw, action, map[string]string{})
}

// LogNetworkEventWithAttributes generates an event related to a network with specific given attributes.
func (daemon *Daemon) LogNetworkEventWithAttributes(nw *libnetwork.Network, action events.Action, attributes map[string]string) {
	attributes["name"] = nw.Name()
	attributes["type"] = nw.Type()
	daemon.EventsService.Log(action, events.NetworkEventType, events.Actor{
		ID:         nw.ID(),
		Attributes: attributes,
	})
}

// LogDaemonEventWithAttributes generates an event related to the daemon itself with specific given attributes.
func (daemon *Daemon) LogDaemonEventWithAttributes(action events.Action, attributes map[string]string) {
	if daemon.EventsService != nil {
		if name := hostName(context.TODO()); name != "" {
			attributes["name"] = name
		}
		daemon.EventsService.Log(action, events.DaemonEventType, events.Actor{
			ID:         daemon.id,
			Attributes: attributes,
		})
	}
}

// SubscribeToEvents returns the currently record of events, a channel to stream new events from, and a function to cancel the stream of events.
func (daemon *Daemon) SubscribeToEvents(since, until time.Time, filter filters.Args) ([]events.Message, chan interface{}) {
	return daemon.EventsService.SubscribeTopic(since, until, daemonevents.NewFilter(filter))
}

// UnsubscribeFromEvents stops the event subscription for a client by closing the
// channel where the daemon sends events to.
func (daemon *Daemon) UnsubscribeFromEvents(listener chan interface{}) {
	daemon.EventsService.Evict(listener)
}

// copyAttributes guarantees that labels are not mutated by event triggers.
func copyAttributes(attributes, labels map[string]string) {
	if labels == nil {
		return
	}
	for k, v := range labels {
		attributes[k] = v
	}
}
