var util = {
	_print:function(mode,msg){
		console.log('['+new Date().toLocaleString()+' '+mode+":] "+msg);
	},
	debugLog:function(msg){
		this._print('DEBUG',msg);
	},
	fatalLog:function(msg){
		this._print('FATAL',msg);
	},
	errorLog:function(msg){
		this._print('ERROR',msg);
	},
}
var weixin = {
	get:function(request,response){
		var pageUrl = decodeURIComponent(decodeURI(request.url)).substr(1);
		//打开pageUrl的跳转链接
		var subPage = require('webpage').create();
		subPage.viewportSize = { width: 800, height: 600 };
		subPage.paperSize = { width: 800 , height: 600, margin: '0px' };
		subPage.open(pageUrl,function(status){
			if (status != "success") {
				response({
					code:500,
					msg:'打开url '+pageUrl+" 失败",
					data:''
				});
				return;
			}
			var html = subPage.evaluate(function(){
				return $('html').html();
			});
			response({
				code:200,
				msg:'',
				data:html
			});
		});
	}
};
var network = {
    _server: null,
    init: function(port) {
        this._server = require('webserver').create();
        this._server.listen('127.0.0.1:' + port, function(request, response) {
			util.debugLog('receiver a request:'+ JSON.stringify(request));
			try{
				weixin.get(request, function(result) {
					util.debugLog('finish request '+JSON.stringify(result));
					if( result.code != 0 )
						util.errorLog('request error '+JSON.stringify(result));
					response.statusCode = result.code;
					response.write(result.data);
					response.close();
				});
			}catch(e){
				util.errorLog('request error '+e);
				var result = {
					code:1,
					msg:e,
					data:'',
				};
				response.statusCode = 500;
				response.write(JSON.stringify(result));
				response.close();
			}
        });
    }
};
var system = require('system');
if (system.args.length != 2) {
    util.errorLog('Usage: main.js networkport');
    phantom.exit(1);
}else{
	var networkPort = system.args[1];
	network.init(networkPort);
}
