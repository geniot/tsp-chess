# Chess for TrimUI Smart Pro

## Overview


## Prerequisites

This app has been developed and tested for CrossMix-OS (https://github.com/cizia64/CrossMix-OS)
Likewise it should work with the stock firmware or any other overlay OS (like CrossMix).

It is not designed for Knulli or launching from EmulationStation. Although theoretically, it should work there as well.

## Installation

1. Download Chess.zip released here on GitHub.
2. Unzip it and copy all files with the folder (Chess) to /mnt/SDCARD/Apps on your TSP.
3. Reboot TSP. Chess app should appear in Apps.
4. If you plan to put this app in some other folder, remember to modify launch.sh.

## Building

Basically, I use Ubuntu and Docker. There is a Docker container running that I use for
cross-compilation.

1. You need to build the image (see https://github.com/geniot/trimui-smart-pro-toolchain), and start a container.
2. Run <code>make dist</code> to prepare the package. See Makefile for details.


You can also deploy the app with make. This is the default <code>make</code> behavior that I use a lot during the
   development. Modify Makefile if you plan to do so.

## Links

1. https://github.com/geniot/trimui-smart-pro-toolchain
2. https://github.com/gen2brain/raylib-go
3. https://github.com/trimui/firmware_smartpro
4. https://github.com/cizia64/CrossMix-OS
5. https://github.com/geniot/tsp-hardware-test

## License

You can use it freely for any purpose that you like. Most importantly, 
I hope it will help you to create something new and also share it with the community.