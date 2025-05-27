<template>
  <div id="app" class="col-md-4">
    <label>Логин</label>
    <input type="text" v-model="login" placeholder="Введите логин" class="form-control" /><br>
    <label>Пароль</label>
    <input type="password" v-model="password" placeholder="Введите пароль" class="form-control" /><br>
    <button @click="loginUser" class="btn btn-primary">Авторизоваться</button>
    <p v-if="errorMessage" class="error-message">{{ errorMessage }}</p>
  </div>
</template>

<script>
import { mapMutations } from 'vuex';

export default {
  name: 'AuthorisationPage',
  data() {
    return {
      login: '',
      password: '',
      errorMessage: ''
    };
  },
  methods: {
    ...mapMutations({
      vuexLogin: 'login' 
    }),
    
   async loginUser() {
  this.errorMessage = '';
  
  if (!this.login || !this.password) {
    this.errorMessage = 'Пожалуйста, введите логин и пароль';
    return;
  }
  
  try {
    const response = await fetch('http://localhost:8080/api/login', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({
        login: this.login,
        password: this.password
      }),
      credentials: 'include' 
    });
    
    const data = await response.json();
        
    if (response.ok) {
      localStorage.setItem('authToken', data.token);
      
      this.$store.commit('login', {
        userName: data.user.login,
        token: data.token,
        userId: data.user.id
      });
      
      this.$router.push('/');
    } else {
      this.errorMessage = data.message || 'Ошибка авторизации';
    }
  } catch (error) {
    console.error('Error:', error);
    this.errorMessage = 'Произошла ошибка при авторизации';
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
  max-width: 400px;
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
}

.btn-primary:hover {
  background-color: #7aa33a;
}
</style>