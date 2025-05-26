<!--<template>
    <p class="userauth">Вы не авторизованы</p>
</template>

<script>
export default {
  name: 'Username'
}
</script>

<style scoped>
.userauth{
    text-align: center;
}
</style>-->

<template>
  <p class="userauth">
    {{ isAuthenticated ? `Добро пожаловать, ${userName}` : 'Вы не авторизованы' }}
  </p>
</template>

<script>
import { mapState } from 'vuex';

export default {
  name: 'Username',
  computed: {
    ...mapState(['isAuthenticated', 'userName'])
  }
}
</script>

<style scoped>
.userauth {
  text-align: center;
}
</style>