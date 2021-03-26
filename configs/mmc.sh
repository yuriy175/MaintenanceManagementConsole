#bash
cd /home/MMC
pkill ServerConsole
pkill node &
#git clone https://github.com/yuriy175/MaintenanceManagementConsole.git /home/MMC
git pull https://github.com/yuriy175/MaintenanceManagementConsole.git

cd /home/MMC/ServerConsole
export GOPATH=/home/MMC/ServerConsole/
go build
./ServerConsole &

cd /home/MMC/ConsoleUI/consoleui
#npm i
#npm run build
serve -s build &
