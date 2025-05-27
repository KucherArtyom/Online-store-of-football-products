<template>
  <div id="app" class="col-md-4">
    <label>Имя*</label>
    <input type="text" v-model="formData.name" placeholder="Введите имя (2-20 символов)" class="form-control" required><br>
    <label>Фамилия*</label>
    <input type="text" v-model="formData.surname" placeholder="Введите фамилию (2-20 символов)" class="form-control" required><br>
    <label>Отчество</label>
    <input type="text" v-model="formData.patronymic" placeholder="Введите отчество (до 20 символов)" class="form-control"><br>
    <label>Телефон*</label>
    <input type="tel" v-model="formData.telephone" placeholder="Введите телефон" class="form-control" required><br>
    <label>Логин*</label>
    <input type="text" v-model="formData.login" placeholder="Введите логин (3-100 символов)" class="form-control" required><br>
    <label>Пароль*</label>
    <input type="password" v-model="formData.password" placeholder="Введите пароль (минимум 6 символов)" class="form-control" required><br>
    <button @click="registerUser" class="btn btn-primary">Зарегистрироваться</button>
    <p v-if="errorMessage" class="error-message">{{ errorMessage }}</p>
  </div>
</template>

<script>
import { mapMutations } from 'vuex';

export default {
  name: 'RegistrationPage',
  data() {
    return {
      formData: {
        name: '',
        surname: '',
        patronymic: '',
        telephone: '',
        login: '',
        password: ''
      },
      errorMessage: ''
    };
  },
  methods: {
    ...mapMutations(['login']),
    async registerUser() {
      this.errorMessage = '';
      
      if (!this.formData.name || !this.formData.surname || !this.formData.telephone || 
          !this.formData.login || !this.formData.password) {
        this.errorMessage = 'Пожалуйста, заполните все обязательные поля (помечены *)';
        return;
      }

      if (this.formData.name.length < 2 || this.formData.name.length > 20) {
        this.errorMessage = 'Имя должно быть от 2 до 20 символов';
        return;
      }

      if (this.formData.surname.length < 2 || this.formData.surname.length > 20) {
        this.errorMessage = 'Фамилия должна быть от 2 до 20 символов';
        return;
      }

      if (this.formData.patronymic && this.formData.patronymic.length > 20) {
        this.errorMessage = 'Отчество не должно превышать 20 символов';
        return;
      }

      if (this.formData.login.length < 3 || this.formData.login.length > 100) {
        this.errorMessage = 'Логин должен быть от 3 до 100 символов';
        return;
      }

      if (this.formData.password.length < 6) {
        this.errorMessage = 'Пароль должен быть не менее 6 символов';
        return;
      }

      try {
        const response = await fetch('http://localhost:8080/api/register', {
          method: 'POST',
          headers: {
            'Content-Type': 'application/json',
          },
          body: JSON.stringify(this.formData)
        });

        if (!response.ok) {
          let errorMessage = `Ошибка сервера: ${response.status}`;
          
          try {
            const text = await response.text();
            try {
              const json = JSON.parse(text);
              errorMessage = json.message || errorMessage;
            } catch (e) {
              errorMessage = text || errorMessage;
            }
          } catch (e) {
          }

          this.errorMessage = errorMessage;
          return;
        }

        const data = await response.json();
        
        this.login({
          userName: data.user.login,
          token: data.token,
          userId: data.user.id
        });
        
        this.$router.push('/');
      } catch (error) {
        console.error('Error:', error);
        this.errorMessage = 'Произошла ошибка при соединении с сервером';
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