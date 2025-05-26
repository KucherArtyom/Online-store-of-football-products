<!--<template>
  <div>
    <div v-for="(row, index) in chunkedProducts" :key="index" class="product">
      <ProductItem 
        v-for="product in row" 
        :key="product.id"
        :product="product"
        @add-to-cart="addToCart"
        @add-to-favorites="addToFavorites"
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
    }
  },
  computed: {
    chunkedProducts() {
      const chunkSize = 4
      const result = []
      for (let i = 0; i < this.products.length; i += chunkSize) {
        result.push(this.products.slice(i, i + chunkSize))
      }
      return result
    }
  },
  methods: {
    addToCart(product) {
      // Handle add to cart logic
      console.log('Added to cart:', product.name)
    },
    addToFavorites(product) {
      // Handle add to favorites logic
      console.log('Added to favorites:', product.name)
    }
  }
}
</script>-->

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
