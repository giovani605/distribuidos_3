// modulo que cria a middle ware Express que lida com os REST
const express = require('express');

const app = express();
// importa o arquivo de pssagem
const passagens = require('./api/rotas/passagem.js')

app.use('/passagem',passagens);

module.exports = app;

// app.use((req,res,next) =>{
//   console.log("chegou um request");
//
//   res.status(200).json(
//     {message: 'It work'});
//   });
