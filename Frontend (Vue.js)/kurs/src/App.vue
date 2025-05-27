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



