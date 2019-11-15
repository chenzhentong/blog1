{{define  "nav"}}
    <style>
        .navbar-default {
            border: none;
            box-shadow: 0px 2px 2px 0px rgba(50,50,50,0.25);
            font-weight: bold;
            border:none;
        }
        .navbar-default .navbar-brand{
            height:60px;
            font-size: 25px;
            line-height: 30px;
        }
        .navbar-default .navbar-collapse a
        {
            height:60px;
            font-size: 15px;
            line-height: 30px;
            color:#202020 ;
        }
        .navbar-default .navbar-nav>li>a {
            color: #666;
        }
        body{
            background-color: #F6F6F6;
        }
        .nav> li:hover .dropdown-menu {display: block;}
    </style>
    <nav class="navbar navbar-default " role="navigation">
        <div class="container">
            <div class="navbar-header">
                <button class="navbar-toggle" data-toggle="collapse" data-target=".navbar-collapse ">
                    <span class="icon-bar"></span>
                    <span class="icon-bar"></span>
                    <span class="icon-bar"></span>

                </button>
                <a href="/" class="navbar-brand" title="admin">admin</a>

            </div>
            <div class="navbar-collapse collapse">
                <ul class="nav navbar-nav ">
                    <li {{if .isHome}} class="active" {{end}}><a href="javascript:void(0)" onclick="GoHome()"><span class="glyphicon glyphicon-home"></span>首页</a></li>
                    <li {{if .isCategory}}class="active"{{end}}><a href="/category"> <i class="fa fa-tags"></i>分类管理</a></li>
                    <li {{if .isTopic}}class="active"{{end}}><a href="/topic/list"><span class="glyphicon glyphicon-list-alt"></span>文章管理</a></li>
                    <li class="dropdown" >
                        <a type="button" data-toggle="dropdown"   id="dLabel">
                            Java<span class="caret"></span>
                        </a>

                        <!--<li ><a href="#admin" title="admin" name="admin">admin</a></li>-->

                        <ul class="dropdown-menu" arial-labelledby="dLabel">
                            <li><a href="index.html">Struct2</a></li>
                            <li><a href="index.html">Spring</a></li>
                            <li><a href="index.html">Hibernate</a></li>
                            <li><a href="index.html">Mybatis</a></li>
                            <li><a href="index.html">SpringBoot</a></li>
                        </ul>
                    </li>
                </ul>
                <ul class="nav navbar-nav navbar-right">
                    <li class="dropdown" >
                        <a type="button" data-toggle="dropdown"  aria-expanded="false" id="dLabel">
                            admin<span class="caret"></span>
                        </a>
                        <ul class="dropdown-menu" arial-labelledby="dLabel">
                            <li><a href="index.html">前台首页</a></li>
                            <li><a href="index.html">个人主页</a></li>
                            <li><a href="index.html">个人设置</a></li>
                            <li><a href="index.html">账户中心</a></li>
                            <li><a href="index.html">我的收藏</a></li>
                        </ul>

                        <!--<li ><a href="#admin" title="admin" name="admin">admin</a></li>-->

                    </li>
                    <li>
                        {{if .isLogin}}
                            <a  href="/login/exit"><span class="glyphicon glyphicon-off">退出</a>

                        {{else}}
                            <a href="javascript:void(0)"  data-toggle="modal"
                               data-target="#login"  id="notLogin">登录</a>
                        {{end}}

                        <div id="login" class="modal fade">
                            <div class="modal-dialog">
                                <div class="modal-content">
                                    <div class="modal-body">
                                        <button class="close" data-dismiss="modal">
                                            <span>&times;</span>
                                        </button>
                                    </div>
                                    <div class="modal-title">
                                        <h1 class="text-center">登录</h1>
                                    </div>

                                    <form class="form-group"
                                          action="/login/go"
                                          method="post" onsubmit="return checkInputInfo()">
                                        <div class="form-group">
                                            <label for="">用户名</label> <input class="form-control"
                                                                             type="text" placeholder="" name="Username" id="Username">
                                        </div>
                                        <div class="form-group">
                                            <label for="">密码</label> <input class="form-control" type="password" placeholder="" name="Password" id="Password">
                                        </div>
                                        <div class="form-group">
                                            <input type="checkbox" placeholder="" name="AutoLogin"> 自动登录
                                        </div>
                                        <div class="text-right">
                                            <button class="btn btn-primary" type="submit">登录</button>
                                            <button class="btn btn-danger" data-dismiss="modal">取消</button>

                                        </div>
                                    </form>
                                </div>
                            </div>
                        </div>

                    </li>
                </ul>
            </div>
        </div>
    </nav>
    <script type="text/javascript">
        //检验用户名密码是否为空
        function checkInputInfo() {
            "".trim()
            if($("#Username").val()==""){
                alert("请输入用户名!")
                return false

            }else if( $("#Password").val()==""){
                alert("请输入密码!")
                return false
            }
            return true
        }
        function  GoHome() {
            location.href="/"

        }

    </script>
{{end}}