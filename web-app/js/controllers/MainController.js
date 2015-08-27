app.controller('MainController', ['$scope','gameOfLife', function ($scope, gameOfLife){
   /*gameOfLife.getBoard().success(function(data) {
		$scope.board = data;
		console.log(data)
   })*/
	var gameId = 0;

	gameOfLife.newBoard().then(function(response){
		gameId = response.data;
		gameOfLife.getBoard(gameId).then(function(response) {
			$scope.board = response.data;
			console.log($scope.board);
		});
	});
	

	$scope.stepGame = function() {
		gameOfLife.updateGame($scope.board, gameId).then(function(response){
			gameOfLife.stepBoard(gameId).then(function(response){
				$scope.board = response.data;
			});
			//return;
		});
		//gameOfLife.stepBoard(gameId).then(function(response){
		//	$scope.board = response.data;
		//});
	}
	$scope.update = function(rowIdx, colIdx, val) {
		$scope.board.grid[rowIdx][colIdx] = !val;
	}
	
}]);
