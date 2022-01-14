FROM public.ecr.aws/lambda/nodejs:latest

COPY package*.json app.js ./

RUN npm install

CMD ["app.handler"]
