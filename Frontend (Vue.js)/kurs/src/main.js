//import './assets/main.css'//

/*import { createApp } from 'vue'
import App from './App.vue'
import router from './router'
import './assets/styles.css'


const app = createApp(App)

app.use(router)

app.mount('#app')*/

import { createApp } from 'vue';
import App from './App.vue';
import router from './router';
import store from './store';
import './assets/styles.css';

const app = createApp(App);

app.use(router);
app.use(store);

// Инициализируем хранилище при запуске приложения
store.commit('initializeStore');



app.mount('#app');

