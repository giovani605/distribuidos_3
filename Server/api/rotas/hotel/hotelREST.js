// esse arquivo lida com os requests nas passagens
"use strict";
const express = require('express');
const router = express.Router();
const gerenciador = require('./GerenciadorHotel.js');
const Passagem = require('./hotelClass');

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

router.get('/', (req,res,next) => {
  // recuperar dados do request
  console.log(req.body);
  var pass = extrairDados(req.body);
  if(gerenciador.consultar(pass)){
    res.status(200).json({
      message:'Quarto disponivel'
    });
  }else{
    res.status(200).json({
      message:'Quarto nao disponivel'
    });
  }
});



// recupera uma lista com todas as passagens
router.get('/todos', (req,res,next) => {
  res.status(200).send(gerenciador.listar());
  console.log("passei");
  return;

});
// aqui lido com os post resquest
router.post('/', (req,res,next) =>{
  console.log(req.body);
  var pass = extrairDados(req.body);
  if(gerenciador.criar(pass)){
    res.status(200).json({
      message:'Quarto comprada com sucesso'
    });
  }else{
    res.status(200).json({
      message:'Quarto negada'
    });
  }
});

module.exports = router;
