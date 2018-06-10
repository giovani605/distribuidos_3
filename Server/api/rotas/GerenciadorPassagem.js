// esse Modulo controla as passagens

// cria uma mapa que tem as passagens
"use strict";
var mapa = new Map();
const Passagem = require('./ClassPassagem');
var cont = 1;

function criarPassagem(passagem) {
  if(testarPassagem(passagem)){
    salvar(passagem);
    console.log("salvei a passagem");
    return true;
  }
  else{
    return false;
  }
}
function salvar(passagem){
  // criar um id
  console.log(passagem);
  let id = cont;
  passagem.id = id;
  // add no map
  mapa.set(id,passagem);
  cont++;
}
// essa funcao testa se a passagem eh valida
function testarPassagem(passagem){
  if(mapa.size == 0){
    return true;
  }

  for(var a of mapa.values()){
    console.log(a);
    if(a.data == passagem.data &&
      a.origem == passagem.origem &&
      a.destino == passagem.destino &&
      a.numero == passagem.numero){
        return false;
      }
    }
    return true;
  }


  function consultarPassagem(passsagem){
    // retorna a lista a lista de passagem
    if(testarPassagem(passagem)){
      return true;
    }
    else{
      return false;
    }
  }

  function listarPassagem(){
    // retorna a lista a lista de passagem

  }
  exports.consultar = consultarPassagem;
  exports.criar = criarPassagem;
