package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"

	"strings"
)

// Terminal spinner function

func spinner() {
	s := []string{"|", "/", "-", "\\"}
	for {
		for _, v := range s {
			fmt.Printf("\r%s", v)
			time.Sleep(100 * time.Millisecond)
		}
	}
}

// Green color

func green() {
	fmt.Printf("\033[32m")
}

// Clear the terminal

func clear() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

// Install teamviewer

func install_teamviewer() {
	// wget the file from teamviewer
	_, err := exec.Command("wget", "https://download.teamviewer.com/download/linux/teamviewer_amd64.deb").Output()
	if err != nil {
		fmt.Println(err)
	}

	// install the file
	_, err = exec.Command("sudo", "dpkg", "-i", "teamviewer_amd64.deb").Output()
	// while the file is installing the spinner will run
	go spinner()
	if err != nil {
		fmt.Println(err)
	}
	// remove the file

	_, err = exec.Command("rm", "teamviewer_amd64.deb").Output()
	if err != nil {
		fmt.Println(err)
	}

}

// Install google chrome

func install_chrome() {
	// wget the file from teamviewer
	_, err := exec.Command("wget", "https://dl.google.com/linux/direct/google-chrome-stable_current_amd64.deb").Output()

	if err != nil {
		fmt.Println(err)
	}

	// install the file
	_, err = exec.Command("sudo", "dpkg", "-i", "google-chrome-stable_current_amd64.deb").Output()
	if err != nil {
		fmt.Println(err)
	}
	// remove the file

	_, err = exec.Command("rm", "google-chrome-stable_current_amd64.deb").Output()
	if err != nil {
		fmt.Println(err)
	}

}

// Software to install

func software_req() {

	// Check if teamviewer is installed

	_, err := exec.Command("teamviewer", "--version").Output()
	if err != nil {
		// Install teamviewer
		fmt.Println("Teamviewer is not installed.....Starting Installation")
		fmt.Println(' ')
		// while the file is installing the spinner will run
		go spinner()
		install_teamviewer()
	} else {
		fmt.Println("Teamviewer is installed")
	}

	// Check if google chrome is installed

	_, err = exec.Command("google-chrome", "--version").Output()
	if err != nil {
		// Install google chrome
		fmt.Println("Google Chrome is not installed.....Starting Installation")
		// while the file is installing the spinner will run
		go spinner()
		install_chrome()
		// Add google chrome to the dock
		_, err = exec.Command("cp", "/usr/share/applications/google-chrome.desktop", "/home/viasat/Desktop/").Output()
		if err != nil {
			fmt.Println(err)
		}
		// Make the google chrome icon on the dock clickable
		_, err = exec.Command("chmod", "+x", "/home/viasat/Desktop/google-chrome.desktop").Output()
		if err != nil {
			fmt.Println(err)
		}
	} else {
		fmt.Println("Google Chrome is installed")
	}

}

// Check if Ubuntu
func isUbuntu() bool {
	system, err := exec.Command("uname", "-a").Output()
	if err != nil {
		fmt.Println(err)
	}
	if strings.Contains(string(system), "Ubuntu") {
		return true
	}
	return false
}

func main() {
	clear()
	if isUbuntu() {
		fmt.Println("****** Ubuntu is the current operating system ******")
	} else {
		fmt.Println("Not Ubuntu")
		fmt.Fprint(os.Stderr, "This script is only for Ubuntu")

	}
	software_req()

	fmt.Println("Done")

}
