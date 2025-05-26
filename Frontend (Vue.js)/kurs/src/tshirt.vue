<!--<template>
    <ProductGrid :products="products" />
</template>

<script>
import ProductGrid from './components/ProductGrid.vue'


export default {
  name: 'TshirtPage',
  components: {
    ProductGrid
  },
  data() {
    return {
      products: [
        { id: 1, name: 'Футболка - Барселона', image: '/Barcelona.png', text: 'Футболка Barcelona', manufacturer: 'Nike', price: 9000 },
        { id: 2, name: 'Футболка - Интер', image: '/Inter.png', text: 'Футболка Inter', manufacturer: 'Nike', price: 9000 },
        { id: 3, name: 'Футболка - Манчестер сити', image: '/Manchester City.png', text: 'Футболка Manchester City', manufacturer: 'Puma', price: 8500 },
        { id: 4, name: 'Футболка - Милан', image: '/Milan.png', text: 'Футболка Milan', manufacturer: 'Puma', price: 8500 },
        { id: 5, name: 'Футболка - Болонья', image: '/Bolonia.png', text: 'Футболка Bolonia', manufacturer: 'Macron', price: 7000 },
        { id: 6, name: 'Футболка - Борнмут', image: '/Bournemouth.png', text: 'Футболка Bournemouth', manufacturer: 'Umbro', price: 7000 },
        { id: 7, name: 'Футболка - Селтик', image: '/Celtic.png', text: 'Футболка Celtic', manufacturer: 'Adidas', price: 8000 },
        { id: 8, name: 'Футболка - Реал Мадрид', image: '/Real Madrid.png', text: 'Футболка Real Madrid', manufacturer: 'Adidas', price: 9500 },
        { id: 9, name: 'Футболка - Байер', image: '/Bayer.png', text: 'Футболка Bayer', manufacturer: 'Castore', price: 7500 },
        { id: 10, name: 'Футболка - Фиорентина', image: '/Fiorentina.png', text: 'Футболка Fiorentina', manufacturer: 'Kappa', price: 7500 },
        { id: 11, name: 'Футболка - Трабзонспор', image: '/Trabsonspor.png', text: 'Футболка Trabsonspor', manufacturer: 'Joma', price: 6500 },
        { id: 12, name: 'Футболка - Вольфсбург', image: '/Wolfsburg.png', text: 'Футболка Wolfsburg', manufacturer: 'Nike', price: 8000 }
      ]
    }
  }
}
</script>-->



<template>
  <ProductGrid :products="products" />
</template>

<script>
import ProductGrid from './components/ProductGrid.vue'

export default {
  name: 'TshirtPage',
  components: {
    ProductGrid
  },
  data() {
    return {
      products: []  // Теперь данные будут загружаться из API
    }
  },
  async created() {
    try {
      const response = await fetch('http://localhost:8080/api/products/category/1')
      this.products = await response.json()
    } catch (error) {
      console.error('Ошибка загрузки товаров:', error)
    }
  }

  
}


</script>


