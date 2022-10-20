cd ./cig-gui
pnpm run build
pnpm run postbuild
cd ..

# Build target
$env:build_os = 'windows'

cd ./src
go build -ldflags "-X main.build_os=$env:build_os"
Move-Item -Path "./src.exe" -Destination "../"
cd ..
Rename-Item -Path "./src.exe" -NewName "CIG.exe"