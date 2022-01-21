FROM nginx:1.21.4-alpine
COPY ./compose/local/nginx/templates/nginx.conf /etc/nginx/nginx.conf

CMD ["nginx", "-g", "daemon off;"]
