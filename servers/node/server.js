const express = require('express')
const bodyParser = require("body-parser")

const app = express()
const port = 3000

// Middleware
app.use(bodyParser.json());
app.use(bodyParser.urlencoded({ extended: true }));




app.get("/", function (req, res) {
    res.send("Server is up and running!")
})


app.post('/api/users', (req, res) => {
  const { name } = req.body;
  console.log(name)
  const newUser = { "body": `hello ${name}` };

  // based on logic resp to the frontend or maybe the db for the request 
  res.status(201).json(newUser);
});



app.listen(port, function () {
    console.log(`Server is running on port: ${port}`)
})