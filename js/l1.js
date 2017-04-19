/*
let costs = Array[n][3]
for(int i = n; i< n; i++){
  for(int j = 0; j < 3; j++){
    costs[i][j]
  }
}
*/
function getMin (costs) {
  var i = 0
  for(i = 0; i < 3; i++) {
    console.log(costs[i] + '|')
  }
}

var costs = [[1,2,3],[3,4,2],[1,2,1],[2,3,4],[1,4,7],[2,6,8],[4,5,3]]
getMin(costs)
//module.exports = getMin;
