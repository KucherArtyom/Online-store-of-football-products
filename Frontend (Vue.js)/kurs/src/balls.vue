<template>
  <ProductGrid :products="products" />
</template>

<script>
import ProductGrid from './components/ProductGrid.vue'

export default {
  name: 'BallsPage',
  components: {
    ProductGrid
  },
  data() {
    return {
      products: []
    }
  },
  async created() {
    try {
      const response = await fetch('http://localhost:8080/api/products/category/3')
      this.products = await response.json()
    } catch (error) {
      console.error('Ошибка загрузки товаров:', error)
    }
  }
}
</script>