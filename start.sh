# Check if binary is present - if not rebuild 
if [ ! -f ./hometer ]; then
    echo "Binary not found - rebuild!"
    go build -o hometer  main.go logger.go routers.go
    echo "Build done"
fi

# Start hometer
nohup ./hometer > nohup.log 2>&1 &
echo $! > pid.txt