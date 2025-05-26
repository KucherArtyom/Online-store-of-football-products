<template>
  <div id="app" class="col-md-4">
    <label>Имя</label>
    <input type="text" v-model="name" placeholder="Введите имя" class="form-control" /><br>
    <label>Фамилия</label><br>
    <input type="text" v-model="surname" placeholder="Введите фамилию" class="form-control" /><br>
    <label>Отчество</label>
    <input type="text" v-model="patronymic" placeholder="Введите отчество" class="form-control" /><br>
    <label>Телефон</label><br>
    <input type="text" v-model="telephone" placeholder="Введите телефон" class="form-control" /><br>
    <label>Логин</label>
    <input type="text" v-model="userLogin" placeholder="Введите логин" class="form-control" /><br>
    <label>Пароль</label>
    <input type="password" v-model="password" placeholder="Введите пароль" class="form-control" /><br>
    <button @click="registerUser" class="btn btn-primary">Зарегистрироваться</button>

</div>
</template>

<script>
import { mapMutations } from 'vuex';

export default {
  name: 'RegistrationPage',
  data() {
    return {
      name: '',
      surname: '',
      patronymic: '',
      telephone: '',
      userLogin: '',
      password: '',
      errorMessage: ''
    };
  },
  methods: {
    ...mapMutations(['login']),
    async registerUser() {
      this.errorMessage = '';
      
      if (!this.name || !this.surname || !this.telephone || !this.login || !this.password) {
        this.errorMessage = 'Пожалуйста, заполните все обязательные поля';
        return;
      }

      try {
        const response = await fetch('http://localhost:8080/api/register', {
          method: 'POST',
          headers: {
            'Content-Type': 'application/json',
          },
          body: JSON.stringify({
            name: this.name,
            surname: this.surname,
            patronymic: this.patronymic,
            telephone: this.telephone,
            login: this.userLogin,
            password: this.password
          })
        });
        
        const data = await response.json();
        if (response.ok) {
          // Автоматически входим после регистрации
          this.login({
            userName: this.userLogin,
            token: data.token || 'dummy-token' // Замените на реальный токен из ответа
          });
          
          // Перенаправляем на главную страницу
          this.$router.push('/');
        } else {
          alert(data.message || 'Ошибка регистрации');
        }
      } catch (error) {
        console.error('Error:', error);
        alert('Произошла ошибка при регистрации');
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