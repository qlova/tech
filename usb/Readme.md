# HID

The hid package aims to be a platform and environment agnostic home for 
reading human input devices. Such as mouse, keyboard and gamepads.

Events will typically be handled by the platform specific window driver
and will be loaded into the HID package for processing. This extra 
layer of indirection enables inputs to be easily stubbed/tested within
your application.