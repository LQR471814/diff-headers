## diff-headers

> A dead simple utility to compare / diff HTTP headers.

### Usage

```
# command
$ diff-headers file1.txt file2.txt

# file1.txt
accept-ranges: bytes
age: 203
cache-control: public,max-age=0,must-revalidate
server: Netlify

# file2.txt (lines without ":" are ignored)
HTTP/2 200
age: 103
content-type: text/html
last-modified: Sun, 18 Jan 2015 00:04:33 GMT
vary: Accept-Encoding,User-Agent

# output
===== file2.txt is missing the following headers from file1.txt

HEADER         VALUE                             
accept-ranges  bytes                             
cache-control  public,max-age=0,must-revalidate  
server         netlify                           

===== file1.txt is missing the following headers from file2.txt

HEADER         VALUE                          
content-type   text/html                      
last-modified  sun, 18 jan 2015 00:04:33 gmt  
vary           accept-encoding,user-agent     

===== Differences in shared header values

HEADER  file1.txt  file2.txt  
age     203        103        
```
