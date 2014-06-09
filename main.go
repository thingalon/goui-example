package main

import (
	"fmt"
	"time"
	"os/user"
	"github.com/thingalon/goui"
)

var nextThreadID = 1

func main() {
	//	Tell goui where to find HTML/CSS/JavaScript content.
	goui.SetAssetPath("assets/")	//	Use a real path during debugging...
	goui.SetBindataSource(Asset)	//	But wrap it all up in a bindata blob for deployment.
	
	//	Assign handlers for incoming goui messages (sent via goui.SendMessage on the JavaScript side)
	goui.SetMessageHandler("examples.showPopup", showExamplePopup)
	goui.SetMessageHandler("examples.spawnThread", spawnThread)
	
	//	Start running goui. This method does not return until the user exits the program.
	goui.Run(func() {
		//	Create a new main window.
		mainWindow := goui.OpenWindow(goui.WindowOptions{
			Template: "templates/main.html",
			StyleFlags: goui.WindowResizable | goui.WindowClosable | goui.WindowMinimizable,
			PercentLeft: 10, PercentTop: 10,
			PercentWidth: 75, PercentHeight: 75,
			RememberGeometry: true,
		})
		
		//	Set a callback to exit the application when the main window is closed.
		mainWindow.SetCloseHandler(func(window *goui.Window) {
			fmt.Println("Main window closed. Exiting.")
			goui.Stop()
		})
		
		//	Push a message to the main window; just throw some general data at it to show this works.
		u, _ := user.Current()
		mainWindow.Send(goui.Message{"showName", goui.Data{"name": u.Name}})
	})
	
	fmt.Println("Exited gracefully.")
}

func showExamplePopup(window *goui.Window, request *goui.Message) (response goui.Data) {
	inputText, ok := request.Params["data"].(string)
	if ! ok || len(inputText) == 0 {
		inputText = "Nothing!"
	}
	
	popupWindow := goui.OpenWindow(goui.WindowOptions{
		Template: "templates/popup.html",
		StyleFlags: goui.WindowModal,
		Centered: true,
		PixelWidth: 500, PixelHeight: 250,
	})

	popupWindow.SetCloseHandler(func(window *goui.Window) {
		//	Code to run when the popup closes (if any) goes here.
		fmt.Println("Looks like you closed a popup.")
	})
	
	popupWindow.Send(goui.Message{"showMessage", goui.Data{"inputText": inputText}})

	return
}

func spawnThread(window *goui.Window, request *goui.Message) (response goui.Data) {
	duration, ok := request.Params["duration"].(float64); 
	if !ok {
		return
	}

	speed, ok := request.Params["speed"].(float64); 
	if !ok {
		return
	}
	
	threadID := nextThreadID
	nextThreadID++
	
	go runThread(window, threadID, duration, speed)
	
	return
}

func runThread(window *goui.Window, threadID int, duration float64, speed float64) {
	threadOut(window, fmt.Sprintf("Starting thread %d for %f seconds, ticking every %f seconds.", threadID, duration, speed))

	startTime := time.Now()
	elapsed := float64(0)
	second := float64(time.Second)
	
	for duration > elapsed {
		elapsed = float64(time.Since(startTime)) / second
		if duration - elapsed > speed {
			time.Sleep(time.Duration(speed * second))
			threadOut(window, fmt.Sprintf("Thread %d says hello!", threadID))
		} else {
			time.Sleep(time.Duration((duration - elapsed) * second))
		}
	}

	threadOut(window, fmt.Sprintf("Exiting thread %d", threadID))
}

func threadOut(window *goui.Window, text string) {
	window.Send(goui.Message{"threadOutput", goui.Data{"text":text}})
}
