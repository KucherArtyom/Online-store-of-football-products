<template>
  <div>
    <h1>Корзина</h1>

    <div v-if="isAuthenticated">
      
      <!--<div v-if="basket.length > 0">-->
      <div v-if="(basket ?? []).length > 0">
        <ProductGrid 
          :products="basket" 
          :is-basket-page="true"
          @remove-from-basket="handleRemoveFromBasket"
        />
        <div class="checkout-section">
          <p class="total-price">Общая сумма: {{ totalPrice }} руб.</p>
          <button @click="proceedToCheckout" class="checkout-button">
            Перейти к оформлению заказа
          </button>
        </div>
      </div>
      <div v-else>
        <p>В корзине пока ничего нет</p>
      </div>
    </div>
    <div v-else>
      <p>Авторизуйтесь, чтобы просматривать корзину</p>
    </div>
  </div>
</template>

<script>
import ProductGrid from './components/ProductGrid.vue'
import { mapState, mapActions } from 'vuex'

export default {
  name: 'BasketPage',
  components: {
    ProductGrid
  },
  computed: {
    ...mapState(['isAuthenticated', 'basket']),
    totalPrice() {
      return this.basket?.reduce((sum, product) => sum + product.price, 0) || 0
    }
  },
  methods: {
    ...mapActions(['removeFromBasket', 'fetchBasket']),
    proceedToCheckout() {
      this.$router.push('/order')
    },
    async loadBasket() {
      try {
        await this.fetchBasket()
      } catch (error) {
        console.error('Ошибка загрузки корзины:', error)
      }
    },
    async handleRemoveFromBasket(productId) {
      try {
        await this.removeFromBasket(productId)
      } catch (error) {
        console.error('Ошибка удаления:', error)
      }
    }
  },
  async created() {
    if (this.isAuthenticated) {
      await this.loadBasket()
    }
  }
}
</script>

<style scoped>
.checkout-section {
  margin-top: 20px;
  padding: 20px;
  background-color: #f5f5f5;
  border-radius: 5px;
  text-align: right;
}

.total-price {
  font-size: 18px;
  font-weight: bold;
  margin-bottom: 15px;
}

.checkout-button {
  background-color: #94BF4C;
  color: white;
  padding: 10px 20px;
  border: none;
  border-radius: 4px;
  cursor: pointer;
  font-size: 16px;
}

.checkout-button:hover {
  background-color: #7aa33a;
}
</style>