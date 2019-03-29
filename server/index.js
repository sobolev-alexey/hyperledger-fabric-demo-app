const express = require('express');
const bodyParser = require('body-parser');
const { create_record, change_holder, get_container, get_all_containers, } = require('./controller.js');

// Save our port
const port = process.env.PORT || 3001;

const app = express();
app.use(bodyParser.urlencoded({ extended: true }));
app.use(bodyParser.json());

app.get('/get/:id', (req, res) => {
    get_container(req, res);
});
app.get('/get_all', (req, res) => {
    get_all_containers(req, res);
});
app.post('/change', (req, res) => {
    change_holder(req, res);
});
app.post('/create', (req, res) => {
    create_record(req, res);
});

// set up a static file server that points to the "client" directory
// app.use(express.static(path.join(__dirname, '../client')));

// Start the server and listen on port 
app.listen(port, () => console.log("Live on port: " + port));

