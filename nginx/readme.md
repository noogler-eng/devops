# Nginx Web Server Setup with Docker and Ubuntu

## Overview

Nginx is a powerful web server that serves static files such as `index.html`, JavaScript, CSS, images, and more. This guide walks you through setting up Nginx inside a Docker container using Ubuntu.

## Running Ubuntu in a Docker Container

To start an Ubuntu container and map port 8080 to port 80 inside the container:

```sh
docker run -it -p 8080:80 ubuntu
```

## Installing Nginx

Once inside the container, update the package list and install Nginx:

```sh
apt-get update
apt-get install nginx
```

Verify the installation:

```sh
nginx -v
```

Start Nginx:

```sh
nginx
```

### Checking Nginx Response Headers

Nginx should now be running. When making an HTTP request, the response header should contain:

```
server: nginx/1.24.0 (Ubuntu)
```

## Exploring Nginx Configuration Files

Navigate to the Nginx configuration directory:

```sh
ls /etc/nginx
cd /etc/nginx
ls -lh
```

The main configuration file is `nginx.conf`, which is important for server configuration.

## Editing Nginx Configuration

Install `vim` for editing:

```sh
apt-get install vim
```

Edit the configuration file:

```sh
vim nginx.conf
```

### Backup and Modify Configuration

Backup the existing configuration file:

```sh
mv nginx.conf nginx-backup.conf
```

Create a new configuration file:

```sh
touch nginx.conf
vim nginx.conf
```

## Reloading Nginx

After making changes to `nginx.conf`, reload the Nginx service:

```sh
nginx -s reload
```

## Setting Up a Website Directory

Create a directory for your website and an `index.html` file:

```sh
mkdir /website
touch /website/index.html
```

## Testing Nginx Configuration

Before restarting Nginx, test the configuration for errors:

```sh
nginx -t
```

### Important Note

When specifying file paths in the configuration, **always use absolute paths** instead of relative paths.

## Conclusion

You have successfully set up Nginx inside a Docker container running Ubuntu. You can now configure it to serve your static files and customize it as needed.
