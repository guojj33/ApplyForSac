<template>
    <el-card :body-style="{ padding: '0px' }" class="infoCard">
    <el-form :model="adminInfo">
        <el-form-item label="用户名">
        <el-input v-model="adminInfo.adminId" disabled></el-input>
        </el-form-item>
        <el-form-item label="邮箱">
        <el-input v-model="adminInfo.email" disabled></el-input>
        </el-form-item>
        <el-form-item class="button">
            <el-button
                size="large"
                type="primary"
                @click="registerFormVisible = true">
                注册
            </el-button>
        </el-form-item>
            <el-dialog title="注册新管理员" :visible.sync="registerFormVisible">
                <el-form :model="registerForm">
                    <el-form-item label="用户名">
                    <el-input v-model="registerForm.adminId"></el-input>
                    </el-form-item>
                    <el-form-item label="密码">
                    <el-input v-model="registerForm.password" show-password></el-input>
                    </el-form-item>
                    <el-form-item label="邮箱">
                    <el-input v-model="registerForm.email"></el-input>
                    </el-form-item>
                </el-form>
                <div slot="footer" class="dialog-footer">
                    <el-button @click="registerFormVisible = false">取 消</el-button>
                    <el-button type="primary" @click="Register();registerFormVisible = false">注册</el-button>
                </div>
            </el-dialog>

        <el-form-item class="button">
            <el-button
                size="large"
                type="danger"
                @click="LogOut()">
                登出
            </el-button>
        </el-form-item>
    </el-form>
    </el-card>
</template>

<script>
import global_ from '../../Global'
import { Loading } from 'element-ui'

export default {
    name: 'Info',
    data() {
        return {
            adminInfo: {
                adminId: '',
                email: '',
            },
            registerForm: {
                adminId: '',
                password: '',
                email: '',
            },
            registerFormVisible: false,
        };
    },
    created() {
        this.GetInfo();
    },
    methods: {
        GetInfo() {
            let loadingInst = Loading.service();
            let self_ = this;
            this.axios.get("/api/" + sessionStorage.getItem('AccountTypeStr') + "/" + sessionStorage.getItem('Id'))
            .then(function(response) {
                console.log(response);
                if (response.status === 200) {
                    console.log("获取个人信息成功");
                    self_.adminInfo.adminId = response.data.AdminId;
                    self_.adminInfo.email = response.data.Email;
                } else {
                    console.log("获取个人信息失败")
                }
                loadingInst.close();
            })
            .catch(function(error) {
                alert(error);
            });
        },
        LogOut() {
            let loadingInst = Loading.service();
            let self_ = this;
            this.axios.post("/api/logout", {})
            .then(function(response) {
                console.log(response);
                if (response.status === 200) {
                    console.log("退出登录成功");
                    sessionStorage.setItem('Token','');
                    sessionStorage.setItem('AccountTypeStr','');
                    sessionStorage.setItem('Id','');
                    self_.$router.replace('/');
                }
                loadingInst.close();
            })
            .catch(function(error) {
                alert(error);
            })
        },
        Register() {
            if (this.registerForm.adminId === '' || this.registerForm.password === '' || this.registerForm.email === '') {
                this.$alert('请补全信息','注册');
                return;
            }
            let self_ = this;
            this.axios.post("/api/admins/" + sessionStorage.getItem("Id") + "/register", this.registerForm)
            .then(function(response) {
                console.log(response);
                if (response.status === 200) {
                    console.log("注册成功")
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
                self_.registerForm = {
                    adminId: '',
                    password: '',
                    email: '',
                }
            })
            .catch(function(error) {
                self_.$alert(error);
            });
        }
    }
}
</script>

<style>
.infoCard {
    margin: auto;
    padding: 20px;
    width: 300px;
}

.button {
    display: block;
    margin: 15px auto;
    width: 70px;
}
</style>