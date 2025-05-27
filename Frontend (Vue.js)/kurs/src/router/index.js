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

const routes = [
  { path: '/tshirt', component: TshirtsPage },
  { path: '/scarfs', component: ScarfsPage },
  { path: '/balls', component: BallsPage },
  { path: '/boots', component: BootsPage },
  { path: '/registration', component: RegistrationPage,  meta: {hideSidebar: true,hideButtonBar: true}},
  { path: '/authorisation', component: AuthoristaionPage, meta: {hideSidebar: true,hideButtonBar: true }},
  { path: '/favourites', component: FavouritesPage },
  { path: '/basket', component: BasketPage },
  { path: '/order', component: OrderPage, meta: {hideSidebar: true,hideButtonBar: true}}
]

const router = createRouter({
  history: createWebHistory(),
  routes
})


export default router
