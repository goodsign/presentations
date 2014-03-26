(function(){
  
  trace = arbor.etc.trace
  objmerge = arbor.etc.objmerge
  objcopy = arbor.etc.objcopy

  var Main = function(elt){
    var dom = $(elt)

    sys = arbor.ParticleSystem({stiffness:50, repulsion:2000, gravity:true, precision: 0.5})
    sys.renderer = Renderer("#viewport") // our newly created renderer will have its .init() method called shortly by sys...
    
    var _ed = dom.find('#editor')
    var _code = dom.find('textarea')
    var _canvas = dom.find('#viewport').get(0)
    
    var that = {
      init:function(){
        $(window).resize(that.resize)
        that.resize()
        that.updateLayout()
        window.setTimeout(that.updateGraph, 500)
        return that
      },
      updateGraph:function(e){
        $.getJSON('graph.json', function(data) {
          console.debug(data)

          if (data.nodeCount > 1) {
            sys.graft(data)
          }
          if (data.finished) {
            $("#logo").attr("class", "logo_icon_stopped")
          } else {
            window.setTimeout(that.updateGraph, 500)
          }
        });
      },
      
      resize:function(){        
        var w = $(window).width()
        var x = w - _ed.width()
        that.updateLayout(x)
        sys.renderer.redraw()
      },
      
      updateLayout:function(){
        var w = dom.width()
        var h = $(window).height() - 150
        var canvW = w
        var canvH = h
        _canvas.width = canvW
        _canvas.height = canvH
        sys.screenSize(canvW, canvH)
        sys.screenPadding(80, 140, 80, 140)
      }
    }
    
    return that.init()    
  }


  $(document).ready(function(){
    var mcp = Main("#container")    
  })

  
})()