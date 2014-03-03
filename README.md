# Remote Commands

## Skills developed

* Concurrency (channels, go routines)
* Network interface
* Much more dependending on the commands you implement

## Assignment

The goal of this exercise it to trigger commands remotely. The
application is going to be a server that can
concurrenctly receive a list of
predefined commands. The status of the server and the status of each
task should also be available via specific commands.
Here is a list of suggested commands:

* Current UTC time
* Current CPU usage
* Available RAM
* CPU usage over last hour
* Available RAM over last hour
* Download url into specific folder
* Make computer "say" something
* Capture and send a screenshot
* Trigger webhook at specific time


## Resources

* Apple's `say` command: https://developer.apple.com/library/mac/documentation/Darwin/Reference/ManPages/man1/say.1.html
* linux `espeak` command: `echo "I wanna Go!"|espeak` http://espeak.sourceforge.net/
* Windows users might want to look at [ptts](http://jampal.sourceforge.net/ptts.html) or [espeak](http://espeak.sourceforge.net/)
* OS X has a `screencapture` command
* Windows has a free tool called [NirCmd](http://www.nirsoft.net/utils/nircmd.html) to do a bunch of things.
