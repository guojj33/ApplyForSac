import Vue from 'vue'
import ElementUI from 'element-ui'
//import 'element-ui/lib/theme-chalk/index.css'
import router from './router'
import App from './App.vue'

import axios from 'axios'
import VueAxios from 'vue-axios'

Vue.use(VueAxios, axios)
Vue.use(ElementUI)
Vue.config.productionTip = false

import moment from 'moment'//导入文件

Vue.prototype.$moment = moment;//赋值使用

// 添加请求拦截器
axios.interceptors.request.use(config => {
  // 在发送请求之前做些什么
  //判断是否存在token，如果存在将每个页面header都添加token
  console.log("请求头添加token:" + sessionStorage.getItem('Token'));
  if(sessionStorage.getItem('Token')){
    config.headers.Authorization = 'Bearer ' + sessionStorage.getItem('Token');
  }

  return config;
  }, error => {
  // 对请求错误做些什么
    return Promise.reject(error);
  });

// http response 拦截器
axios.interceptors.response.use(response => {
    return response;
  }, error => {
    if (error.response) {
    switch (error.response.status) {
      case 401:
        sessionStorage.setItem('Token','');
        sessionStorage.setItem('AccountTypeStr','');
        router.replace('/');
      }
    }
  return Promise.reject(error.response.data)
  });

//拦截器放在上面，解决了刷新页面时拦截器不运行的问题
new Vue({
  el: '#app',
  router,
  components: { App },
  render: h => h(App)
})