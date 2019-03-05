for GOOS in darwin linux; do
    for GOARCH in 386 amd64 arm; do
        echo $GOOS-$GOARCH
        CGO_ENABLED=0 GOOS=$GOOS GOARCH=$GOARCH go build -v -o ./build/pdf-designer-$GOOS-$GOARCH .
    done
done

for GOOS in windows; do
    for GOARCH in 386 amd64; do
        echo $GOOS-$GOARCH
        CGO_ENABLED=0 GOOS=$GOOS GOARCH=$GOARCH go build -v -o ./build/pdf-designer-$GOOS-$GOARCH.exe .
    done
done
