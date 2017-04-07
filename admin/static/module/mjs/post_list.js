layui.define(['laytpl','form','laypage','sblog_edit','sblog_op_load'],function(exports){ //提示：模块也可以依赖其它模块，如：layui.define('layer', callback);
    var laytpl = layui.laytpl,laypage = layui.laypage;
    var obj = {
        url:"/back/module/html/post_list.js.html",
        gurl:"/sadmin/post/find",
        content :"",
        data_list:[],
        target:{},
		findData:{},
        entry : function(target){
          this.target = target
        },
        body:function(sdata){
          self = this
		  var renderData = {
			"post_list":sdata||[]
		  }

		  var htmlEl = document.getElementById('sblog_post_list_content');
		  self.content = sblog_post_list_content_tpl.innerHTML;
		  //self.content = htmlEl.innerHTML
          laytpl(self.content).render(renderData, function(ehtml){
            htmlEl.innerHTML = ehtml;
          });
          //self.target.innerHTML = self.content;

          layui.form().render(); //更新全部
			return self.content;
        },
        page:function(){
            laypage({
                cont: 'sblog_post_list_page'
                ,pages: 100 //总页数
                ,groups: 5 //连续显示分页数
                ,jump:function(object,first){
					if (first){
						return this
					}
					obj.findData.page=object.curr
					obj.load({NotFulshPage:true});
                }
            });
        },
        flush:function(){
            var self =obj
            sblog.ajax({
                url:self.gurl,
				data:self.findData,
                type:"GET",
                success:function(data,request_status){
                    if (!(data.constructor == Array)) {
                        return
                    }
                    self.data_list = data;//JSON.parse(data);
					self.render(self.data_list);
					self.listen();
                }
            })
        },
		render:function(obj){
			this.body(obj)
		},
        load:function(param){
			param = param || {}
          var self = obj

          var $ = layui.jquery, form = layui.form();

          layui.form().render(); //更新全部

          form.on('checkbox(allChoose)', function(data){
            var child = $(data.elem).parents('table').find('tbody input[type="checkbox"]');
              child.each(function(index, item){
                item.checked = data.elem.checked;
                });
            form.render('checkbox');
          });
		  if (!param.NotFulshPage) {
			  //pagenation
			  self.page();
		  }
		  self.flush();
        },
		listen:function(){
		  //edit operation.
		  $(".sblog_post_edit_icon").on("click", function(){
		    var el = $(this);
			var id = el.data("val")
			layui.sblog_op_load.entry(sblog_body,layui.sblog_edit);
			layui.sblog_op_load.load(null,layui.sblog_edit.load)
			layui.sblog_edit.load({ID:id})
		  });
		  $(".sblog_post_del_icon").on("click", function(){
			var el = $(this)
			layui.layer.confirm("Are you sure want to delete the data?",{
				btn:['confirm','cancel'],
				shade:true
			},function(index){
				sblog.ajax({
					url:"/sadmin/post/delete",
					data:{ID:el.data("val")},
					success:function(res){
						obj.load();
					}
				});
				layui.layer.close(index);
			});
		  });
		  $(".sblog_post_list_top").on("click", function(){
			var el = $(this)
			var ptr = el.parent().parent();
			var id = ptr.data("id")
			sblog.ajax({
				url:"/sadmin/post/top",
				data:{ID:id},
				success:function(){
					obj.load();
				}
			});
		  });
		}
    };

    //输出接口
    exports('sblog_post_list', obj);
});
