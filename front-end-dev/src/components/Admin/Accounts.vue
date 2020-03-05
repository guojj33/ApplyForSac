<template>
<el-container>
    <el-main>
    <el-card :body-style="{ padding: '0px' }" class="usersCard accountCard">
        <el-table
        v-loading="isLoading"
        :data="users"
        style="width: 100%">
            <el-table-column
            prop="userId"
            label="用户名"
            width="150px">
            </el-table-column>
            <el-table-column
            prop="email"
            label="邮箱"
            width="150px">
            </el-table-column>
        </el-table>
    </el-card>
    <el-card :body-style="{ padding: '0px' }" class="adminsCard accountCard">
        <el-table
        v-loading="isLoading"
        :data="admins"
        style="width: 100%">
            <el-table-column
            prop="adminId"
            label="管理员名"
            width="150px">
            </el-table-column>
            <el-table-column
            prop="email"
            label="邮箱"
            width="150px">
            </el-table-column>
        </el-table>
    </el-card>
    </el-main>
</el-container>
</template>

<script>
export default {
    name: 'Accounts',
    data() {
        return {
            isLoading: true,
            users: [],
            admins: [],
        };
    },
    created() {
        this.GetAccounts();
    },
    methods: {
        GetAccounts() {
            var createUserRow = function() {
                var userRow = new Object ({
                    userId: '',
                    email: '',
                });
                return userRow;
            };
            var createAdminRow = function() {
                var adminRow = new Object ({
                    adminId: '',
                    email: '',
                });
                return adminRow;
            };
            let self_ = this;
            let userRows = [];
            let adminRows = [];
            this.axios.get("/api/users")
            .then(function(response) {
                console.log(response);
                if (response.status === 200) {
                    console.log("获取用户列表成功");
                    let usersData = response.data.users;
                    for (let index in usersData) {
                        let u = usersData[index];
                        //console.log(u)
                        let userRow = createUserRow();
                        userRow.userId = u.UserId;
                        userRow.email = u.Email;
                        userRows.push(userRow);
                    }
                    self_.users = userRows
                    self_.axios.get("/api/admins")
                    .then(function(response) {
                        console.log(response);
                        if (response.status === 200) {
                            console.log("获取管理员列表成功")
                            let adminsData = response.data.admins;
                            for (let index in adminsData) {
                                let a = adminsData[index];
                                let adminRow = createAdminRow();
                                adminRow.adminId = a.AdminId;
                                adminRow.email = a.Email;
                                adminRows.push(adminRow);
                            }
                            self_.admins = adminRows;
                            self_.isLoading = false;
                        }
                    })
                    .catch(function(error) {
                        self_.$alert(error);
                    });
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
.accountCard {
    width: 300px;
    padding: 20px;
    margin: 20px auto;
}
</style>