<template>
  <div id="app" class="col-md-4">
    <h2>Стоимость заказа: {{ totalPrice }} руб.</h2>
    <h3>Данные доставки</h3>
    <label>Страна</label>
    <input type="text" v-model="country" placeholder="Введите страну" class="form-control" /><br>
    <label>Город</label><br>
    <input type="text" v-model="city" placeholder="Введите город" class="form-control" /><br>
    <label>Улица</label>
    <input type="text" v-model="street" placeholder="Введите улицу" class="form-control" /><br>
    <label>Дом</label><br>
    <input type="number" v-model="house" placeholder="Введите номер дом" class="form-control" /><br>
    <label>Квартира</label>
    <input type="number" v-model="apartment" placeholder="Введите номер квартиры" class="form-control" /><br>
    <h3>Платежные данные</h3>
    <label>Номер банковской карты</label>
    <input type="number " v-model="bank_card_number" placeholder="Введите номер банковской карты" class="form-control" /><br>
    <!--<button @click="registerUser" class="btn btn-primary">Оформить заказ</button>-->
    <button @click="placeOrder" class="btn btn-primary" :disabled="isProcessing">
      {{ isProcessing ? 'Оформляем заказ...' : 'Оформить заказ' }}
    </button>
    
    <p v-if="errorMessage" class="error-message">{{ errorMessage }}</p>
</div>
</template>


<script>
import { mapState, mapMutations } from 'vuex';

export default {
  name: 'OrderPage',
  data() {
    return {
      country: '',
      city: '',
      street: '',
      house: '',
      apartment: '',
      bank_card_number: '',
      errorMessage: '',
      isProcessing: false
    };
  },
  computed: {
    ...mapState(['basket', 'userId', 'token']), // Добавляем token в mapState
    totalPrice() {
      return this.basket?.reduce((sum, product) => sum + product.price, 0) || 0
    }
  },
  methods: {
    ...mapMutations(['clearBasket']),
    
async placeOrder() {
      this.errorMessage = '';
      this.isProcessing = true;

      if (!this.country || !this.city || !this.street || !this.house || !this.bank_card_number) {
        this.errorMessage = 'Пожалуйста, заполните все обязательные поля';
        this.isProcessing = false;
        return;
      }

      if (!this.basket?.length) {
        this.errorMessage = 'Корзина пуста';
        this.isProcessing = false;
        return;
      }

      try {
        const orderData = {
          customer_id: this.userId,
          order_price: this.totalPrice,
          card_number: this.bank_card_number,
          products: this.basket.map(p => p.id),
          address: {
            country: this.country,
            city: this.city,
            street: this.street,
            house: Number(this.house),
            apartment: Number(this.apartment) || null
          }
        };

        const response = await fetch('http://localhost:8080/api/orders', {
          method: 'POST',
          headers: {
            'Content-Type': 'application/json',
            'Authorization': `Bearer ${this.token}`
          },
          body: JSON.stringify(orderData)
        });

        if (!response.ok) {
          const errorText = await response.text();
          throw new Error(errorText || `Ошибка сервера: ${response.status}`);
        }

        const data = await response.json();
        await this.clearBasket();
        alert('Заказ успешно оформлен!');
        this.$router.push({ path: '/', query: { orderSuccess: 'true' } });
      } catch (error) {
        console.error("Ошибка оформления заказа:", error);
        this.errorMessage = error.message.includes("Failed to") 
          ? "Ошибка на сервере. Пожалуйста, попробуйте позже."
          : error.message;
      } finally {
        this.isProcessing = false;
      }
    }
  }
}
</script>

<style scoped>
.error-message {
  color: red;
  margin-top: 10px;
}

.col-md-4 {
  max-width: 600px;
  margin: 0 auto;
  padding: 20px;
}

.form-control {
  width: 100%;
  padding: 8px;
  margin-bottom: 10px;
}

.btn-primary {
  background-color: #94BF4C;
  color: white;
  padding: 10px 15px;
  border: none;
  border-radius: 4px;
  cursor: pointer;
  font-size: 16px;
}

.btn-primary:hover {
  background-color: #7aa33a;
}

.btn-primary:disabled {
  background-color: #cccccc;
  cursor: not-allowed;
}

h2 {
  color: #333;
  margin-bottom: 20px;
}

h3 {
  color: #555;
  margin: 20px 0 10px 0;
}
</style>