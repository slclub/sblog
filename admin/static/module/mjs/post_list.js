layui.define(['laytpl','form'],function(exports){ //提示：模块也可以依赖其它模块，如：layui.define('layer', callback);
    var laytpl = layui.laytpl;
    var obj = {
		url:"/static/module/html/post_list.js.html",
		content : "",
		target:{},
		entry : function(target){
			this.target = target
		},
        body:function(){
			self = this

			laytpl(self.content).render({}, function(html){
				self.content = html;
			});
			self.target.innerHTML = self.content;

            return self.content;
        },
		load:function(){

			var $ = layui.jquery, form = layui.form();

			layui.form().render(); //更新全部

			form.on('checkbox(allChoose)', function(data){
				var child = $(data.elem).parents('table').find('tbody input[type="checkbox"]');
					child.each(function(index, item){
						item.checked = data.elem.checked;
				    });
				form.render('checkbox');
			});

		}
    };


    //输出test接口
    exports('sblog_post_list', obj);
});
