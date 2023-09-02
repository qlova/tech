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

package hid

import "qlova.tech/lib/sdl/v2"

const (
	/*
		A variable controlling whether the HIDAPI joystick drivers should be used.

		This variable can be set to the following values:
			"0"       - HIDAPI drivers are not used
			"1"       - HIDAPI drivers are used (the default)

		This variable is the default for all drivers, but can be overridden by the hints for specific drivers below.
	*/
	HintJoystick sdl.Hint = "SDL_JOYSTICK_HIDAPI"
	/*
		A variable controlling whether the HIDAPI driver for Nintendo GameCube controllers should be used.

		This variable can be set to the following values:
			"0"       - HIDAPI driver is not used
			"1"       - HIDAPI driver is used

		The default is the value of SDL_HINT_JOYSTICK_HIDAPI
	*/
	HintJoystickGamecube sdl.Hint = "SDL_JOYSTICK_HIDAPI_GAMECUBE"
	/*
		A variable controlling whether "low_frequency_rumble" and "high_frequency_rumble" is used to implement
			the GameCube controller's 3 rumble modes, Stop(0), Rumble(1), and StopHard(2)
			this is useful for applications that need full compatibility for things like ADSR envelopes.
			Stop is implemented by setting "low_frequency_rumble" to "0" and "high_frequency_rumble" ">0"
			Rumble is both at any arbitrary value,
			StopHard is implemented by setting both "low_frequency_rumble" and "high_frequency_rumble" to "0"

		This variable can be set to the following values:
			"0"       - Normal rumble behavior is behavior is used (default)
			"1"       - Proper GameCube controller rumble behavior is used
	*/
	HintJoystickGamecubeRumbleBrake sdl.Hint = "SDL_JOYSTICK_GAMECUBE_RUMBLE_BRAKE"
	/*
		A variable controlling whether the HIDAPI driver for Nintendo Switch Joy-Cons should be used.

		This variable can be set to the following values:
		  "0"       - HIDAPI driver is not used
		  "1"       - HIDAPI driver is used

		The default is the value of SDL_HINT_JOYSTICK_HIDAPI
	*/
	HintJoystickJoyCons sdl.Hint = "SDL_JOYSTICK_HIDAPI_JOY_CONS"
	/*
		A variable controlling whether Nintendo Switch Joy-Con controllers will be combined into a single Pro-like controller when using the HIDAPI driver

		This variable can be set to the following values:
		  "0"       - Left and right Joy-Con controllers will not be combined and each will be a mini-gamepad
		  "1"       - Left and right Joy-Con controllers will be combined into a single controller (the default)
	*/
	HintJoystickJoyConsCombine sdl.Hint = "SDL_JOYSTICK_HIDAPI_COMBINE_JOY_CONS"
	/*
		A variable controlling whether Nintendo Switch Joy-Con controllers will be in vertical mode when using the HIDAPI driver

		This variable can be set to the following values:
		  "0"       - Left and right Joy-Con controllers will not be in vertical mode (the default)
		  "1"       - Left and right Joy-Con controllers will be in vertical mode

		This hint must be set before calling SDL_Init(SDL_INIT_GAMECONTROLLER)
	*/
	HintJoystickVerticalJoyCons sdl.Hint = "SDL_JOYSTICK_HIDAPI_VERTICAL_JOY_CONS"
	/*
		A variable controlling whether the HIDAPI driver for Amazon Luna controllers connected via Bluetooth should be used.

		This variable can be set to the following values:
		  "0"       - HIDAPI driver is not used
		  "1"       - HIDAPI driver is used

		The default is the value of SDL_HINT_JOYSTICK_HIDAPI
	*/
	HintJoystickLuna sdl.Hint = "SDL_JOYSTICK_HIDAPI_LUNA"
	/*
		A variable controlling whether the HIDAPI driver for Nintendo Online classic controllers should be used.

		This variable can be set to the following values:
		  "0"       - HIDAPI driver is not used
		  "1"       - HIDAPI driver is used

		The default is the value of SDL_HINT_JOYSTICK_HIDAPI
	*/
	HintJoystickNintendoClassic sdl.Hint = "SDL_JOYSTICK_HIDAPI_NINTENDO_CLASSIC"
	/*
		A variable controlling whether the HIDAPI driver for NVIDIA SHIELD controllers should be used.

		This variable can be set to the following values:
		  "0"       - HIDAPI driver is not used
		  "1"       - HIDAPI driver is used

		The default is the value of SDL_HINT_JOYSTICK_HIDAPI
	*/
	HintJoystickShield sdl.Hint = "SDL_JOYSTICK_HIDAPI_SHIELD"
	/*
		A variable controlling whether the HIDAPI driver for PS3 controllers should be used.

		This variable can be set to the following values:
			"0"       - HIDAPI driver is not used
			"1"       - HIDAPI driver is used

		The default is the value of SDL_HINT_JOYSTICK_HIDAPI on macOS, and "0" on other platforms.

		It is not possible to use this driver on Windows, due to limitations in the default drivers
		installed. See https://github.com/ViGEm/DsHidMini for an alternative driver on Windows.
	*/
	HintJoystickPS3 sdl.Hint = "SDL_JOYSTICK_HIDAPI_PS3"
	/*
		A variable controlling whether the HIDAPI driver for PS4 controllers should be used.

		This variable can be set to the following values:
			"0"       - HIDAPI driver is not used
			"1"       - HIDAPI driver is used

		The default is the value of SDL_HINT_JOYSTICK_HIDAPI
	*/
	HintJoystickPS4 sdl.Hint = "SDL_JOYSTICK_HIDAPI_PS4"
	/*
		A variable controlling whether extended input reports should be used for PS4 controllers when using the HIDAPI driver.

		This variable can be set to the following values:
			"0"       - extended reports are not enabled (the default)
			"1"       - extended reports

		Extended input reports allow rumble on Bluetooth PS4 controllers, but
		break DirectInput handling for applications that don't use SDL.

		Once extended reports are enabled, they can not be disabled without
		power cycling the controller.

		For compatibility with applications written for versions of SDL prior
		to the introduction of PS5 controller support, this value will also
		control the state of extended reports on PS5 controllers when the
		SDL_HINT_JOYSTICK_HIDAPI_PS5_RUMBLE hint is not explicitly set.
	*/
	HintJoystickRumblePS4 sdl.Hint = "SDL_JOYSTICK_HIDAPI_PS4_RUMBLE"
	/*
		A variable controlling whether the HIDAPI driver for PS5 controllers should be used.

		This variable can be set to the following values:
			"0"       - HIDAPI driver is not used
			"1"       - HIDAPI driver is used

		The default is the value of SDL_HINT_JOYSTICK_HIDAPI
	*/
	HintJoystickPS5 sdl.Hint = "SDL_JOYSTICK_HIDAPI_PS5"
	/*
		A variable controlling whether the player LEDs should be lit to indicate which player is associated with a PS5 controller.

		This variable can be set to the following values:
			"0"       - player LEDs are not enabled
			"1"       - player LEDs are enabled (the default)
	*/
	HintJoystickPS5LED sdl.Hint = "SDL_JOYSTICK_HIDAPI_PS5_PLAYER_LED"
	/*
		A variable controlling whether extended input reports should be used for PS5 controllers when using the HIDAPI driver.

		This variable can be set to the following values:
			"0"       - extended reports are not enabled (the default)
			"1"       - extended reports

		Extended input reports allow rumble on Bluetooth PS5 controllers, but
		break DirectInput handling for applications that don't use SDL.

		Once extended reports are enabled, they can not be disabled without
		power cycling the controller.

		For compatibility with applications written for versions of SDL prior
		to the introduction of PS5 controller support, this value defaults to
		the value of SDL_HINT_JOYSTICK_HIDAPI_PS4_RUMBLE.
	*/
	HintJoystickRumblePS5 sdl.Hint = "SDL_JOYSTICK_HIDAPI_PS5_RUMBLE"
	/*
		A variable controlling whether the HIDAPI driver for Google Stadia controllers should be used.

		This variable can be set to the following values:
			"0"       - HIDAPI driver is not used
			"1"       - HIDAPI driver is used

		The default is the value of SDL_HINT_JOYSTICK_HIDAPI
	*/
	HintJoystickStadia sdl.Hint = "SDL_JOYSTICK_HIDAPI_STADIA"
	/*
		A variable controlling whether the HIDAPI driver for Bluetooth Steam Controllers should be used.

		This variable can be set to the following values:
			"0"       - HIDAPI driver is not used
			"1"       - HIDAPI driver is used for Steam Controllers, which requires Bluetooth access
						and may prompt the user for permission on iOS and Android.

		The default is "0"
	*/
	HintJoystickSteam sdl.Hint = "SDL_JOYSTICK_HIDAPI_STEAM"
	/*
		A variable controlling whether the HIDAPI driver for Nintendo Switch controllers should be used.

		This variable can be set to the following values:
			"0"       - HIDAPI driver is not used
			"1"       - HIDAPI driver is used

		The default is the value of SDL_HINT_JOYSTICK_HIDAPI
	*/
	HintJoystickSwitch sdl.Hint = "SDL_JOYSTICK_HIDAPI_SWITCH"
	/*
		A variable controlling whether the Home button LED should be turned on when a Nintendo Switch Pro controller is opened

		This variable can be set to the following values:
			"0"       - home button LED is turned off
			"1"       - home button LED is turned on

		By default the Home button LED state is not changed. This hint can also be set to a floating point value between 0.0 and 1.0 which controls the brightness of the Home button LED.
	*/
	HintJoystickSwitchHomeLED sdl.Hint = "SDL_JOYSTICK_HIDAPI_SWITCH_HOME_LED"
	/*
		A variable controlling whether the Home button LED should be turned on when a Nintendo Switch Joy-Con controller is opened

		This variable can be set to the following values:
			"0"       - home button LED is turned off
			"1"       - home button LED is turned on

		By default the Home button LED state is not changed. This hint can also be set to a floating point value between 0.0 and 1.0 which controls the brightness of the Home button LED.
	*/
	HintJoystickJoyconHomeLED sdl.Hint = "SDL_JOYSTICK_HIDAPI_JOYCON_HOME_LED"
	/*
		A variable controlling whether the player LEDs should be lit to indicate which player is associated with a Nintendo Switch controller.

		This variable can be set to the following values:
			"0"       - player LEDs are not enabled
			"1"       - player LEDs are enabled (the default)
	*/
	HintJoystickSwitchPlayerLED sdl.Hint = "SDL_JOYSTICK_HIDAPI_SWITCH_PLAYER_LED"
	/*
		A variable controlling whether the HIDAPI driver for Nintendo Wii and Wii U controllers should be used.

		This variable can be set to the following values:
			"0"       - HIDAPI driver is not used
			"1"       - HIDAPI driver is used

		This driver doesn't work with the dolphinbar, so the default is SDL_FALSE for now.
	*/
	HintJoystickWii sdl.Hint = "SDL_JOYSTICK_HIDAPI_WII"
	/*
		A variable controlling whether the player LEDs should be lit to indicate which player is associated with a Wii controller.

		This variable can be set to the following values:
			"0"       - player LEDs are not enabled
			"1"       - player LEDs are enabled (the default)
	*/
	HintJoystickWiiPlayerLED sdl.Hint = "SDL_JOYSTICK_HIDAPI_WII_PLAYER_LED"
	/*
		A variable controlling whether the HIDAPI driver for XBox controllers should be used.

		This variable can be set to the following values:
			"0"       - HIDAPI driver is not used
			"1"       - HIDAPI driver is used

		The default is "0" on Windows, otherwise the value of SDL_HINT_JOYSTICK_HIDAPI
	*/
	HintJoystickXbox sdl.Hint = "SDL_JOYSTICK_HIDAPI_XBOX"
	/*
		A variable controlling whether the HIDAPI driver for XBox 360 controllers should be used.

		This variable can be set to the following values:
			"0"       - HIDAPI driver is not used
			"1"       - HIDAPI driver is used

		The default is the value of SDL_HINT_JOYSTICK_HIDAPI_XBOX
	*/
	HintJoystickXbox360 sdl.Hint = "SDL_JOYSTICK_HIDAPI_XBOX_360"
	/*
		A variable controlling whether the player LEDs should be lit to indicate which player is associated with an Xbox 360 controller.

		This variable can be set to the following values:
			"0"       - player LEDs are not enabled
			"1"       - player LEDs are enabled (the default)
	*/
	HintJoystickXbox360PlayerLED sdl.Hint = "SDL_JOYSTICK_HIDAPI_XBOX_360_PLAYER_LED"
	/*
		A variable controlling whether the HIDAPI driver for XBox 360 wireless controllers should be used.

		This variable can be set to the following values:
			"0"       - HIDAPI driver is not used
			"1"       - HIDAPI driver is used

		The default is the value of SDL_HINT_JOYSTICK_HIDAPI_XBOX_360
	*/
	HintJoystickXbox360Wireless sdl.Hint = "SDL_JOYSTICK_HIDAPI_XBOX_360_WIRELESS"
	/*
		A variable controlling whether the HIDAPI driver for XBox One controllers should be used.

		This variable can be set to the following values:
			"0"       - HIDAPI driver is not used
			"1"       - HIDAPI driver is used

		The default is the value of SDL_HINT_JOYSTICK_HIDAPI_XBOX
	*/
	HintJoystickXboxOne sdl.Hint = "SDL_JOYSTICK_HIDAPI_XBOX_ONE"
	/*
		A variable controlling whether the Home button LED should be turned on when an Xbox One controller is opened

		This variable can be set to the following values:
			"0"       - home button LED is turned off
			"1"       - home button LED is turned on

		By default the Home button LED state is not changed. This hint can also be set to a floating point value between 0.0 and 1.0 which controls the brightness of the Home button LED. The default brightness is 0.4.
	*/
	HintJoystickXboxOneHomeLED sdl.Hint = "SDL_JOYSTICK_HIDAPI_XBOX_ONE_HOME_LED"
)
