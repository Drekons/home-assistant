<template>
  <v-container fluid fill-height class="preloader-container">
    <v-row justify="center" align="center">
      <v-col cols="auto">
        <v-progress-circular
          :size="70"
          :width="7"
          color="primary"
          indeterminate
        >
          <v-icon size="40">mdi-home-assistant</v-icon>
        </v-progress-circular>
      </v-col>
    </v-row>
  </v-container>
</template>

<script setup>
// No additional logic needed
</script>

<style scoped>
.preloader-container {
  background-color: rgba(0, 0, 0, 0.1);
  position: fixed;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  z-index: 9999;
}
</style>
