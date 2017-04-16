
layui.define(['laytpl','form', 'layedit'],function(exports){ //提示：模块也可以依赖其它模块，如：layui.define('layer', callback);
    var laytpl = layui.laytpl;
    var form = layui.form()
	//创建一个编辑器
    var obj = {
        url:"/back/module/html/edit.js.html",
        content : "",
        target:{},
        entry : function(target){
          this.target = target
        },
        hello: function(str){
          alert('Hello '+ (str||'test'));
        },
        render:function(id){
            self = this
            sblog.ajax({
                url:"/sadmin/post/addhtml",
                data:{ID:id},
                async:true,
                success:function(data,req_status){
                    form.render();
                    data = data||{};
                    data.ID = data.ID || $("#ID").value;
                    data.submit_val = (data.ID) ? "Edit":"Add";
                    data.title = data.title || "";
                    data.tags = data.tags || "";

                    data.edit_content = data.content || "";
                    delete(data.content)

                    //self.content=sblog_post_add_tpl.innerHTML;
                    self.content=sblog_post_add_html.innerHTML;
                    laytpl(self.content).render(data, function(html){
                      self.content = html;
                    });
                    sblog_post_add_html.innerHTML = self.content;

                    var editIndex = layui.layedit.build('LAY_demo_editor');
                    //自定义验证规则
                    form.verify({
                      title: function(value){
                        if(value.length < 3){
                        return '标题至少得3个字符啊';
                        }
                      }
                      ,content: function(value){
                        layui.layedit.sync(editIndex);
                      }
                    });

                }

            });
            return self.content;
        },
        load:function(param){
            param = param ||{}
            self = obj;
            form.render();
            self.render(param.ID||0)
            //监听提交
            form.on('submit(post_edit_submit)', function(data){
                self.save(data.field);
                return false;
            });

        },
        save:function(adata){
            self = obj;
            adata = adata || {};
            console.log(adata)
            sblog.ajax({
                url:"/sadmin/post/save",
                data:adata,
                type:"POST",
                success:function(res,res_status){
                    layer.alert("OK success!!!"+JSON.stringify(res||{}),{
                        title:"after save"
                    })
                    //self.render(res.ID);
                }
            });
        }
    };



    //输出test接口
    exports('sblog_edit', obj);
});
