CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o "./out/$1" fix-workshop-ue &&
cp -r ./templates "./out/" &&
cp -r ./settings "./out/" &&
cp -r ./static "./out/" &&
zip -r fix-workshop.v2.zip ./out/*