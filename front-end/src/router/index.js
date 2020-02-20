import Vue from 'vue'
import Router from 'vue-router'

import UserFrame from '../components/User/UserFrame'
import UserOccupation from '../components/User/Occupation'
import UserApply from '../components/User/Apply'
import UserReview from '../components/User/Review'
import UserInfo from '../components/User/Info'

import AdminFrame from '../components/Admin/AdminFrame'
import AdminOccupation from '../components/Admin/Occupation'
import AdminReview from '../components/Admin/Review'
import AdminCheck from '../components/Admin/Check'
import AdminInfo from '../components/Admin/Info'

import Login from '../components/Login'

Vue.use(Router)

const router = new Router({
  mode: "history",
  routes: [
    {
      path: '/',
      component: Login
    },
    {
      path: '/user',
      component: UserFrame,
      children: [
        {
          path: 'occupation',
          component: UserOccupation,
        },
        {
          path: 'apply',
          component: UserApply,
        },
        {
          path: 'review',
          component: UserReview,
        },
        {
          path: 'info',
          component: UserInfo,
        }
      ]
    },
    {
      path: '/admin',
      component: AdminFrame,
      children: [
        {
          path: 'occupation',
          component: AdminOccupation,
        },
        {
          path: 'review',
          component: AdminReview,
        },
        {
          path: 'check',
          component: AdminCheck,
        },
        {
          path: 'info',
          component: AdminInfo,
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