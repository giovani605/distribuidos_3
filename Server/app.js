// modulo que cria a middle ware Express que lida com os REST
const express = require('express');

const app = express();
// importa o arquivo de pssagem
const passagens = require('./api/rotas/passagem.js');
const hotel = require('./api/rotas/hotel/hotelREST.js');

app.use('/passagem',passagens);
app.use('/hotel',hotel);

module.exports = app;

// app.use((req,res,next) =>{
//   console.log("chegou um request");
//
//   res.status(200).json(
//     {message: 'It work'});
//   });
