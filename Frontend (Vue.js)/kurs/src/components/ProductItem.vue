<!--<template>
  <div class="product-line">
    <img :src="product.image" :alt="product.name" class="tshirt"/>
    <div class="buttons-line">
      <button @click="$emit('add-to-cart', product)" class="button-basket">
        В корзину
      </button>
      <button @click="$emit('add-to-favorites', product)" class="button-favourites">
        В избранное
      </button>
    </div>
  </div>
</template>

<script>
export default {
  name: 'ProductItem',
  props: {
    product: {
      type: Object,
      required: true
    }
  }
}
</script>-->


<!--<template>
  <div class="product-line">
    <img :src="product.image" :alt="product.name" class="tshirt"/>
    <div class="product-info">
      <h3 class="product-title">{{ product.text || product.name }}</h3>
      <p class="product-manufacturer" v-if="product.manufacturer">Производитель: {{ product.manufacturer }}</p>
      <p class="product-price" v-if="product.price">Цена: {{ formattedPrice }}</p>
    </div>
    <div class="buttons-line">
      <button @click="$emit('add-to-cart', product)" class="button-basket">
        В корзину
      </button>
      <button @click="$emit('add-to-favorites', product)" class="button-favourites">
        В избранное
      </button>
    </div>
  </div>
</template>

<script>
export default {
  name: 'ProductItem',
  props: {
    product: {
      type: Object,
      required: true
    }
  },
  computed: {
    formattedPrice() {
      if (!this.product.price) return 'Цена не указана';
      return new Intl.NumberFormat('ru-RU', {
        style: 'currency',
        currency: 'RUB',
        minimumFractionDigits: 0
      }).format(this.product.price);
    }
  }
}
</script>

<style scoped>
.product-line {
  display: flex;
  flex-direction: column;
  align-items: center;
  padding: 15px;
  background-color: #94BF4C;
  border-radius: 8px;
  margin: 10px;
  width: calc(100% - 20px);
  min-height: 400px; /* Увеличиваем минимальную высоту */
  justify-content: space-between; /* Равномерное распределение */
}

.tshirt {
  max-width: 100%;
  height: 180px;
  object-fit: contain;
  margin-bottom: 15px;
  background: white;
  padding: 10px;
  border-radius: 4px;
}

.product-info {
  width: 100%;
  text-align: center;
  color: white;
  margin: 10px 0;
  flex-grow: 1; /* Занимает доступное пространство */
}

.product-title {
  font-size: 18px;
  margin-bottom: 8px;
  font-weight: bold;
}

.product-manufacturer {
  font-size: 14px;
  margin-bottom: 8px;
}

.product-price {
  font-size: 16px;
  font-weight: bold;
  color: #3e3f3a;
}

.buttons-line {
  display: flex;
  justify-content: space-between;
  width: 100%;
  margin-top: 15px;
}

.button-basket, .button-favourites {
  padding: 8px 15px;
  border: none;
  border-radius: 4px;
  cursor: pointer;
  font-size: 14px;
  transition: all 0.3s;
  width: 48%; /* Ширина кнопок с небольшим отступом */
}

.button-basket {
  background-color: #3e3f3a;
  color: white;
}

.button-favourites {
  background-color: white;
  color: #3e3f3a;
  border: 1px solid #3e3f3a;
}

.button-basket:hover {
  background-color: #2a2b26;
}

.button-favourites:hover {
  background-color: #f0f0f0;
}
</style>-->






<!--<template>
  <div class="product-line">
    <img :src="product.image_url" :alt="product.name" class="tshirt"/>
    <div class="product-info">
      <h3 class="product-title">{{ product.name }}</h3>
      <p class="product-manufacturer">Производитель: {{ product.manufacturer }}</p>
      <p class="product-price">Цена: {{ formattedPrice }}</p>
    </div>
    <div class="buttons-line">
      <button @click="$emit('add-to-cart', product)" class="button-basket">
        В корзину
      </button>
      <button @click="$emit('add-to-favorites', product)" class="button-favourites">
        Избранное
      </button>
    </div>
  </div>
</template>

<script>
export default {
  name: 'ProductItem',
  props: {
    product: {
      type: Object,
      required: true
    }
  },
  computed: {
    formattedPrice() {
      return new Intl.NumberFormat('ru-RU', {
        style: 'currency',
        currency: 'RUB',
        minimumFractionDigits: 0
      }).format(this.product.price);
    }
  }
}
</script>

<style scoped>
.product-line {
  display: flex;
  flex-direction: column;
  align-items: center;
  padding: 15px;
  background-color: #94BF4C;
}

.product-info {
  width: 100%;
  margin: 10px 0;
  text-align: center;
}

.product-title {
  font-size: 18px;
  margin-bottom: 5px;
  color: #333;
}

.product-manufacturer {
  font-size: 14px;
  color: #666;
  margin-bottom: 5px;
}

.product-price {
  font-size: 16px;
  font-weight: bold;
  color:#3e3f3a;
}

.buttons-line {
  display: flex;
  justify-content: space-between;
  width: 100%;
  margin-top: 10px;
}
</style>-->


<!--<template>
  <div class="product-line">
    <img :src="product.image_url" :alt="product.name" class="tshirt"/>
    <div class="product-info">
      <h3 class="product-title">{{ product.name }}</h3>
      <p class="product-manufacturer">Производитель: {{ product.manufacturer }}</p>
      <p class="product-price">Цена: {{ formattedPrice }}</p>
    </div>
    <div class="buttons-line">
      <button @click="$emit('add-to-cart', product)" class="button-basket">
        В корзину
      </button>
      <button 
        @click="toggleFavorite(product)" 
        class="button-favourites"
        v-if="!isFavoritePage"
      >
        {{ isFavorite(product.id) ? 'В избранном' : 'В избранное' }}
      </button>
      <button 
        @click="$emit('remove-from-favorites', product.id)" 
        class="button-remove"
        v-else
      >
        Удалить из избр.
      </button>
    </div>
  </div>
</template>

<script>
import { mapState, mapActions } from 'vuex';

export default {
  name: 'ProductItem',
  props: {
    product: {
      type: Object,
      required: true
    },
    isFavoritePage: {
      type: Boolean,
      default: false
    }
  },
  computed: {
    formattedPrice() {
      return new Intl.NumberFormat('ru-RU', {
        style: 'currency',
        currency: 'RUB',
        minimumFractionDigits: 0
      }).format(this.product.price);
    },
    ...mapState(['favorites', 'userId'])
  },
  methods: {
    ...mapActions(['addToFavorites', 'removeFromFavorites']),
    isFavorite(productId) {
      return this.favorites.some(f => f.id === productId);
    },
    handleRemove() {
    this.$emit('remove-from-favorites', this.product.id);
  },
   async toggleFavorite(product) {
    try {
      if (this.isFavorite(product.id)) {
        await this.removeFromFavorites(product.id);
      } else {
        await this.addToFavorites(product.id); // Отправляем только ID
      }
    } catch (error) {
      console.error('Ошибка:', error);
      alert(error.message || 'Произошла ошибка');
    }
  }
}
  }
</script>

<style scoped>
.product-line {
  display: flex;
  flex-direction: column;
  align-items: center;
  padding: 15px;
  background-color: #94BF4C;
  position: relative;
}

.product-info {
  width: 100%;
  margin: 10px 0;
  text-align: center;
}

.product-title {
  font-size: 18px;
  margin-bottom: 5px;
  color: #333;
}

.product-manufacturer {
  font-size: 14px;
  color: #666;
  margin-bottom: 5px;
}

.product-price {
  font-size: 16px;
  font-weight: bold;
  color:#3e3f3a;
}

.buttons-line {
  display: flex;
  justify-content: space-between;
  width: 100%;
  margin-top: 10px;
}

.auth-warning {
  color: red;
  font-size: 12px;
  margin-top: 5px;
}

.button-remove {
  background-color: #ff6b6b;
  color: white;
  padding: 8px 12px;
  border: none;
  border-radius: 4px;
  cursor: pointer;
}

.button-remove:hover {
  background-color: #ff5252;
}
</style>-->

<template>
  <div class="product-line">
    <img :src="product.image_url" :alt="product.name" class="tshirt"/>
    <div class="product-info">
      <h3 class="product-title">{{ product.name }}</h3>
      <p class="product-manufacturer">Производитель: {{ product.manufacturer }}</p>
      <p class="product-price">Цена: {{ formattedPrice }}</p>
    </div>
    <div class="buttons-line">

      <button 
        @click="handleBasketAction" 
        class="button-basket"
        v-if="!isBasketPage"
      >
        {{ isInBasket ? 'В корзине' : 'В корзину' }}
      </button>

      <button 
        @click="$emit('remove-from-basket', product.id)" 
        class="button-remove"
        v-else
      >
        Удалить из корз.
      </button>

      <button 
        @click="toggleFavorite(product)" 
        class="button-favourites"
        v-if="!isFavoritePage && !isBasketPage"
      >
        {{ isFavorite(product.id) ? 'В избранном' : 'В избранное' }}
      </button>

      <button 
        @click="$emit('remove-from-favorites', product.id)" 
        class="button-remove"
        v-if="isFavoritePage && !isBasketPage"
      >
       Удалить из избр.
      </button>

    </div>
  </div>
</template>

<script>
import { mapState, mapActions, mapGetters } from 'vuex';

export default {
  name: 'ProductItem',
  props: {
    product: {
      type: Object,
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
    formattedPrice() {
      return new Intl.NumberFormat('ru-RU', {
        style: 'currency',
        currency: 'RUB',
        minimumFractionDigits: 0
      }).format(this.product.price);
    },
    ...mapState(['favorites', 'basket', 'userId']),
    isInBasket() {
      /*return this.basket.some(b => b.id === this.product.id);*/
      return this.basket ? this.basket.some(b => b.id === this.product.id) : false;
    }
  },
  methods: {
    
    ...mapActions(['addToFavorites', 'removeFromFavorites', 'addToBasket', 'removeFromBasket']),
    isFavorite(productId) {
      return this.favorites.some(f => f.id === productId);
    },
    
    async toggleFavorite(product) {
      try {
        if (this.isFavorite(product.id)) {
          await this.removeFromFavorites(product.id);
        } else {
          await this.addToFavorites(product.id);
        }
      } catch (error) {
        console.error('Ошибка:', error);
        alert(error.message || 'Произошла ошибка');
      }
    },
    async handleBasketAction() {
      try {
        if (this.isInBasket) {
          await this.removeFromBasket(this.product.id);
        } else {
          await this.addToBasket(this.product.id);
        }
      } catch (error) {
        console.error('Ошибка:', error);
        alert(error.message || 'Произошла ошибка');
      }
    }
  }
}
</script>

<style scoped>
.product-line {
  display: flex;
  flex-direction: column;
  align-items: center;
  padding: 15px;
  background-color: #94BF4C;
  position: relative;
}

.product-info {
  width: 100%;
  margin: 10px 0;
  text-align: center;
}

.product-title {
  font-size: 18px;
  margin-bottom: 5px;
  color: #333;
}

.product-manufacturer {
  font-size: 14px;
  color: #666;
  margin-bottom: 5px;
}

.product-price {
  font-size: 16px;
  font-weight: bold;
  color:#3e3f3a;
}

.buttons-line {
  display: flex;
  justify-content: space-between;
  width: 100%;
  margin-top: 10px;
}

.button-remove {
  background-color: #ff6b6b;
  color: white;
  padding: 8px 12px;
  border: none;
  border-radius: 4px;
  cursor: pointer;
}

.button-remove:hover {
  background-color: #ff5252;
}
</style>