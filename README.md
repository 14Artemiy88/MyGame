# MyGame
Just a game in the terminal

![game.gif](screenshots/game.gif)

## Requirements
```agsl
go 1.20
```

## Installation
### Snap 
```bash
sudo snap install --beta piupiu
```
#### Usage
```bash
piupiu
```

### GO
Make sure the Go executables directory ($GOPATH/bin) is added to your PATH environment variable. You can achieve this using the following command:
```bash
export PATH=$PATH:$(go env GOPATH)/bin
```
and then
```bash
go install github.com/14Artemiy88/MyGame@latest
```
#### Usage
```bash
Mygame
```