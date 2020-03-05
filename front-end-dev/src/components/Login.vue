<template>
<el-container>
    <el-main class="backGround">
        <el-card :body-style="{ padding: '0px' }" class="loginCard">
            <div class="title">
            <embed src="/static/logo.svg" type="image/svg+xml" width="50px" height="50px"/>
            </div>
            <el-form :model="loginForm">
                <el-form-item label="用户名">
                <el-input v-model="loginForm.Id"></el-input>
                </el-form-item>
                <el-form-item label="密码">
                <el-input v-model="loginForm.Password" show-password></el-input>
                </el-form-item>
                <el-form-item class="button">
                    <el-button
                        size="large"
                        type="danger"
                        @click="registerFormVisible = true">
                        注册
                    </el-button>
                    <el-dialog title="注册新用户" :visible.sync="registerFormVisible">
                        <el-form :model="registerForm">
                            <el-form-item label="用户名">
                            <el-input v-model="registerForm.UserId"></el-input>
                            </el-form-item>
                            <el-form-item label="密码">
                            <el-input v-model="registerForm.Password" show-password></el-input>
                            </el-form-item>
                            <el-form-item label="邮箱">
                            <el-input v-model="registerForm.Email"></el-input>
                            </el-form-item>
                        </el-form>
                        <div slot="footer" class="dialog-footer">
                            <el-button @click="registerFormVisible = false">取 消</el-button>
                            <el-button type="primary" @click="Register();registerFormVisible = false">注册</el-button>
                        </div>
                    </el-dialog>
                </el-form-item>
                <el-form-item class="button">
                    <el-button
                        size="large"
                        type="primary"
                        @click="Login()">
                        登录
                    </el-button>
                </el-form-item>
            </el-form>
        </el-card>
        <el-card class="commentCard">
            <div v-for="(coment, index) in comments" :key="index">
                <div class="cName">{{coment.name}}</div>
                <div class="cContent">{{coment.content}}</div>
                <el-divider></el-divider>
            </div>
            <div class="addCommentBox">
                <el-form v-model="addCommentForm">
                    <el-form-item>
                        <el-input style="width: 110px;" v-model="addCommentForm.name" placeholder="随便的昵称">
                        </el-input>
                    </el-form-item>
                    <el-form-item>
                        <el-input type="textarea"  :autosize="{minRows: 2}" v-model="addCommentForm.content" placeholder="随便的内容...">
                        </el-input>
                    </el-form-item>
                    <el-form-item>
                        <el-button
                        size="large"
                        type="primary"
                        @click="AddComment()">
                        发表
                        </el-button>
                    </el-form-item>
                </el-form>
            </div>
        </el-card>
    </el-main>
    <el-footer>
        <div class="footer">
        <a target="_blank" href="https://github.com/guojj33/ApplyForSac">Apply For Sac</a>
        &nbsp;by&nbsp;
        <a target="_blank" href="https://github.com/guojj33">guojj33</a>
        </div>
    </el-footer>
</el-container>
</template>

<script>
import global_ from '../Global'
import { Loading } from 'element-ui'

export default {
    name: 'Login',
    data() {
        return {
            addCommentForm: {
                name: '',
                content: '',
            },
            comments: [{
                name: 'Admin',
                content: 'Hi',
            },{
                name: 'JJ',
                content: 'Hello',
            }],
            loginForm: {
                Id: '',
                Password: '',
            },
            registerForm: {
                UserId: '',
                Password: '',
                Email: '',
            },
            registerFormVisible: false,
        };
    },
    created() {
        this.registerFormVisible = false;
        this.GetComments();
    },
    methods: {
        GetComments() {
            var createCommentRow = function() {
                var commentRow = new Object ({
                    name: '',
                    content: '',
                })
                return commentRow;
            }
            let loadingInst = Loading.service();
            let self_ = this;
            let commentRows = [];
            this.axios.get("/api/comments")
            .then(function(response) {
                console.log(response)
                if (response.status === 200) {
                    console.log("获取评论列表成功");
                    let commentsData = response.data.comments;
                    for (let index in commentsData) {
                        let c = commentsData[index];
                        let commentRow = createCommentRow()
                        commentRow.name = c.Name
                        commentRow.content = c.Content
                        commentRows.push(commentRow);
                    }
                    self_.comments = commentRows;
                } else {
                    console.log("获取评论列表失败")
                }
                loadingInst.close();
            })
            .catch(function(error) {
                self_.$alert(error)
            })
        },
        AddComment() {
            if (this.addCommentForm.name === '' || this.addCommentForm.content === '') {
                this.$alert('请补全评论信息', '评论');
                return;
            }
            let self_ = this;
            this.axios.post("api/comments", this.addCommentForm)
            .then(function(response) {
                if (response.status === 200) {
                    self_.$alert('发表成功','评论');
                    self_.GetComments();
                } else {
                    self_.$alert('发表失败','评论');
                }
                self_.addCommentForm = {
                    name: '',
                    content: '',
                };
            })
            .catch(function(error) {
                self_.$alert(error);
            })
        },
        Login() {
            var self_ = this
            if (self_.loginForm.Id === '' || self_.loginForm.Password === '') {
                //loadingInst.close();
                self_.$alert('用户名或密码不能为空','登录',{
                    confirmButtonText: "确认"
                });
                return;
            }
            let loadingInst = Loading.service();
            this.axios.post("/api/login", this.loginForm)
            .then(function (response) {
                console.log(response);
                if (response.status === 200) {
                    console.log("登录成功");
                    var responseData = response.data;
                    console.log(responseData);
                    // //保存 token
                    sessionStorage.setItem('Token',responseData.Token.SAC_TOKEN);
                    sessionStorage.setItem('Id',self_.loginForm.Id);
                    switch (responseData.CurAccountType) {
                        case 0:
                            self_.$router.replace('/user/apply');
                            console.log("用户登录");
                            sessionStorage.setItem('AccountTypeStr','users');
                            break;
                        case 1:
                            self_.$router.replace('/admin/review');
                            console.log("管理员登陆");
                            sessionStorage.setItem('AccountTypeStr','admins');
                            break;
                    }
                } else {
                    self_.$alert('登陆失败','登录',{
                        confirmButtonText: "确认"
                    });
                    console.log("登录失败");
                }
                loadingInst.close();
            })
            .catch(function (error) {
                self_.$alert(error,'系统',{
                    confirmButtonText: "确认"
                });
                loadingInst.close();
            });
        },
        Register() {
            if (this.registerForm.UserId === '' || this.registerForm.Password === ''
            || this.registerForm.error) {
                this.$alert('请补全信息','注册');
                return;
            }
            let self_ = this;
            this.axios.post("/api/register", this.registerForm)
            .then(function (response) {
                console.log(response);
                if (response.status === 200) {
                    console.log("注册成功");
                    var responseData = response.data;
                    console.log(responseData);
                    self_.$alert('注册成功','注册',{
                        confirmButtonText: "确认"
                    });
                } else {
                    console.log("注册失败")
                    self_.$alert('注册失败','注册',{
                        confirmButtonText: "确认"
                    });
                }
            })
            .catch(function (error) {
                self_.$alert(error);
            });
        }
    }
}
</script>

<style>
.title {
    margin: auto;
    width: 30px;
}

.loginCard {
    margin: 20px auto;
    padding: 20px;
    width: 300px;
}

.backGround {
    margin: auto;
}

.button {
    width: 70px;
    margin: 15px auto;
}

.footer {
    text-align: center;
    color: #909399;
    font-size: 13px;
}

a:link {
    color: #909399;
}

a:visited {
    color: #909399;
}

.commentCard {
    margin: auto;
    padding: 10px;
    width: 320px;
}

.cName {
    color: #fb7299;
    font-weight: bold;
    font-size: 13px;
    margin-bottom: 10px;
}

.cContent {
    font-size: 16px;
}
</style>