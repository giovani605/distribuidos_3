// esse arquivo lida com os requests nas passagens
"use strict";
const express = require('express');
const router = express.Router();
const gerenciador = require('./GerenciadorPassagem.js');
const Passagem = require('./ClassPassagem');

// aqui eu vou linda com os get request
var cont = 0;

function extrairDados(body){
  var pass = new Passagem();
  pass.data = body.data;
  pass.origem = body.origem;
  pass.destino =  body.destino;
  pass.numero = body.numero;
  pass.tipo = body.tipo;
  return pass;
}

router.get('/:id', (req,res,next) => {
  // recuperar dados do request
  //  console.log(req);
  var id = req.params.id;
  var obj = gerenciador.consultar(id);
  res.status(200).json({
    message:'OK',
    dados:obj
  });

});


// recupera uma lista com todas as passagens disponiveis
router.get('/', (req,res,next) => {
  var a = gerenciador.listar();
  res.status(200).json({
    dados: a
  });
  return;
});

// aqui o cliente quer comprar a passagem
router.post('/:id', (req,res,next) =>{
  var id = req.params.id;
  if(gerenciador.comprar(id)){
    res.status(200).json({
      message:'comprada com sucesso'
    });
  }else{
    res.status(200).json({
      message:'passagem nao disponivel'
    });
  }
});

module.exports = router;
