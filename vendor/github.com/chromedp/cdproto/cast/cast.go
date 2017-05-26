// Package cast provides the Chrome DevTools Protocol
// commands, types, and events for the Cast domain.
//
// A domain for interacting with Cast, Presentation API, and Remote Playback
// API functionalities.
//
// Generated by the cdproto-gen command.
package cast

// Code generated by cdproto-gen. DO NOT EDIT.

import (
	"context"

	"github.com/chromedp/cdproto/cdp"
)

// EnableParams starts observing for sinks that can be used for tab
// mirroring, and if set, sinks compatible with |presentationUrl| as well. When
// sinks are found, a |sinksUpdated| event is fired. Also starts observing for
// issue messages. When an issue is added or removed, an |issueUpdated| event is
// fired.
type EnableParams struct {
	PresentationURL string `json:"presentationUrl,omitempty"`
}

// Enable starts observing for sinks that can be used for tab mirroring, and
// if set, sinks compatible with |presentationUrl| as well. When sinks are
// found, a |sinksUpdated| event is fired. Also starts observing for issue
// messages. When an issue is added or removed, an |issueUpdated| event is
// fired.
//
// parameters:
func Enable() *EnableParams {
	return &EnableParams{}
}

// WithPresentationURL [no description].
func (p EnableParams) WithPresentationURL(presentationURL string) *EnableParams {
	p.PresentationURL = presentationURL
	return &p
}

// Do executes Cast.enable against the provided context.
func (p *EnableParams) Do(ctxt context.Context, h cdp.Executor) (err error) {
	return h.Execute(ctxt, CommandEnable, p, nil)
}

// DisableParams stops observing for sinks and issues.
type DisableParams struct{}

// Disable stops observing for sinks and issues.
func Disable() *DisableParams {
	return &DisableParams{}
}

// Do executes Cast.disable against the provided context.
func (p *DisableParams) Do(ctxt context.Context, h cdp.Executor) (err error) {
	return h.Execute(ctxt, CommandDisable, nil, nil)
}

// SetSinkToUseParams sets a sink to be used when the web page requests the
// browser to choose a sink via Presentation API, Remote Playback API, or Cast
// SDK.
type SetSinkToUseParams struct {
	SinkName string `json:"sinkName"`
}

// SetSinkToUse sets a sink to be used when the web page requests the browser
// to choose a sink via Presentation API, Remote Playback API, or Cast SDK.
//
// parameters:
//   sinkName
func SetSinkToUse(sinkName string) *SetSinkToUseParams {
	return &SetSinkToUseParams{
		SinkName: sinkName,
	}
}

// Do executes Cast.setSinkToUse against the provided context.
func (p *SetSinkToUseParams) Do(ctxt context.Context, h cdp.Executor) (err error) {
	return h.Execute(ctxt, CommandSetSinkToUse, p, nil)
}

// StartTabMirroringParams starts mirroring the tab to the sink.
type StartTabMirroringParams struct {
	SinkName string `json:"sinkName"`
}

// StartTabMirroring starts mirroring the tab to the sink.
//
// parameters:
//   sinkName
func StartTabMirroring(sinkName string) *StartTabMirroringParams {
	return &StartTabMirroringParams{
		SinkName: sinkName,
	}
}

// Do executes Cast.startTabMirroring against the provided context.
func (p *StartTabMirroringParams) Do(ctxt context.Context, h cdp.Executor) (err error) {
	return h.Execute(ctxt, CommandStartTabMirroring, p, nil)
}

// StopCastingParams stops the active Cast session on the sink.
type StopCastingParams struct {
	SinkName string `json:"sinkName"`
}

// StopCasting stops the active Cast session on the sink.
//
// parameters:
//   sinkName
func StopCasting(sinkName string) *StopCastingParams {
	return &StopCastingParams{
		SinkName: sinkName,
	}
}

// Do executes Cast.stopCasting against the provided context.
func (p *StopCastingParams) Do(ctxt context.Context, h cdp.Executor) (err error) {
	return h.Execute(ctxt, CommandStopCasting, p, nil)
}

// Command names.
const (
	CommandEnable            = "Cast.enable"
	CommandDisable           = "Cast.disable"
	CommandSetSinkToUse      = "Cast.setSinkToUse"
	CommandStartTabMirroring = "Cast.startTabMirroring"
	CommandStopCasting       = "Cast.stopCasting"
)
