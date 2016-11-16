# phprocess module for phpStudy

I tried rewrote phprocess module to resolve when a php program call anothoer php program in same php-cgi.exe instance (such as in same phpStudy instance) will cause HTTP 504 eeror.

### Why will cause 504 error?

When program (A) call program (B) in same php-cgi.exe instance, php-cgi.exe will wait for (A) exit, but (A) wait for (B).

![why](http://chuantu.biz/t5/41/1479282390x974406806.png)

### How to use?

1. Build executable file:

```
C:\> go build phprocess.go
```

2. Copy and replace ```phprocess.exe``` in phpStudy tools folder.

3. Adding following code in nginx configuration file:

```
upstream php-fastcgi {
    server 127.0.0.1:9000 weight=1;
    server 127.0.0.1:9001 weight=1;
    server 127.0.0.1:9002 weight=1;
    server 127.0.0.1:9003 weight=1;
    # ... ... for more
}
```

4. Replace ```fastcgi_pass 127.0.0.1:9000;``` to ```fastcgi_pass php-fastcgi;``` .