src_dir = ./src/cmd/clock-quiz
icon_path = ../../../docs/clock.png # Path relative to src_dir
app_name = ClockQuiz
app_version = 0.0.1

build:
	go build ./src/cmd/clock-quiz

package_macos:
	cd $(src_dir) && fyne package -os darwin -name $(app_name) -appVersion $(app_version) -icon $(icon_path) && mv $(app_name).app ../../../

zip_macos:
	zip --symlinks -r ClockQuiz.zip ClockQuiz.app/

