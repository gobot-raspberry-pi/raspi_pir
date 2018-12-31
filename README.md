# raspi_pir
PIR motion sensor application using GOBOT. If the sensor would sense movement the LED would turn-on.

## Diagram
![image](https://github.com/gobot-raspberry-pi/raspi_pir/blob/master/images/raspi_pir.png)
## How to run
1. Run this command
   ```
   GOARM=7 GOARCH=arm GOOS=linux go build -o=raspi_pir *.go
   ```
2. Then `SCP` the output to your Pi.
    ```
    scp raspi_pir pi@192.168.0.XXX:/home/pi
    ```
    - Change the the IP with your PI IP address.
3. `SSH` into your Pi and execute this command
   ```
   ./raspi_pir
   ```
    - You should see something like this
    ```
        2019/01/01 01:45:33 Initializing connections...
        2019/01/01 01:45:33 Initializing connection RaspberryPi-697CF05F ...
        2019/01/01 01:45:33 Initializing devices...
        2019/01/01 01:45:33 Initializing device PIRMotion-704FF824 ...
        2019/01/01 01:45:33 Initializing device LED-773B3AC ...
        2019/01/01 01:45:33 Robot PIR-motion-bot initialized.
        2019/01/01 01:45:33 Starting Robot PIR-motion-bot ...
        2019/01/01 01:45:33 Starting connections...
        2019/01/01 01:45:33 Starting connection RaspberryPi-697CF05F...
        2019/01/01 01:45:33 Starting devices...
        2019/01/01 01:45:33 Starting device PIRMotion-704FF824 on pin 11...
        2019/01/01 01:45:33 Starting device LED-773B3AC on pin 7...
        2019/01/01 01:45:33 Starting work...
    ```
    If PIR sensed a motion, you should see this...
    ```
        Motion Detected at: Jan  1 01:45:33.402
    ```

    If no motion is detected, you should see this...
    ```
        Motion Stopped at: Jan  1 01:47:11.333
    ```
    
    - `Ctrl + c` to stop
 
## Notes

