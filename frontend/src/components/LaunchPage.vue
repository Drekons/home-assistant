<template>
  <div>
    <h1>Launch Page</h1>
    <!-- Здесь будет содержимое страницы Launch -->
  </div>
</template>

<script setup>
import { onMounted } from 'vue';
import { getLaunchData } from '@/services/api';

onMounted(async () => {
  try {
    const data = await getLaunchData();
    console.log('Launch data:', data);
    // Здесь вы можете обработать полученные данные
  } catch (error) {
    console.error('Failed to get launch data:', error);
  }
});
</script>
