<template>
  <Preloader v-if="loading" />
  <LoginForm v-else-if="!isAuthenticated" />
  <LaunchPage v-else />
</template>

<script setup>
import { ref, onMounted } from 'vue';
import { useRouter } from 'vue-router';
import LoginForm from "@/components/LoginForm.vue";
import LaunchPage from "@/components/LaunchPage.vue";
import Preloader from "@/components/Preloader.vue";
import { isTokenValid } from '@/utils/auth';
import { getLaunchData } from '@/services/api';

const router = useRouter();
const loading = ref(true);
const isAuthenticated = ref(false);

onMounted(async () => {
  if (isTokenValid()) {
    try {
      await getLaunchData();
      isAuthenticated.value = true;
    } catch (error) {
      console.error('Failed to get launch data:', error);
      isAuthenticated.value = false;
    }
  } else {
    isAuthenticated.value = false;
  }
  loading.value = false;
});
</script>
