## Colorize
print alternating lines in different colors.

### Usage

`colorize -color "foregroundColor+attributes:backgroundColor+attributes" -alt=false`

#### Examples

Format all lines:

`cat file | colorize -color "red:yellow"`

Format alternating lines:

`cat file | colorize -color "reg:yellow" -alt`

---
### Options

`-alt` : Alternate lines are colored

`-color` :

* black
* red
* green
* yellow
* blue
* magenta
* cyan
* white
* 0...255 (256 colors)

### Attributes

*   b = bold foreground
*   B = Blink foreground
*   u = underline foreground
*   i = inverse
*   h = high intensity (bright) foreground, background

    does not work with 256 colors
