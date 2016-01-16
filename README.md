# screenshot-tool-mac
A simple tool for uploading screenshots to imgur on OSX. Written in golang.

## How does this work?
This tool watches the directory of your choice (I use ~/Screenshots) for new images. When it detects one, it also takes that image and uploads it to imgur's api, and then copies the link of the image to the clipboard.62118

## How do I use it?
same

## Todo
* Parameterize values (directory, client ID stuff, etc)
* Clean up code (!!!!)
* Create install scripts