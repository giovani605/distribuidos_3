// esse arquivo lida com os requests nas passagens
const express = require('express');
const router = express.Router();
const gerenciador = require('./GerenciadorPassagem.js');
// aqui eu vou linda com os get request
router.get('/', (req,res,next) => {
  console.log('Consegui' + cont);
  cont++;
  res.status(200).json({
    message: 'lidando o get'
  });
});


// aqui lido com os post resquest
router.post('/', (req,res,next) =>{
  res.status(200).json({
    message:'Lidando com post'

  });

});

module.exports = router;
