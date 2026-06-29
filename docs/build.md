# Build for multiple platforms
```
GOOS=windows GOARCH=amd64  go build -ldflags="-s -w" -o dist/todo-app-windows-amd64.exe 
GOOS=linux   GOARCH=amd64  go build -ldflags="-s -w" -o dist/todo-app-linux-amd64   
GOOS=darwin  GOARCH=amd64  go build -ldflags="-s -w" -o dist/todo-app-darwin-amd64  
GOOS=darwin  GOARCH=arm64  go build -ldflags="-s -w" -o dist/todo-app-darwin-arm64  
```

# Release all at once
```
gh release create v1.0.0 dist/* \
--title "Release v1.0.0" \
--notes "What's changed in this release"
```