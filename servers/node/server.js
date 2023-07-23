const express = require('express')
const https = require('https')
const bodyParser = require("body-parser")
const fs  = require('fs')

const app = express()
const port = 3000

app.use(bodyParser.json());
app.use(bodyParser.urlencoded({ extended: true }));

process.env.NODE_TLS_REJECT_UNAUTHORIZED = '0';

// info regarding the middleware connection 
const CACERT = fs.readFileSync('/Users/avinesh/go/src/skeleton/middleware/certs/ca.crt');
const HOSTNAME = 'localhost';
const PORT = 8080;

const options = {
    hostname: 'localhost',
    port: 8080,
    path: '/',
    method: 'GET',
    rejectUnauthorised: false,
    ca: CACERT
};


app.get("/", function (req, res) {
    res.send("Server is up and running!")
})

app.post('/user', (req, res) => {
    const { name } = req.body;
    
    const data = JSON.stringify({
        name: name
    })

    const options = {
        hostname: HOSTNAME,
        port: PORT,
        path: `/user`,
        method: "POST",
        rejectedUnauthorised: false,
        ca: CACERT
    }

    const middlewareCall = https.request(options, (response) => {
        let resp = response

        res.send(resp)
    })

    middlewareCall.write(data)
    middlewareCall.end()
});

app.get('/users/:id', (req, res) => {
    const id = parseInt(req.params.id)

    const options = {
        hostname: HOSTNAME,
        port: PORT,
        path: `/id?${id}`,
        method: "GET",
        rejectedUnauthorised: false,
        ca: CACERT
    }

    const middlewareCall = https.request(options, (response) => {
        let data = response

        response.on('data', (chunk) => {
            data += chunk 
        })

        response.on('end', () => {
            console.log(data)
            res.send(data)
        })
    })

    middlewareCall.on("error", (error) => {
        console.error("ERROR making the middleware call: ", error )
        res.status(500).send('Error making the middleware call.')
    })

    middlewareCall.end()
})

app.get("/ping", (req, res)  => {

    const options = {
        hostname: HOSTNAME,
        port: PORT,
        path: "/",
        method: "GET",
        rejectedUnauthorised: false,
        ca: CACERT
    }

    const middlewareCall = https.request(options, (response) => {
        let data = ''

        response.on('data', (chunk) => {
            data += chunk 
        })

        response.on('end', () => {
            console.log(data)
            res.send(data)
        })
    }) 

    middlewareCall.on("error", (error) => {
        console.error("ERROR making the middleware call: ", error )
        res.status(500).send('Error making the middleware call.')
    })

    middlewareCall.end();
})

app.listen(port, function () {
    console.log(`Server is running on port: ${port}`)
})