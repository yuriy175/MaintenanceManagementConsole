server {

#        root /var/www/mprom.ml/html;
	root /home/epo/ConsoleUI/consoleui/build;
        index index.html index.htm index.nginx-debian.html;

        server_name mprom.ml www.mprom.ml;

        location / {
                try_files $uri $uri/ =404;
        }

    location /manager/rabbit/ {
        if ($request_uri ~* "/manager/rabbit/(.*)") {
            proxy_pass http://127.0.0.1:15672/$1;
        }
        proxy_pass http://127.0.0.1:15672;
        proxy_buffering                     off;
        proxy_set_header Host               $http_host;
        proxy_set_header X-Real-Ip          $remote_addr;
        proxy_set_header X-Forwarded-For    $proxy_add_x_forwarded_for;
        proxy_set_header X-Forwarded-Proto  $scheme;

   }

    location /epo/ {
#	root /home/epo/ConsoleUI/consoleui/build;
	proxy_pass http://127.0.0.1:5000;
        proxy_buffering                     off;
        proxy_set_header Host               $http_host;
        proxy_set_header X-Real-Ip          $remote_addr;
        proxy_set_header X-Forwarded-For    $proxy_add_x_forwarded_for;
        proxy_set_header X-Forwarded-Proto  $scheme;

   }

#    location /equips/GetAllEquips {
     location /equips/ {
	if ($request_uri ~* "/equips/(.*)") {
            proxy_pass http://127.0.0.1:8181/equips/$1;
        }

#        proxy_pass http://127.0.0.1:8181/equips/GetAllEquips;
        proxy_buffering                     off;
        proxy_set_header Host               $http_host;
        proxy_set_header X-Real-Ip          $remote_addr;
        proxy_set_header X-Forwarded-For    $proxy_add_x_forwarded_for;
        proxy_set_header X-Forwarded-Proto  $scheme;

   }


       location /wsapp/ {

       if ($request_uri ~* "/wsapp/(.*)") {
            proxy_pass http://127.0.0.1:8080/$1;
        }

       proxy_pass http://127.0.0.1:8080;
       proxy_http_version 1.1;
       proxy_set_header Upgrade $http_upgrade;
       proxy_set_header Connection "Upgrade";
       proxy_set_header Host $host;
   }  



    listen [::]:443 ssl ipv6only=on; # managed by Certbot
    listen 443 ssl; # managed by Certbot
    ssl_certificate /etc/letsencrypt/live/mprom.ml/fullchain.pem; # managed by Certbot
    ssl_certificate_key /etc/letsencrypt/live/mprom.ml/privkey.pem; # managed by Certbot
    include /etc/letsencrypt/options-ssl-nginx.conf; # managed by Certbot
    ssl_dhparam /etc/letsencrypt/ssl-dhparams.pem; # managed by Certbot


}
server {
    if ($host = www.mprom.ml) {
        return 301 https://$host$request_uri;
    } # managed by Certbot


    if ($host = mprom.ml) {
        return 301 https://$host$request_uri;
    } # managed by Certbot


        listen 80;
        listen [::]:80;

        server_name mprom.ml www.mprom.ml;
    return 404; # managed by Certbot




}
