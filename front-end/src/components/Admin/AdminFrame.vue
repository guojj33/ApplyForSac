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
            <el-menu-item index="/admin/info">
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
