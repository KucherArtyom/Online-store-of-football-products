import { createStore } from 'vuex';

export default createStore({
  state() {
    return {
      isAuthenticated: false,
      userName: '',
      userId: null, // Добавляем ID пользователя
      token: null,
      favorites: [], // Добавляем список избранного
      basket: [] // Добавляем состояние для корзины
    };
  },
  mutations: {
    login(state, { userName, token, userId }) {
      state.isAuthenticated = true;
      state.userName = userName;
      state.token = token;
      state.userId = userId;  // Сохраняем userId
      localStorage.setItem('authToken', token);
      localStorage.setItem('userName', userName);
      localStorage.setItem('userId', userId);  // Сохраняем в localStorage
  },
    logout(state) {
      state.isAuthenticated = false;
      state.userName = '';
      state.userId = null;
      state.token = null;
      state.favorites = [];
      localStorage.removeItem('authToken');
      localStorage.removeItem('userName');
      localStorage.removeItem('userId');
    },
    initializeStore(state) {
      const token = localStorage.getItem('authToken');
      const userName = localStorage.getItem('userName');
      const userId = localStorage.getItem('userId');  // Восстанавливаем userId
      if (token && userName && userId) {
          state.isAuthenticated = true;
          state.userName = userName;
          state.token = token;
          state.userId = parseInt(userId);  // Преобразуем в число
          this.dispatch('fetchFavorites');
          this.dispatch('fetchBasket');
      }
  },
    setFavorites(state, favorites) {
      state.favorites = favorites;
    },
    addFavorite(state, product) {
      if (!state.favorites.some(f => f.id === product.id)) {
        state.favorites.push(product);
      }
    },
    removeFavorite(state, productId) {
      state.favorites = state.favorites.filter(f => f.id !== productId);
    },
    setBasket(state, basket) {
      state.basket = basket;
    },
    addToBasket(state, product) {
      if (!state.basket.some(b => b.id === product.id)) {
        state.basket.push(product);
      }
    },
    removeFromBasket(state, productId) {
      state.basket = state.basket.filter(b => b.id !== productId);
    },
    clearBasket(state) {
      state.basket = [];
    }
  },
  actions: {
    async addToFavorites({ commit, state }, productId) {
      if (!state.isAuthenticated) {
          throw new Error('Необходима авторизация');
      }
      if (!state.userId) {
        throw new Error('User ID not available');
    }
      
      try {
          const response = await fetch('http://localhost:8080/api/favorites/add', {
              method: 'POST',
              headers: {
                  'Content-Type': 'application/json',
                  'Authorization': `Bearer ${state.token}`
              },
              body: JSON.stringify({ 
                user_id: state.userId,
                product_id: productId 
              })
          });
          
          // Клонируем ответ перед чтением
          const responseClone = response.clone();
          
          let data;
          try {
              data = await response.json();
          } catch (e) {
              // Если не удалось распарсить JSON, читаем как текст из клонированного ответа
              const text = await responseClone.text();
              throw new Error(text || `HTTP error! status: ${response.status}`);
          }
          
          if (!response.ok) {
              throw new Error(data.error || data.message || `Ошибка сервера: ${response.status}`);
          }
          
          commit('addFavorite', { id: productId });
          return true;
      } catch (error) {
          console.error('Ошибка при добавлении в избранное:', error);
          throw error;
      }
  },
    async removeFromFavorites({ commit, state }, productId) {
  try {
    console.log('Sending remove request:', { 
      user_id: state.userId, 
      product_id: productId 
    });

    const response = await fetch('http://localhost:8080/api/favorites/remove', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
        'Authorization': `Bearer ${state.token}`
      },
      credentials: 'include', // Важно для CORS с куками
      body: JSON.stringify({
        user_id: state.userId,
        product_id: productId
      })
    });

    console.log('Response status:', response.status);

    if (!response.ok) {
      const errorText = await response.text();
      console.error('Server error:', errorText);
      throw new Error(errorText || 'Server error');
    }

    const data = await response.json();
    console.log('Remove success:', data);
    
    commit('removeFavorite', productId);
    return data;
    
  } catch (error) {
    console.error('Remove favorite error:', error);
    throw new Error(`Не удалось удалить из избранного: ${error.message}`);
  }
},
    async fetchFavorites({ commit, state }) {
    if (!state.isAuthenticated || !state.userId) return;
    
    try {
      const response = await fetch(`http://localhost:8080/api/favorites?user_id=${state.userId}`, {
        headers: {
          'Authorization': `Bearer ${state.token}`
        }
      });
      
      if (!response.ok) throw new Error('Ошибка загрузки избранного');
      
      const favorites = await response.json();
      commit('setFavorites', favorites);
    } catch (error) {
      console.error('Ошибка при загрузке избранного:', error);
    }
  },
  async addToBasket({ commit, state }, productId) {
      if (!state.isAuthenticated) {
          throw new Error('Необходима авторизация');
      }
      if (!state.userId) {
          throw new Error('User ID not available');
      }
      
      try {
          const response = await fetch('http://localhost:8080/api/basket/add', {
              method: 'POST',
              headers: {
                  'Content-Type': 'application/json',
                  'Authorization': `Bearer ${state.token}`
              },
              body: JSON.stringify({ 
                  user_id: state.userId,
                  product_id: productId 
              })
          });
          
          if (!response.ok) {
              const errorData = await response.json();
              throw new Error(errorData.message || 'Ошибка сервера');
          }
          
          const data = await response.json();
          commit('addToBasket', { id: productId });
          return data;
      } catch (error) {
          console.error('Ошибка при добавлении в корзину:', error);
          throw error;
      }
    },
    async removeFromBasket({ commit, state }, productId) {
      try {
          const response = await fetch('http://localhost:8080/api/basket/remove', {
              method: 'POST',
              headers: {
                  'Content-Type': 'application/json',
                  'Authorization': `Bearer ${state.token}`
              },
              body: JSON.stringify({
                  user_id: state.userId,
                  product_id: productId
              })
          });
          
          if (!response.ok) {
              const errorData = await response.json();
              throw new Error(errorData.message || 'Ошибка сервера');
          }
          
          commit('removeFromBasket', productId);
          return await response.json();
      } catch (error) {
          console.error('Ошибка при удалении из корзины:', error);
          throw error;
      }
    },
    async fetchBasket({ commit, state }) {
      if (!state.isAuthenticated || !state.userId) return;
      
      try {
          const response = await fetch(`http://localhost:8080/api/basket?user_id=${state.userId}`, {
              headers: {
                  'Authorization': `Bearer ${state.token}`
              }
          });
          
          if (!response.ok) throw new Error('Ошибка загрузки корзины');
          
          const basket = await response.json();
          commit('setBasket', basket);
      } catch (error) {
          console.error('Ошибка при загрузке корзины:', error);
      }
    },
    async clearBasket({ commit, state }) {
  if (!state.isAuthenticated || !state.userId) return;
  
  try {
    const response = await fetch('http://localhost:8080/api/basket/clear', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
        'Authorization': `Bearer ${state.token}`
      },
      body: JSON.stringify({
        user_id: state.userId
      })
    });
    
    if (!response.ok) throw new Error('Ошибка очистки корзины');
    
    commit('clearBasket');
  } catch (error) {
    console.error('Ошибка при очистке корзины:', error);
    throw error;
  }
}
  }
});