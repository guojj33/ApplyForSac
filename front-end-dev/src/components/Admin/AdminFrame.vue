<template>
    <div id="adminFrame">
      <el-container>
        <el-header>
          <el-menu router :default-active="defaultActive()" mode="horizontal">
            <el-menu-item index="/admin/occupation">
                <i class="el-icon-date"></i>
                <span slot="title">查看占用</span>
            </el-menu-item>
            <el-menu-item index="/admin/review">
                <i class="el-icon-notebook-2"></i>
                <span slot="title">进行审核</span>
            </el-menu-item>
            <el-menu-item index="/admin/check">
                <i class="el-icon-circle-check"></i>
                <span slot="title">进行签到</span>
            </el-menu-item>
            <el-menu-item index="/admin/room">
                <i class="el-icon-house"></i>
                <span slot="title">房间管理</span>
            </el-menu-item>
            <el-menu-item index="/admin/time">
                <i class="el-icon-time"></i>
                <span slot="title">时间调整</span>
            </el-menu-item>
            <el-menu-item index="/admin/accounts">
                <i class="el-icon-user"></i>
                <span slot="title">账户管理</span>
            </el-menu-item>
            <el-menu-item index="/admin/info">
                <i class="el-icon-user"></i>
                <span slot="title">登出/注册</span>
            </el-menu-item>
          </el-menu>
        </el-header>
        <el-main>
          <router-view></router-view>
        </el-main>
      </el-container>
    </div>
</template>

<script>
export default {
  name: 'AdminFrame',
  methods: {
    defaultActive() {
      return this.$route.path
    }
  }
}
</script>

<style>

</style>
