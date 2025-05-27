<template>
  <div>
    <h1>Избранное</h1>
    <div v-if="isAuthenticated">
      <div v-if="favorites.length > 0">
        <ProductGrid :products="favorites" :is-favorite-page="true" @remove-from-favorites="removeFromFavorites"/>
      </div>
      <div v-else>
        <p>В избранном пока ничего нет</p>
      </div>
    </div>
    <div v-else>
      <p>Авторизуйтесь, чтобы просматривать избранное</p>
    </div>
  </div>
</template>


<script>
import ProductGrid from './components/ProductGrid.vue'
import { mapState, mapActions } from 'vuex';

export default {
  name: 'FavouritesPage',
  components: {
    ProductGrid
  },
  computed: {
    ...mapState(['isAuthenticated', 'favorites'])
  },
  async created() {
    if (this.isAuthenticated) {
      await this.loadFavorites();
    }
  },
  methods: {
    ...mapActions(['removeFromFavorites']),
    async loadFavorites() {
      try {
        await this.$store.dispatch('fetchFavorites');
      } catch (error) {
        console.error('Ошибка загрузки избранного:', error);
        alert('Не удалось загрузить избранное');
      }
    },
    async handleRemoveFavorite(productId) {
    try {
      console.log('Removing product:', productId);
      await this.$store.dispatch('removeFromFavorites', productId);
      await this.$store.dispatch('fetchFavorites');
    } catch (error) {
      console.error('Remove failed:', error);
      this.$notify({
        title: 'Ошибка',
        message: error.message,
        type: 'error',
        duration: 5000
      });
    }
  }
  }
}
</script>