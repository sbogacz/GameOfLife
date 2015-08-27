var app = angular.module('GameOfLifeApp', []);

app.config(['$httpProvider', function ($httpProvider) {

    $httpProvider.defaults.useXDomain = true;
    //$httpProvider.defaults.headers.common = {};

    console.log('logging out headers');
    console.log($httpProvider.defaults);
    console.log($httpProvider.defaults.headers.common);
    console.log($httpProvider.defaults.headers.post);
    console.log($httpProvider.defaults.headers.put);
    console.log($httpProvider.defaults.headers.patch);
    console.log('end logging out headers');

    $httpProvider.defaults.headers.common = {Accept: "application/json, text/plain, */*"};
    $httpProvider.defaults.headers.post = {"Content-Type": "application/json;charset=utf-8"};

    console.log('after: logging out headers');
    console.log($httpProvider.defaults.headers.common);
    console.log($httpProvider.defaults.headers.post);
    console.log($httpProvider.defaults.headers.put);
    console.log($httpProvider.defaults.headers.patch);
    console.log('after: end logging out headers');

}]);
