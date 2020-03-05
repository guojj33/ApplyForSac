import Vue from 'vue'
import Router from 'vue-router' 

Vue.use(Router)

const router = new Router({
  //mode: "history",
  mode: "hash",
  routes: [
    {
      path: '/',
      component: resolve => require(['../components/Login'], resolve)
    },
    {
      path: '/user',
      component: resolve => require(['../components/User/UserFrame'], resolve),
      children: [
        {
          path: 'occupation',
          component: resolve => require(['../components/User/Occupation'], resolve),
        },
        {
          path: 'apply',
          component: resolve => require(['../components/User/Apply'], resolve),
        },
        {
          path: 'review',
          component: resolve => require(['../components/User/Review'], resolve),
        },
        {
          path: 'info',
          component: resolve => require(['../components/User/Info'], resolve),
        }
      ]
    },
    {
      path: '/admin',
      component: resolve => require(['../components/Admin/AdminFrame'], resolve),
      children: [
        {
          path: 'occupation',
          component: resolve => require(['../components/Admin/Occupation'], resolve),
        },
        {
          path: 'review',
          component: resolve => require(['../components/Admin/Review'], resolve),
        },
        {
          path: 'check',
          component: resolve => require(['../components/Admin/Check'], resolve),
        },
        {
          path: 'info',
          component: resolve => require(['../components/Admin/Info'], resolve),
        },
        {
          path: 'room',
          component: resolve => require(['../components/Admin/Room'], resolve),
        },
        {
          path: 'accounts',
          component: resolve => require(['../components/Admin/Accounts'], resolve),
        },
        {
          path: 'time',
          component: resolve => require(['../components/Admin/TimeManagement'], resolve),
        }
      ]
    }
  ]
})

// 导航守卫
// 使用 router.beforeEach 注册一个全局前置守卫，判断用户是否登陆
router.beforeEach((to, from, next) => {
  if (to.path === '/') {  //不让已登录的浏览器进入登陆界面
    let token = sessionStorage.getItem('Token')
    console.log("sesstoken:",token);
    if (token !== null && token !== '') {
      let AccountTypeStr = sessionStorage.getItem('AccountTypeStr');
      console.log("seeA:",AccountTypeStr);
      if (AccountTypeStr === 'users') {
        next('/user/apply');
      } else {
        next('/admin/occupation');
      }
    }
    next();
  } else {
    if (sessionStorage.getItem('Token') === null || sessionStorage.getItem('Token') === '') { //没有检测到 token ，则跳转到登陆界面
      next('/');
    } else {
      next();
    }
  }
});

export default router;