import { createRouter, createWebHistory } from 'vue-router'
import TshirtsPage from '/src/tshirt.vue'
import ScarfsPage from '/src/skarf.vue'
import BallsPage from '/src/balls.vue'
import BootsPage from '/src/boots.vue'
import RegistrationPage from '/src/registration.vue'
import AuthoristaionPage from '/src/authorisation.vue'
import FavouritesPage from '/src/favourites.vue'
import BasketPage from '/src/basket.vue'
import OrderPage from '/src/order.vue'
// импортируйте другие страницы

const routes = [
  { path: '/tshirt', component: TshirtsPage },
  { path: '/scarfs', component: ScarfsPage },
  { path: '/balls', component: BallsPage },
  { path: '/boots', component: BootsPage },
  { path: '/registration', component: RegistrationPage,  meta: {hideSidebar: true,hideButtonBar: true}},
  { path: '/authorisation', component: AuthoristaionPage, meta: {hideSidebar: true,hideButtonBar: true }},
  { path: '/favourites', component: FavouritesPage },
  { path: '/basket', component: BasketPage },
  { path: '/order', component: OrderPage }
  // другие маршруты
]

const router = createRouter({
  history: createWebHistory(),
  routes
})


export default router
/*
import { createRouter, createWebHistory } from 'vue-router'
import TshirtsPage from '/src/tshirt.vue'
import ScarfsPage from '/src/skarf.vue'
import BallsPage from '/src/balls.vue'
import BootsPage from '/src/boots.vue'
import RegistrationPage from '/src/registration.vue'
import AuthorisationPage from '/src/authorisation.vue'
import FavouritesPage from '/src/favourites.vue'
import BasketPage from '/src/basket.vue'
import HomePage from '/src/App.vue' // Импортируем главную страницу
import store from '../store' // Импортируем хранилище Vuex

const routes = [
  {
    path: '/',
    component: HomePage,
    meta: { requiresAuth: false }
  },
  { 
    path: '/tshirt', 
    component: TshirtsPage,
    meta: { requiresAuth: false }
  },
  { 
    path: '/scarfs', 
    component: ScarfsPage,
    meta: { requiresAuth: false }
  },
  { 
    path: '/balls', 
    component: BallsPage,
    meta: { requiresAuth: false }
  },
  { 
    path: '/boots', 
    component: BootsPage,
    meta: { requiresAuth: false }
  },
  { 
    path: '/registration', 
    component: RegistrationPage,
    meta: {
      hideSidebar: true,
      hideButtonBar: true,
      requiresAuth: false
    }
  },
  { 
    path: '/authorisation',  // Исправлена опечатка в названии
    component: AuthorisationPage,
    meta: {
      hideSidebar: true,
      hideButtonBar: true,
      requiresAuth: false
    }
  },
  { 
    path: '/favourites', 
    component: FavouritesPage,
    meta: { requiresAuth: true } // Только для авторизованных
  },
  { 
    path: '/basket', 
    component: BasketPage,
    meta: { requiresAuth: true } // Только для авторизованных
  },
  // Дополнительный маршрут для перенаправления
  {
    path: '/:pathMatch(.*)*',
    redirect: '/'
  }
]

const router = createRouter({
  history: createWebHistory(),
  routes
})

// Глобальный навигационный хук
router.beforeEach((to, from, next) => {
  // Проверяем, требует ли маршрут аутентификации
  if (to.meta.requiresAuth && !store.state.isAuthenticated) {
    // Перенаправляем на страницу авторизации
    next('/authorisation')
  } else if ((to.path === '/authorisation' || to.path === '/registration') && store.state.isAuthenticated) {
    // Если пользователь уже авторизован, перенаправляем с этих страниц
    next('/')
  } else {
    // Продолжаем навигацию
    next()
  }
})

export default router*/