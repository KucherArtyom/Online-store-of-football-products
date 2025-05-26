<template>
  <div class="sidebar">
    <div class="sidebar-content">
    <router-link 
      to="/tshirt" 
      class="side-button side-button-tshitr"
      active-class="active"
    >
      Футболки
    </router-link>
    <router-link 
      to="/scarfs" 
      class="side-button side-button-sharf"
      active-class="active"
    >
      Шарфы
    </router-link>
    <router-link 
      to="/balls" 
      class="side-button side-button-balls"
      active-class="active"
    >
      Мячи
    </router-link>
    <router-link 
      to="/boots" 
      class="side-button side-button-flags"
      active-class="active"
    >
      Бутсы
    </router-link>
    </div>
  </div>
</template>

<style scoped>
/* Общие стили для всех кнопок */
.side-button {
  display: flex;
  align-items: center;
  justify-content: center;
  height: 80px;
  width: auto;
  flex-grow: 1;
  font-size: 50px;
  color: #ffffff;
  text-decoration: none;
  font-family: "Red Hat Text";
  cursor: pointer;
  transition: background-color 0.3s;
}

.sidebar-content {
  display: flex;
  justify-content: space-between;
  padding: 0 20px;
  height: 80px;
}

/* Стили для конкретных кнопок */
.side-button-tshitr {
  background-color: #94BF4C;
}

.side-button-sharf {
  background-color: #94BF4C;
}

.side-button-balls {
  background-color: #94BF4C;
}

.side-button-flags {
  background-color: #94BF4C;
}

.side-button-caps {
  background-color: #94BF4C;
}

/* Стиль для активной ссылки */
.active {
  background-color: #7aa33a;
  font-weight: bold;
}

/* Эффект при наведении */
.side-button:hover {
  background-color: #7aa33a;
}
</style>