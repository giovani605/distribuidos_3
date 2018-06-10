// esse arquivo lida com os requests nas passagens
"use strict";
const express = require('express');
const router = express.Router();
const gerenciador = require('./GerenciadorPassagem.js');
const Passagem = require('./ClassPassagem');

// aqui eu vou linda com os get request
var cont = 0;

router.get('/', (req,res,next) => {
  // recuperar dados do request
  console.log(req.body);


  res.status(200).json({
    message: 'lidando o get'
  });
});


// aqui lido com os post resquest
router.post('/', (req,res,next) =>{
  console.log(req.body);

  var pass = new Passagem();
  pass.data = req.body.data;
  pass.origem = req.body.origem;
  pass.destino =  req.body.destino;
  pass.numero = req.body.numero;
  pass.tipo = req.body.tipo;
  pass.id = 0;
  if(gerenciador.criar(pass)){
    res.status(200).json({
      message:'Passagem comprada com sucesso'
    });
  }else{
    res.status(200).json({
      message:'Passagem negada'
    });
  }




});

module.exports = router;
