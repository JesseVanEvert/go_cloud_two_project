<script>
import Navigation from './components/Navbar.vue'
export default {
  name: "App",
  components: {
    Navigation
  }
};
</script>

<template>
  
   <navigation />
   <RouterView />
</template>

<style scoped>
</style>