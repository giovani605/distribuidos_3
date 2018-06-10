// esse Modulo controla as passagens

// cria uma mapa que tem as passagens
"use strict";
var mapa = new Map();

mapa.set(1, "teste");
mapa.set(2,"tese");
const Passagem = require('./ClassPassagem');
var teste1 = new Passagem();
var teste2 = new Passagem();
teste1.data = "hoje";
teste2.data = "ontem";
console.log(teste1);
console.log(teste2);

function criarPassagem(passagem) {
  if(testarPassagem(passagem)){
    salvar();
    return true;
  }
  else{
    return false;
  }
}
function salvar(passagem){
  // criar um id
  var id = "1";
  // add no map
  mapa.set(id,passagem);
}
// essa funcao testa se a passagem eh valida
function testarPassagem(passagem){

  for(var a of mapa.values()){
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
  exports.consultar = listarPassagem;
