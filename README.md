# go-unix-socket
A daemon-and-client sample app for testing out UNIX Domain Sockets. `.sock` files are usually observed on applications such as MySQL and Postgres where a user connects to them without using credentials such as username/password.

UNIX Sockets are insanely fast as it operates between processes and not through a network (e.g., TCP/UDP). Very useful when you are developing a local-only applications that run on separate processes.

Another good example of UNIX socket is [PHP-FPM](https://php-fpm.org/), which daemonizes PHP runtime. Web servers such as Nginx can then communicate to it thru a socket, see example:
```conf
server {
    listen       80;
    server_name  example.journaldev.com;
    root         /var/www/html/wordpress;

    access_log /var/log/nginx/example.journaldev.com-access.log;
    error_log  /var/log/nginx/example.journaldev.com-error.log error;
    index index.html index.htm index.php;

    location / {
      try_files $uri $uri/ /index.php$is_args$args;
    }

    location ~ \.php$ {
      fastcgi_split_path_info ^(.+\.php)(/.+)$;
      fastcgi_pass unix:/var/run/php7.2-fpm-wordpress-site.sock; 👈 this
      fastcgi_index index.php;
      include fastcgi.conf;
    }
}
```

## Try it
Run daemon:
```shell
$ go run daemon/daemon.go
Starting daemon...
```

Run CLI (client):
```shell
$ go run cli/cli.go
Starting CLI...
Start typing
> 
```

### References
- https://dev.to/douglasmakey/understanding-unix-domain-sockets-in-golang-32n8
- https://www.digitalocean.com/community/tutorials/php-fpm-nginx
