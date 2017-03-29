layui.define(['laytpl'],function(exports){ //提示：模块也可以依赖其它模块，如：layui.define('layer', callback);
    var laytpl = layui.laytpl;
    var obj = {
        content : "",
        target:{},
        from:{},
        entry : function(target,from){
          this.target = target;
          this.from = from;
          return this;
        },
        hello: function(str){
          alert('Hello '+ (str||'test'));
        },
        load:function(callback, load){
            self = this;
            $.ajax({
                url:self.from.url,
                async:true,
                success:function(data,req_status){
                    self.content = data;

                    laytpl(self.content).render({}, function(html){
                      self.content = html;
                    });
                    self.target.innerHTML = self.content;

                    if(self.valideFunc(callback)){
                      callback(self.target);
                    }
                    if(self.valideFunc(load)){
                      load();
                    }
                }

            });
            return self.content;
        },
        valideFunc:function(func){
          if(func && (typeof func == 'function'||typeof func == 'object')){
            return true
          }
          return false;
        }
    };



    //输出test接口
    exports('sblog_op_load', obj);
});
