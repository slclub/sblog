layui.define(['laytpl','form','laypage'],function(exports){ //提示：模块也可以依赖其它模块，如：layui.define('layer', callback);
    var laytpl = layui.laytpl,laypage = layui.laypage;
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
        page:function(){
            laypage({
                cont: 'sblog_post_list_page'
                ,pages: 100 //总页数
                ,groups: 5 //连续显示分页数
                ,jump:function(){
                }
            });
        },
        load:function(){
          self = obj

          var $ = layui.jquery, form = layui.form();

          layui.form().render(); //更新全部

          form.on('checkbox(allChoose)', function(data){
            var child = $(data.elem).parents('table').find('tbody input[type="checkbox"]');
              child.each(function(index, item){
                item.checked = data.elem.checked;
                });
            form.render('checkbox');
          });
          //pagenation
          self.page()
        }
    };

    //输出接口
    exports('sblog_post_list', obj);
});
