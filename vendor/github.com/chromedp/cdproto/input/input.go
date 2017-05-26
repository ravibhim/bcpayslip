// Package input provides the Chrome DevTools Protocol
// commands, types, and events for the Input domain.
//
// Generated by the cdproto-gen command.
package input

// Code generated by cdproto-gen. DO NOT EDIT.

import (
	"context"

	"github.com/chromedp/cdproto/cdp"
)

// DispatchKeyEventParams dispatches a key event to the page.
type DispatchKeyEventParams struct {
	Type                  KeyType         `json:"type"`                            // Type of the key event.
	Modifiers             Modifier        `json:"modifiers"`                       // Bit field representing pressed modifier keys. Alt=1, Ctrl=2, Meta/Command=4, Shift=8 (default: 0).
	Timestamp             *TimeSinceEpoch `json:"timestamp,omitempty"`             // Time at which the event occurred.
	Text                  string          `json:"text,omitempty"`                  // Text as generated by processing a virtual key code with a keyboard layout. Not needed for for keyUp and rawKeyDown events (default: "")
	UnmodifiedText        string          `json:"unmodifiedText,omitempty"`        // Text that would have been generated by the keyboard if no modifiers were pressed (except for shift). Useful for shortcut (accelerator) key handling (default: "").
	KeyIdentifier         string          `json:"keyIdentifier,omitempty"`         // Unique key identifier (e.g., 'U+0041') (default: "").
	Code                  string          `json:"code,omitempty"`                  // Unique DOM defined string value for each physical key (e.g., 'KeyA') (default: "").
	Key                   string          `json:"key,omitempty"`                   // Unique DOM defined string value describing the meaning of the key in the context of active modifiers, keyboard layout, etc (e.g., 'AltGr') (default: "").
	WindowsVirtualKeyCode int64           `json:"windowsVirtualKeyCode,omitempty"` // Windows virtual key code (default: 0).
	NativeVirtualKeyCode  int64           `json:"nativeVirtualKeyCode,omitempty"`  // Native virtual key code (default: 0).
	AutoRepeat            bool            `json:"autoRepeat"`                      // Whether the event was generated from auto repeat (default: false).
	IsKeypad              bool            `json:"isKeypad"`                        // Whether the event was generated from the keypad (default: false).
	IsSystemKey           bool            `json:"isSystemKey"`                     // Whether the event was a system key event (default: false).
	Location              int64           `json:"location,omitempty"`              // Whether the event was from the left or right side of the keyboard. 1=Left, 2=Right (default: 0).
}

// DispatchKeyEvent dispatches a key event to the page.
//
// parameters:
//   type - Type of the key event.
func DispatchKeyEvent(typeVal KeyType) *DispatchKeyEventParams {
	return &DispatchKeyEventParams{
		Type: typeVal,
	}
}

// WithModifiers bit field representing pressed modifier keys. Alt=1, Ctrl=2,
// Meta/Command=4, Shift=8 (default: 0).
func (p DispatchKeyEventParams) WithModifiers(modifiers Modifier) *DispatchKeyEventParams {
	p.Modifiers = modifiers
	return &p
}

// WithTimestamp time at which the event occurred.
func (p DispatchKeyEventParams) WithTimestamp(timestamp *TimeSinceEpoch) *DispatchKeyEventParams {
	p.Timestamp = timestamp
	return &p
}

// WithText text as generated by processing a virtual key code with a
// keyboard layout. Not needed for for keyUp and rawKeyDown events (default:
// "").
func (p DispatchKeyEventParams) WithText(text string) *DispatchKeyEventParams {
	p.Text = text
	return &p
}

// WithUnmodifiedText text that would have been generated by the keyboard if
// no modifiers were pressed (except for shift). Useful for shortcut
// (accelerator) key handling (default: "").
func (p DispatchKeyEventParams) WithUnmodifiedText(unmodifiedText string) *DispatchKeyEventParams {
	p.UnmodifiedText = unmodifiedText
	return &p
}

// WithKeyIdentifier unique key identifier (e.g., 'U+0041') (default: "").
func (p DispatchKeyEventParams) WithKeyIdentifier(keyIdentifier string) *DispatchKeyEventParams {
	p.KeyIdentifier = keyIdentifier
	return &p
}

// WithCode unique DOM defined string value for each physical key (e.g.,
// 'KeyA') (default: "").
func (p DispatchKeyEventParams) WithCode(code string) *DispatchKeyEventParams {
	p.Code = code
	return &p
}

// WithKey unique DOM defined string value describing the meaning of the key
// in the context of active modifiers, keyboard layout, etc (e.g., 'AltGr')
// (default: "").
func (p DispatchKeyEventParams) WithKey(key string) *DispatchKeyEventParams {
	p.Key = key
	return &p
}

// WithWindowsVirtualKeyCode windows virtual key code (default: 0).
func (p DispatchKeyEventParams) WithWindowsVirtualKeyCode(windowsVirtualKeyCode int64) *DispatchKeyEventParams {
	p.WindowsVirtualKeyCode = windowsVirtualKeyCode
	return &p
}

// WithNativeVirtualKeyCode native virtual key code (default: 0).
func (p DispatchKeyEventParams) WithNativeVirtualKeyCode(nativeVirtualKeyCode int64) *DispatchKeyEventParams {
	p.NativeVirtualKeyCode = nativeVirtualKeyCode
	return &p
}

// WithAutoRepeat whether the event was generated from auto repeat (default:
// false).
func (p DispatchKeyEventParams) WithAutoRepeat(autoRepeat bool) *DispatchKeyEventParams {
	p.AutoRepeat = autoRepeat
	return &p
}

// WithIsKeypad whether the event was generated from the keypad (default:
// false).
func (p DispatchKeyEventParams) WithIsKeypad(isKeypad bool) *DispatchKeyEventParams {
	p.IsKeypad = isKeypad
	return &p
}

// WithIsSystemKey whether the event was a system key event (default: false).
func (p DispatchKeyEventParams) WithIsSystemKey(isSystemKey bool) *DispatchKeyEventParams {
	p.IsSystemKey = isSystemKey
	return &p
}

// WithLocation whether the event was from the left or right side of the
// keyboard. 1=Left, 2=Right (default: 0).
func (p DispatchKeyEventParams) WithLocation(location int64) *DispatchKeyEventParams {
	p.Location = location
	return &p
}

// Do executes Input.dispatchKeyEvent against the provided context.
func (p *DispatchKeyEventParams) Do(ctxt context.Context, h cdp.Executor) (err error) {
	return h.Execute(ctxt, CommandDispatchKeyEvent, p, nil)
}

// InsertTextParams this method emulates inserting text that doesn't come
// from a key press, for example an emoji keyboard or an IME.
type InsertTextParams struct {
	Text string `json:"text"` // The text to insert.
}

// InsertText this method emulates inserting text that doesn't come from a
// key press, for example an emoji keyboard or an IME.
//
// parameters:
//   text - The text to insert.
func InsertText(text string) *InsertTextParams {
	return &InsertTextParams{
		Text: text,
	}
}

// Do executes Input.insertText against the provided context.
func (p *InsertTextParams) Do(ctxt context.Context, h cdp.Executor) (err error) {
	return h.Execute(ctxt, CommandInsertText, p, nil)
}

// DispatchMouseEventParams dispatches a mouse event to the page.
type DispatchMouseEventParams struct {
	Type        MouseType                     `json:"type"`                  // Type of the mouse event.
	X           float64                       `json:"x"`                     // X coordinate of the event relative to the main frame's viewport in CSS pixels.
	Y           float64                       `json:"y"`                     // Y coordinate of the event relative to the main frame's viewport in CSS pixels. 0 refers to the top of the viewport and Y increases as it proceeds towards the bottom of the viewport.
	Modifiers   Modifier                      `json:"modifiers"`             // Bit field representing pressed modifier keys. Alt=1, Ctrl=2, Meta/Command=4, Shift=8 (default: 0).
	Timestamp   *TimeSinceEpoch               `json:"timestamp,omitempty"`   // Time at which the event occurred.
	Button      ButtonType                    `json:"button,omitempty"`      // Mouse button (default: "none").
	Buttons     int64                         `json:"buttons,omitempty"`     // A number indicating which buttons are pressed on the mouse when a mouse event is triggered. Left=1, Right=2, Middle=4, Back=8, Forward=16, None=0.
	ClickCount  int64                         `json:"clickCount,omitempty"`  // Number of times the mouse button was clicked (default: 0).
	DeltaX      float64                       `json:"deltaX,omitempty"`      // X delta in CSS pixels for mouse wheel event (default: 0).
	DeltaY      float64                       `json:"deltaY,omitempty"`      // Y delta in CSS pixels for mouse wheel event (default: 0).
	PointerType DispatchMouseEventPointerType `json:"pointerType,omitempty"` // Pointer type (default: "mouse").
}

// DispatchMouseEvent dispatches a mouse event to the page.
//
// parameters:
//   type - Type of the mouse event.
//   x - X coordinate of the event relative to the main frame's viewport in CSS pixels.
//   y - Y coordinate of the event relative to the main frame's viewport in CSS pixels. 0 refers to the top of the viewport and Y increases as it proceeds towards the bottom of the viewport.
func DispatchMouseEvent(typeVal MouseType, x float64, y float64) *DispatchMouseEventParams {
	return &DispatchMouseEventParams{
		Type: typeVal,
		X:    x,
		Y:    y,
	}
}

// WithModifiers bit field representing pressed modifier keys. Alt=1, Ctrl=2,
// Meta/Command=4, Shift=8 (default: 0).
func (p DispatchMouseEventParams) WithModifiers(modifiers Modifier) *DispatchMouseEventParams {
	p.Modifiers = modifiers
	return &p
}

// WithTimestamp time at which the event occurred.
func (p DispatchMouseEventParams) WithTimestamp(timestamp *TimeSinceEpoch) *DispatchMouseEventParams {
	p.Timestamp = timestamp
	return &p
}

// WithButton mouse button (default: "none").
func (p DispatchMouseEventParams) WithButton(button ButtonType) *DispatchMouseEventParams {
	p.Button = button
	return &p
}

// WithButtons a number indicating which buttons are pressed on the mouse
// when a mouse event is triggered. Left=1, Right=2, Middle=4, Back=8,
// Forward=16, None=0.
func (p DispatchMouseEventParams) WithButtons(buttons int64) *DispatchMouseEventParams {
	p.Buttons = buttons
	return &p
}

// WithClickCount number of times the mouse button was clicked (default: 0).
func (p DispatchMouseEventParams) WithClickCount(clickCount int64) *DispatchMouseEventParams {
	p.ClickCount = clickCount
	return &p
}

// WithDeltaX X delta in CSS pixels for mouse wheel event (default: 0).
func (p DispatchMouseEventParams) WithDeltaX(deltaX float64) *DispatchMouseEventParams {
	p.DeltaX = deltaX
	return &p
}

// WithDeltaY Y delta in CSS pixels for mouse wheel event (default: 0).
func (p DispatchMouseEventParams) WithDeltaY(deltaY float64) *DispatchMouseEventParams {
	p.DeltaY = deltaY
	return &p
}

// WithPointerType pointer type (default: "mouse").
func (p DispatchMouseEventParams) WithPointerType(pointerType DispatchMouseEventPointerType) *DispatchMouseEventParams {
	p.PointerType = pointerType
	return &p
}

// Do executes Input.dispatchMouseEvent against the provided context.
func (p *DispatchMouseEventParams) Do(ctxt context.Context, h cdp.Executor) (err error) {
	return h.Execute(ctxt, CommandDispatchMouseEvent, p, nil)
}

// DispatchTouchEventParams dispatches a touch event to the page.
type DispatchTouchEventParams struct {
	Type        TouchType       `json:"type"`                // Type of the touch event. TouchEnd and TouchCancel must not contain any touch points, while TouchStart and TouchMove must contains at least one.
	TouchPoints []*TouchPoint   `json:"touchPoints"`         // Active touch points on the touch device. One event per any changed point (compared to previous touch event in a sequence) is generated, emulating pressing/moving/releasing points one by one.
	Modifiers   Modifier        `json:"modifiers"`           // Bit field representing pressed modifier keys. Alt=1, Ctrl=2, Meta/Command=4, Shift=8 (default: 0).
	Timestamp   *TimeSinceEpoch `json:"timestamp,omitempty"` // Time at which the event occurred.
}

// DispatchTouchEvent dispatches a touch event to the page.
//
// parameters:
//   type - Type of the touch event. TouchEnd and TouchCancel must not contain any touch points, while TouchStart and TouchMove must contains at least one.
//   touchPoints - Active touch points on the touch device. One event per any changed point (compared to previous touch event in a sequence) is generated, emulating pressing/moving/releasing points one by one.
func DispatchTouchEvent(typeVal TouchType, touchPoints []*TouchPoint) *DispatchTouchEventParams {
	return &DispatchTouchEventParams{
		Type:        typeVal,
		TouchPoints: touchPoints,
	}
}

// WithModifiers bit field representing pressed modifier keys. Alt=1, Ctrl=2,
// Meta/Command=4, Shift=8 (default: 0).
func (p DispatchTouchEventParams) WithModifiers(modifiers Modifier) *DispatchTouchEventParams {
	p.Modifiers = modifiers
	return &p
}

// WithTimestamp time at which the event occurred.
func (p DispatchTouchEventParams) WithTimestamp(timestamp *TimeSinceEpoch) *DispatchTouchEventParams {
	p.Timestamp = timestamp
	return &p
}

// Do executes Input.dispatchTouchEvent against the provided context.
func (p *DispatchTouchEventParams) Do(ctxt context.Context, h cdp.Executor) (err error) {
	return h.Execute(ctxt, CommandDispatchTouchEvent, p, nil)
}

// EmulateTouchFromMouseEventParams emulates touch event from the mouse event
// parameters.
type EmulateTouchFromMouseEventParams struct {
	Type       MouseType       `json:"type"`                 // Type of the mouse event.
	X          int64           `json:"x"`                    // X coordinate of the mouse pointer in DIP.
	Y          int64           `json:"y"`                    // Y coordinate of the mouse pointer in DIP.
	Button     ButtonType      `json:"button"`               // Mouse button.
	Timestamp  *TimeSinceEpoch `json:"timestamp,omitempty"`  // Time at which the event occurred (default: current time).
	DeltaX     float64         `json:"deltaX,omitempty"`     // X delta in DIP for mouse wheel event (default: 0).
	DeltaY     float64         `json:"deltaY,omitempty"`     // Y delta in DIP for mouse wheel event (default: 0).
	Modifiers  Modifier        `json:"modifiers"`            // Bit field representing pressed modifier keys. Alt=1, Ctrl=2, Meta/Command=4, Shift=8 (default: 0).
	ClickCount int64           `json:"clickCount,omitempty"` // Number of times the mouse button was clicked (default: 0).
}

// EmulateTouchFromMouseEvent emulates touch event from the mouse event
// parameters.
//
// parameters:
//   type - Type of the mouse event.
//   x - X coordinate of the mouse pointer in DIP.
//   y - Y coordinate of the mouse pointer in DIP.
//   button - Mouse button.
func EmulateTouchFromMouseEvent(typeVal MouseType, x int64, y int64, button ButtonType) *EmulateTouchFromMouseEventParams {
	return &EmulateTouchFromMouseEventParams{
		Type:   typeVal,
		X:      x,
		Y:      y,
		Button: button,
	}
}

// WithTimestamp time at which the event occurred (default: current time).
func (p EmulateTouchFromMouseEventParams) WithTimestamp(timestamp *TimeSinceEpoch) *EmulateTouchFromMouseEventParams {
	p.Timestamp = timestamp
	return &p
}

// WithDeltaX X delta in DIP for mouse wheel event (default: 0).
func (p EmulateTouchFromMouseEventParams) WithDeltaX(deltaX float64) *EmulateTouchFromMouseEventParams {
	p.DeltaX = deltaX
	return &p
}

// WithDeltaY Y delta in DIP for mouse wheel event (default: 0).
func (p EmulateTouchFromMouseEventParams) WithDeltaY(deltaY float64) *EmulateTouchFromMouseEventParams {
	p.DeltaY = deltaY
	return &p
}

// WithModifiers bit field representing pressed modifier keys. Alt=1, Ctrl=2,
// Meta/Command=4, Shift=8 (default: 0).
func (p EmulateTouchFromMouseEventParams) WithModifiers(modifiers Modifier) *EmulateTouchFromMouseEventParams {
	p.Modifiers = modifiers
	return &p
}

// WithClickCount number of times the mouse button was clicked (default: 0).
func (p EmulateTouchFromMouseEventParams) WithClickCount(clickCount int64) *EmulateTouchFromMouseEventParams {
	p.ClickCount = clickCount
	return &p
}

// Do executes Input.emulateTouchFromMouseEvent against the provided context.
func (p *EmulateTouchFromMouseEventParams) Do(ctxt context.Context, h cdp.Executor) (err error) {
	return h.Execute(ctxt, CommandEmulateTouchFromMouseEvent, p, nil)
}

// SetIgnoreInputEventsParams ignores input events (useful while auditing
// page).
type SetIgnoreInputEventsParams struct {
	Ignore bool `json:"ignore"` // Ignores input events processing when set to true.
}

// SetIgnoreInputEvents ignores input events (useful while auditing page).
//
// parameters:
//   ignore - Ignores input events processing when set to true.
func SetIgnoreInputEvents(ignore bool) *SetIgnoreInputEventsParams {
	return &SetIgnoreInputEventsParams{
		Ignore: ignore,
	}
}

// Do executes Input.setIgnoreInputEvents against the provided context.
func (p *SetIgnoreInputEventsParams) Do(ctxt context.Context, h cdp.Executor) (err error) {
	return h.Execute(ctxt, CommandSetIgnoreInputEvents, p, nil)
}

// SynthesizePinchGestureParams synthesizes a pinch gesture over a time
// period by issuing appropriate touch events.
type SynthesizePinchGestureParams struct {
	X                 float64     `json:"x"`                           // X coordinate of the start of the gesture in CSS pixels.
	Y                 float64     `json:"y"`                           // Y coordinate of the start of the gesture in CSS pixels.
	ScaleFactor       float64     `json:"scaleFactor"`                 // Relative scale factor after zooming (>1.0 zooms in, <1.0 zooms out).
	RelativeSpeed     int64       `json:"relativeSpeed,omitempty"`     // Relative pointer speed in pixels per second (default: 800).
	GestureSourceType GestureType `json:"gestureSourceType,omitempty"` // Which type of input events to be generated (default: 'default', which queries the platform for the preferred input type).
}

// SynthesizePinchGesture synthesizes a pinch gesture over a time period by
// issuing appropriate touch events.
//
// parameters:
//   x - X coordinate of the start of the gesture in CSS pixels.
//   y - Y coordinate of the start of the gesture in CSS pixels.
//   scaleFactor - Relative scale factor after zooming (>1.0 zooms in, <1.0 zooms out).
func SynthesizePinchGesture(x float64, y float64, scaleFactor float64) *SynthesizePinchGestureParams {
	return &SynthesizePinchGestureParams{
		X:           x,
		Y:           y,
		ScaleFactor: scaleFactor,
	}
}

// WithRelativeSpeed relative pointer speed in pixels per second (default:
// 800).
func (p SynthesizePinchGestureParams) WithRelativeSpeed(relativeSpeed int64) *SynthesizePinchGestureParams {
	p.RelativeSpeed = relativeSpeed
	return &p
}

// WithGestureSourceType which type of input events to be generated (default:
// 'default', which queries the platform for the preferred input type).
func (p SynthesizePinchGestureParams) WithGestureSourceType(gestureSourceType GestureType) *SynthesizePinchGestureParams {
	p.GestureSourceType = gestureSourceType
	return &p
}

// Do executes Input.synthesizePinchGesture against the provided context.
func (p *SynthesizePinchGestureParams) Do(ctxt context.Context, h cdp.Executor) (err error) {
	return h.Execute(ctxt, CommandSynthesizePinchGesture, p, nil)
}

// SynthesizeScrollGestureParams synthesizes a scroll gesture over a time
// period by issuing appropriate touch events.
type SynthesizeScrollGestureParams struct {
	X                     float64     `json:"x"`                               // X coordinate of the start of the gesture in CSS pixels.
	Y                     float64     `json:"y"`                               // Y coordinate of the start of the gesture in CSS pixels.
	XDistance             float64     `json:"xDistance,omitempty"`             // The distance to scroll along the X axis (positive to scroll left).
	YDistance             float64     `json:"yDistance,omitempty"`             // The distance to scroll along the Y axis (positive to scroll up).
	XOverscroll           float64     `json:"xOverscroll,omitempty"`           // The number of additional pixels to scroll back along the X axis, in addition to the given distance.
	YOverscroll           float64     `json:"yOverscroll,omitempty"`           // The number of additional pixels to scroll back along the Y axis, in addition to the given distance.
	PreventFling          bool        `json:"preventFling,omitempty"`          // Prevent fling (default: true).
	Speed                 int64       `json:"speed,omitempty"`                 // Swipe speed in pixels per second (default: 800).
	GestureSourceType     GestureType `json:"gestureSourceType,omitempty"`     // Which type of input events to be generated (default: 'default', which queries the platform for the preferred input type).
	RepeatCount           int64       `json:"repeatCount,omitempty"`           // The number of times to repeat the gesture (default: 0).
	RepeatDelayMs         int64       `json:"repeatDelayMs,omitempty"`         // The number of milliseconds delay between each repeat. (default: 250).
	InteractionMarkerName string      `json:"interactionMarkerName,omitempty"` // The name of the interaction markers to generate, if not empty (default: "").
}

// SynthesizeScrollGesture synthesizes a scroll gesture over a time period by
// issuing appropriate touch events.
//
// parameters:
//   x - X coordinate of the start of the gesture in CSS pixels.
//   y - Y coordinate of the start of the gesture in CSS pixels.
func SynthesizeScrollGesture(x float64, y float64) *SynthesizeScrollGestureParams {
	return &SynthesizeScrollGestureParams{
		X: x,
		Y: y,
	}
}

// WithXDistance the distance to scroll along the X axis (positive to scroll
// left).
func (p SynthesizeScrollGestureParams) WithXDistance(xDistance float64) *SynthesizeScrollGestureParams {
	p.XDistance = xDistance
	return &p
}

// WithYDistance the distance to scroll along the Y axis (positive to scroll
// up).
func (p SynthesizeScrollGestureParams) WithYDistance(yDistance float64) *SynthesizeScrollGestureParams {
	p.YDistance = yDistance
	return &p
}

// WithXOverscroll the number of additional pixels to scroll back along the X
// axis, in addition to the given distance.
func (p SynthesizeScrollGestureParams) WithXOverscroll(xOverscroll float64) *SynthesizeScrollGestureParams {
	p.XOverscroll = xOverscroll
	return &p
}

// WithYOverscroll the number of additional pixels to scroll back along the Y
// axis, in addition to the given distance.
func (p SynthesizeScrollGestureParams) WithYOverscroll(yOverscroll float64) *SynthesizeScrollGestureParams {
	p.YOverscroll = yOverscroll
	return &p
}

// WithPreventFling prevent fling (default: true).
func (p SynthesizeScrollGestureParams) WithPreventFling(preventFling bool) *SynthesizeScrollGestureParams {
	p.PreventFling = preventFling
	return &p
}

// WithSpeed swipe speed in pixels per second (default: 800).
func (p SynthesizeScrollGestureParams) WithSpeed(speed int64) *SynthesizeScrollGestureParams {
	p.Speed = speed
	return &p
}

// WithGestureSourceType which type of input events to be generated (default:
// 'default', which queries the platform for the preferred input type).
func (p SynthesizeScrollGestureParams) WithGestureSourceType(gestureSourceType GestureType) *SynthesizeScrollGestureParams {
	p.GestureSourceType = gestureSourceType
	return &p
}

// WithRepeatCount the number of times to repeat the gesture (default: 0).
func (p SynthesizeScrollGestureParams) WithRepeatCount(repeatCount int64) *SynthesizeScrollGestureParams {
	p.RepeatCount = repeatCount
	return &p
}

// WithRepeatDelayMs the number of milliseconds delay between each repeat.
// (default: 250).
func (p SynthesizeScrollGestureParams) WithRepeatDelayMs(repeatDelayMs int64) *SynthesizeScrollGestureParams {
	p.RepeatDelayMs = repeatDelayMs
	return &p
}

// WithInteractionMarkerName the name of the interaction markers to generate,
// if not empty (default: "").
func (p SynthesizeScrollGestureParams) WithInteractionMarkerName(interactionMarkerName string) *SynthesizeScrollGestureParams {
	p.InteractionMarkerName = interactionMarkerName
	return &p
}

// Do executes Input.synthesizeScrollGesture against the provided context.
func (p *SynthesizeScrollGestureParams) Do(ctxt context.Context, h cdp.Executor) (err error) {
	return h.Execute(ctxt, CommandSynthesizeScrollGesture, p, nil)
}

// SynthesizeTapGestureParams synthesizes a tap gesture over a time period by
// issuing appropriate touch events.
type SynthesizeTapGestureParams struct {
	X                 float64     `json:"x"`                           // X coordinate of the start of the gesture in CSS pixels.
	Y                 float64     `json:"y"`                           // Y coordinate of the start of the gesture in CSS pixels.
	Duration          int64       `json:"duration,omitempty"`          // Duration between touchdown and touchup events in ms (default: 50).
	TapCount          int64       `json:"tapCount,omitempty"`          // Number of times to perform the tap (e.g. 2 for double tap, default: 1).
	GestureSourceType GestureType `json:"gestureSourceType,omitempty"` // Which type of input events to be generated (default: 'default', which queries the platform for the preferred input type).
}

// SynthesizeTapGesture synthesizes a tap gesture over a time period by
// issuing appropriate touch events.
//
// parameters:
//   x - X coordinate of the start of the gesture in CSS pixels.
//   y - Y coordinate of the start of the gesture in CSS pixels.
func SynthesizeTapGesture(x float64, y float64) *SynthesizeTapGestureParams {
	return &SynthesizeTapGestureParams{
		X: x,
		Y: y,
	}
}

// WithDuration duration between touchdown and touchup events in ms (default:
// 50).
func (p SynthesizeTapGestureParams) WithDuration(duration int64) *SynthesizeTapGestureParams {
	p.Duration = duration
	return &p
}

// WithTapCount number of times to perform the tap (e.g. 2 for double tap,
// default: 1).
func (p SynthesizeTapGestureParams) WithTapCount(tapCount int64) *SynthesizeTapGestureParams {
	p.TapCount = tapCount
	return &p
}

// WithGestureSourceType which type of input events to be generated (default:
// 'default', which queries the platform for the preferred input type).
func (p SynthesizeTapGestureParams) WithGestureSourceType(gestureSourceType GestureType) *SynthesizeTapGestureParams {
	p.GestureSourceType = gestureSourceType
	return &p
}

// Do executes Input.synthesizeTapGesture against the provided context.
func (p *SynthesizeTapGestureParams) Do(ctxt context.Context, h cdp.Executor) (err error) {
	return h.Execute(ctxt, CommandSynthesizeTapGesture, p, nil)
}

// Command names.
const (
	CommandDispatchKeyEvent           = "Input.dispatchKeyEvent"
	CommandInsertText                 = "Input.insertText"
	CommandDispatchMouseEvent         = "Input.dispatchMouseEvent"
	CommandDispatchTouchEvent         = "Input.dispatchTouchEvent"
	CommandEmulateTouchFromMouseEvent = "Input.emulateTouchFromMouseEvent"
	CommandSetIgnoreInputEvents       = "Input.setIgnoreInputEvents"
	CommandSynthesizePinchGesture     = "Input.synthesizePinchGesture"
	CommandSynthesizeScrollGesture    = "Input.synthesizeScrollGesture"
	CommandSynthesizeTapGesture       = "Input.synthesizeTapGesture"
)
