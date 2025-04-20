<!-- README.md -->

<strong>TCP Echo Server</strong><br>
<strong>Joseph Koop</strong><br>
<strong>Systems Programming and Computer Organization</strong><br>
<strong>Test #2</strong><br>
<strong>April 20, 2025</strong><br>

<br>
<strong>Link to video:</strong>

________________________________________________________________________________________________________________________________________

Overview:<br>
<br>
This project is a basic TCP echo server that creates connections to clients, echos back their input, and stores the results in a file.
________________________________________________________________________________________________________________________________________

How to Use:<br>

You can run this program in your IDE using this base command:<br>
>    go run main.go<br>

There are two command line flags that can be adjusted: <br>
>    go run main.go -port=4000 (replace 4000 with any valid port number)<br>
>    go run main.go -personality=false (turn off personalized echoe messages)<br>

________________________________________________________________________________________________________________________________________

The functionality that I found to be the most educationally beneficial was creating a seperate file for each client and logging their 
messages to it. It's something I've only dont a little of in Programming II and don't have a good grasp on it yet.<br><br>

The functionality that required the most research would have either been that same one or logging addresses and timestamps for
connections and disconnections. Most of the others were somewhat intuitive.