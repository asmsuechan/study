(function() {
"use strict";

require('pdfjs-dist');
PDFJS.workerSrc = "node_modules/pdfjs-dist/build/pdf.worker.js";

var currentPage = 1;
var pdfDoc;
var url = "https://www.nii.ac.jp/userdata/karuizawa/h23/111104_3rdlecueda.pdf";
PDFJS.getDocument(url).then(function(pdf){
  pdfDoc = pdf;
  renderPage(pdfDoc);
});

var renderPage = function (pdf){
  //console.log(pdf)
  //console.log(pdfDoc)
  pdf.getPage(currentPage).then(function(page){
    var scale = 1.0;
    var viewport = page.getViewport(scale);
    var canvas = document.getElementById('pdf-canvas');
    var ctx = canvas.getContext('2d');
    var renderContext = {
      canvasContext: ctx,
      viewport: viewport
    };
    canvas.height = viewport.height;
    canvas.width = viewport.width;

    document.getElementById('pdf-canvas');
    page.render(renderContext);
  });
}

var nextPage = function (){
  currentPage++;
  renderPage(pdfDoc);
}

var loopPage = function(){
  currentPage++
  if (currentPage > pdfDoc.numPages){
    currentPage -= pdfDoc.numPages
  }
  renderPage(pdfDoc)
}

exports.nextPage = nextPage
exports.loopPage = loopPage

//console.log(nextPage)
//setInterval("nextPage()", 5000)
//setTimeout("self.nextPage", 5000)

//document.getElementById("next").onclick = function() {
//  currentPage++;
//  renderPage(pdfDoc);
//}
})();
