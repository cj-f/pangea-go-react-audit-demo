# pangea-go-react-audit-demo
A example local application use a Go server with proxy endpoints for embedding an audit log viewer in the client

PANGEA_TOKEN and PANGEA_DOMAIN are needed to run the Go server

# Server

```
cd server/
go run main.go
```

localhost:4000 should be available with the following APIs, /audit/search, /audit/results, /audit/root

# Client

```
cd client/
npm install
yarn start
```

localhost:3000 should be available showing a simple React app that is just rendering the react-mui-audit-log-viewer component
