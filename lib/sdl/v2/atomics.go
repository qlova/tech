/*

  Simple DirectMedia Layer
  Copyright (C) 1997-2023 Sam Lantinga <slouken@libsdl.org>

  This software is provided 'as-is', without any express or implied
  warranty.  In no event will the authors be held liable for any damages
  arising from the use of this software.

  Permission is granted to anyone to use this software for any purpose,
  including commercial applications, and to alter it and redistribute it
  freely, subject to the following restrictions:

  1. The origin of this software must not be misrepresented; you must not
     claim that you wrote the original software. If you use this software
     in a product, an acknowledgment in the product documentation would be
     appreciated but is not required.
  2. Altered source versions must be plainly marked as such, and must not be
     misrepresented as being the original software.
  3. This notice may not be removed or altered from any source distribution.

*/

package sdl

import (
	"qlova.tech/abi"
)

type SpinLock abi.AtomicInt
type AtomicInt abi.AtomicInt

var Atomics struct {
	Lib

	TryLock func(*SpinLock) Bool `ffi:"SDL_AtomicTryLock"` // Try to lock a spin lock by setting it to a non-zero value.
	Lock    func(*SpinLock)      `ffi:"SDL_AtomicLock"`    // Lock a spin lock by setting it to a non-zero value.
	Unlock  func(*SpinLock)      `ffi:"SDL_AtomicUnlock"`  // Unlock a spin lock by setting it to 0.

	CompareAndSwap func(*AtomicInt, abi.Int, abi.Int) Bool `ffi:"SDL_AtomicCAS"` // Set an atomic variable to a new value if it is currently an old value.
	Set            func(*AtomicInt, abi.Int)               `ffi:"SDL_AtomicSet"` // Set an atomic variable to a value.
	Get            func(*AtomicInt) abi.Int                `ffi:"SDL_AtomicGet"` // Get the value of an atomic variable.
	Add            func(*AtomicInt, abi.Int) abi.Int       `ffi:"SDL_AtomicAdd"` // Add to an atomic variable.

	CompareAndSwapPointer func(*abi.AtomicUintptr, abi.Pointer, abi.Pointer) Bool `ffi:"SDL_AtomicCASPtr"` // Set an atomic variable to a new value if it is currently an old value.
	SetPointer            func(*abi.AtomicUintptr, abi.Pointer)                   `ffi:"SDL_AtomicSetPtr"` // Set an atomic variable to a value.
	GetPointer            func(*abi.AtomicUintptr) abi.Pointer                    `ffi:"SDL_AtomicGetPtr"` // Get the value of an atomic variable.
}
