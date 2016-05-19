# screenshot-tool-mac
A simple tool for uploading screenshots to imgur on OSX. Written in golang.

## How does this work?
This tool watches the directory of your choice (I use ~/Screenshots) for new images. When it detects one, it also takes that image and uploads it to imgur's api, and then copies the link of the image to the clipboard.62118

## How do I use it?
* Set your screenshot directory appropriately in the golang source code
* Compile the go file to binary
* Set the binary to run at boot on your system (differs from OS to OS, check out documentation relating to your OS separately)
* Take screenshots as normal

## Todo
* Parameterize values (directory, client ID stuff, etc)
* Clean up code (!!!!)
* Create install scripts
