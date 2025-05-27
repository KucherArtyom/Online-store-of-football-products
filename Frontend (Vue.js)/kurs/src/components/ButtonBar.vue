<template>
  <div class="button-bar">
    <button variant="success" v-if="!isAuthenticated" type="submit" formmethod="get" formaction="registration.html" class="registration-button">
      <router-link to="/registration">
        Регистрация
      </router-link>
    </button>
    <button type="submit" class="authorisation-button" @click="handleAuthButton">
      {{ isAuthenticated ? 'Выйти из акк.' : 'Авторизация' }}
    </button>
    <button type="submit" formmethod="get" formaction="favourites.html" class="favourites-button">
      <router-link to="/favourites">
        Избранное
      </router-link>
    </button>
    <button type="submit" formmethod="get" formaction="basket.html" class="basket-button">
      <router-link to="/basket">
        Корзина
      </router-link>
    </button>
  </div>
</template>

<script>
import { mapState, mapMutations } from 'vuex';

export default {
  name: 'ButtonBar',
  computed: {
    ...mapState(['isAuthenticated'])
  },
  methods: {
    ...mapMutations(['logout']),
    handleAuthButton() {
      if (this.isAuthenticated) {
        this.logout();
        this.$router.push('/');
      } else {
        this.$router.push('/authorisation');
      }
    }
  }
}
</script>



