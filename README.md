# clock-quiz

What is this? It's a basic Go app to learn the [fyne](https://fyne.io/) toolkit and to teach my daughter to 
calculate the difference in hours and minutes between two analog clocks. 

![docs/demo.png](docs/demo.png)

## To Do List

- fix layout to look nicer
- update `parseInputTime` to also support `0h15m` format
- include hour markers on the clocks?
 
## Fix macOS quarantine?

`sudo xattr -r -d com.apple.quarantine MyApp.app`

## Credits

- Many thanks to the Fyne developer team who built the [analog clock example](https://github.com/fyne-io/examples/tree/develop), 
  which was my inspiration for building this project.
