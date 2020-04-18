# http-address-printer
Tool to print the http adress along with its Hash 
# Getting Started
 Tool is already built so just run ./main to test the application 
 Some examples
```
./main adjust.com 
./main parallel=4 adjust.com
```
# How it is done 
 I have used bueffered channel to solve the problem .By default the concurrency limit is 10 .By passing the parallel flags this limit can be overriden.
 
## Go inbuilt library used 

 * [URL](https://golang.org/pkg/net/url/)-to validate if the given URL is correct .If not it will correct the URL and provide the result 
 * [Crypto/md5](https://golang.org/pkg/crypto/md5/)-Used for calcuating MD5 hash of the response returned from http request
 * [http](https://golang.org/pkg/net/http/)-for intiating http request 
 * [flag](https://golang.org/pkg/flag/) - used for fetching the flags passed in command line arguments 
 
 # Tests
  Added basic unit testing

 
