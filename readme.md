
<h3 align="center">Watchkeeper</h3>

# About The Project

A simple port scanner

## Usage

```shell
kjoedicker@arch ~ % ./watchkeeper --host=github.com --interval=10

Sep 30 09:37:30 - github.com: 
  open:
    -22
    -80
  closed:
    -8081
    -3000

Sep 30 09:37:43 - github.com: 
  open:
    -80
    -22
  closed:
    -8081
    -3000
```