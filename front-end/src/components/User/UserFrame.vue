<template>
    <div id="userFrame">
      <el-container>
        <el-header>
          <el-menu router :default-active="defaultActive()" mode="horizontal">
            <el-menu-item index="/user/occupation">
                <i class="el-icon-date"></i>
                <span slot="title">查看占用情况</span>
            </el-menu-item>
            <el-menu-item index="/user/apply">
                <i class="el-icon-document"></i>
                <span slot="title">创建房间申请</span>
            </el-menu-item>
            <el-menu-item index="/user/review">
                <i class="el-icon-notebook-2"></i>
                <span slot="title">查看审核情况</span>
            </el-menu-item>
            <el-menu-item index="/user/info">
                <i class="el-icon-user"></i>
                <span slot="title">个人信息</span>
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
  name: 'UserFrame',
  methods: {
    defaultActive() {
      return this.$route.path
    }
  }
}
</script>

<style>

</style>
