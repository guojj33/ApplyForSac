<template>
    <el-card :body-style="{ padding: '0px' }" class="infoCard">
    <el-form :model="userInfo">
        <el-form-item label="用户名">
        <el-input v-model="userInfo.userId" disabled></el-input>
        </el-form-item>
        <el-form-item label="邮箱">
        <el-input v-model="userInfo.email" disabled></el-input>
        </el-form-item>
        <el-form-item class="buttons">
          <el-button
            size="large"
            type="danger"
            @click="LogOut()">
            退出登录
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
            userInfo: {
                userId: '',
                email: '',
            }
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
                    self_.userInfo.userId = response.data.UserId;
                    self_.userInfo.email = response.data.Email;
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

.buttons {
    margin: auto;
    width: 98px;
}
</style>