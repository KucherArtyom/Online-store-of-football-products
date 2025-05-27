<template>
  <div>
    <div v-for="(row, index) in chunkedProducts" :key="index" class="product-row">
      <ProductItem 
        v-for="product in row" 
        :key="product.id"
        :product="product"
        :is-favorite-page="isFavoritePage"
        :is-basket-page="isBasketPage"
        @remove-from-favorites="handleRemoveFavorite"
        @remove-from-basket="handleRemoveFromBasket"
      />
    </div>
  </div>
</template>

<script>
import ProductItem from './ProductItem.vue'

export default {
  name: 'ProductGrid',
  components: {
    ProductItem
  },
  props: {
    products: {
      type: Array,
      required: true
    },
    isFavoritePage: {
      type: Boolean,
      default: false
    },
    isBasketPage: {
      type: Boolean,
      default: false
    }
  },
  computed: {
    chunkedProducts() {
      const chunkSize = 3
      const result = []
      for (let i = 0; i < this.products.length; i += chunkSize) {
        result.push(this.products.slice(i, i + chunkSize))
      }
      return result
    }
  },
  
  methods: {
    handleRemoveFavorite(productId) {
      this.$emit('remove-from-favorites', productId);
    },
    handleRemoveFromBasket(productId) {
      this.$emit('remove-from-basket', productId);
    }
  }
}
</script>

<style scoped>
.product-row {
  display: flex;
  flex-wrap: wrap;
  justify-content: space-between;
  margin-bottom: 20px;
}
</style>
