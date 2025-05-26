<!--<script setup>

</script>

<template>
<header class = "flex justify-between">
  <div>
    <p>Регистрация</p>
    <p>Авторизация</p>
  </div>
  <div class="flex items-center">
    <img class ="max-w-[8rem] h-auto object-contain" src="/ball.png" alt="logo"/>
    <div>
      <h1 class="title-logo">Footballiks</h1>
    </div>
  </div>
  <div>
    <p>Избранное</p>
    <p>Корзина</p>
  </div>
</header>
</template>

<style scoped>

</style>-->



<template>
  <div id="app">
    <Header />
    <UserName />
    <ButtonBar v-if="!$route.meta.hideButtonBar" />
    <Sidebar v-if="!$route.meta.hideSidebar" />
    <main class="main-content">
      <router-view></router-view>
    </main>
    <Footer />
  </div>
</template>

<script>
import Header from './components/Header.vue'
import UserName from './components/UserName.vue'
import ButtonBar from './components/ButtonBar.vue'
import Sidebar from './components/Sidebar.vue'
import Footer from './components/Footer.vue'


export default {
  name: 'App',
  components: {
    Header,
    UserName,
    ButtonBar,
    Sidebar,
    Footer
  },
}
</script>

<style scoped>
html, body {
  height: 100%;
  margin: 0;
}

#app {
  min-height: 100vh;
  display: flex;
  flex-direction: column;
}

.main-content {
  flex: 1;
}
</style>



