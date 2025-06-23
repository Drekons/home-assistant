<template>
  <v-container class="fill-height" fluid>
    <v-row align="center" justify="center">
      <v-col cols="12" sm="8" md="6" lg="4">
        <v-card class="elevation-12">
          <v-toolbar color="primary" dark flat>
            <v-toolbar-title>Login</v-toolbar-title>
          </v-toolbar>
          <v-card-text>
            <v-form @submit.prevent="login" ref="form">
              <v-text-field
                v-model="username"
                :rules="[v => !!v || 'Username is required']"
                label="Username"
                name="username"
                prepend-icon="mdi-account"
                type="text"
                required
              ></v-text-field>
              <v-text-field
                v-model="password"
                :rules="[v => !!v || 'Password is required']"
                label="Password"
                name="password"
                prepend-icon="mdi-lock"
                type="password"
                required
              ></v-text-field>
            </v-form>
            <v-alert
              v-if="error"
              type="error"
              class="mt-4"
            >
              {{ error }}
            </v-alert>
          </v-card-text>
          <v-card-actions>
            <v-btn
              color="primary"
              @click="login"
              block
              class="mt-4"
              :loading="loading"
            >
              Login
            </v-btn>
          </v-card-actions>
        </v-card>
      </v-col>
    </v-row>
  </v-container>
</template>

<script>
import axios from 'axios';
import {setToken, isTokenValid} from '@/utils/auth';

export default {
  name: 'LoginForm',
  data() {
    return {
      username: '',
      password: '',
      error: null,
      loading: false,
    }
  },
  methods: {
    async login() {
      this.error = null;
      if (!await this.$refs.form.validate()) {
        this.error = 'Please fill in all required fields';
        return;
      }

      this.loading = true;
      try {
        const response = await axios.post('/api/login', {
          username: this.username,
          password: this.password
        });
        const {token, expiresIn} = response.data;
        setToken(token, expiresIn);
        if (!isTokenValid()) {
          this.error = 'Token is invalid. Please try again.';
        }
      } catch (error) {
        console.error('Login failed', error);
        this.error = error.response?.data?.message || 'Login failed. Please try again.';
      } finally {
        this.loading = false;
      }
    }
  }
}
</script>
