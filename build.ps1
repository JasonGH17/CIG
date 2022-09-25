cd ./cig-gui
pnpm run build
pnpm run postbuild
cd ..
go build ./src
Rename-Item -Path "./src.exe" -NewName "CIG.exe"