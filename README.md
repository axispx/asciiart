# ASCIIart

**ASCIIart** allows you to convert images into ASCII art.

**Original Image:**
![Original Image](https://image.ibb.co/kp7sR6/ashish.jpg)

**ASCIIart Image:**
![ASCIIart Image](https://image.ibb.co/iJEgm6/out.png)


## Getting Started

### Prerequisites
You will need Go language installed on you computer.

### Installing
Go get the project.
`go get github.com/anarchyrucks/asciiart`

### Example
`asciiart -file ashish.png -set "ASHISH" -charsize 12 -serial`

## Flags
| Name     | Description                                                                                   | Default                   |
|----------|-----------------------------------------------------------------------------------------------|---------------------------|
| char     | The character to be used. Random characters will be used for every pixel if not provided.     | None                      |
| charsize | The size of the character.                                                                    | 10                        |
| file     | The path of the file to be used.                                                              | None                      |
| out      | The name of the output image(PNG).                                                            | out.png                   |
| serial   | Use the characters from set serially rather than randomly.                                    | false                     |
| set      | The set of characters to be used.                                                             | Visible ASCII characters. |

## License
MIT
