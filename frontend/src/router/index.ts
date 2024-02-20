import { createRouter, createWebHistory } from 'vue-router'
import TokenServer from "@/token";

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/login',
      name: 'login',
      component: () => import('@/views/login/login.vue')
    },
    {
      path: '/layout',
      name: 'layout',
      component: () => import('@/views/layout/layout.vue')
    }
  ]
})

router.beforeEach((to,from,next)=>{
  if (to.path === '/login'){
    next()
  }else if (!TokenServer.getToken()){
    next('/login')
  }else {
    next()
  }
})

export default router
