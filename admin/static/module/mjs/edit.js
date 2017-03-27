
layui.define(['laytpl'],function(exports){ //提示：模块也可以依赖其它模块，如：layui.define('layer', callback);
    var laytpl = layui.laytpl;
    var obj = {
		url:"/static/module/html/edit.js.html",
		content : "",
		target:{},
		entry : function(target){
			this.target = target
		},
        hello: function(str){
          alert('Hello '+ (str||'test'));
        },
        body:function(){
			self = this
            $.ajax({
                url:"/static/module/html/edit.js.html",
                async:true,
                success:function(data,req_status){

                    self.content = data;

					laytpl(self.content).render({}, function(html){
					    self.content = html;
				    });
					self.target.innerHTML = self.content;
                }

            });
            return self.content;
        }
    };



    //输出test接口
    exports('sblog_edit', obj);
});
