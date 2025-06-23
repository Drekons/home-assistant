import axios from 'axios';
import CryptoJS from 'crypto-js';

const SECRET_KEY = 'your-secret-key'; // Замените на реальный секретный ключ

const encrypt = (data) => {
  return CryptoJS.AES.encrypt(JSON.stringify(data), SECRET_KEY).toString();
};

const decrypt = (ciphertext) => {
  const bytes = CryptoJS.AES.decrypt(ciphertext, SECRET_KEY);
  return JSON.parse(bytes.toString(CryptoJS.enc.Utf8));
};

export const getToken = () => {
  const encryptedToken = localStorage.getItem('encryptedToken');
  if (!encryptedToken) return null;
  return decrypt(encryptedToken);
};

export const setToken = (token, expiresIn) => {
  const tokenData = {
    token,
    expiresAt: Date.now() + expiresIn * 1000
  };
  const encryptedToken = encrypt(tokenData);
  localStorage.setItem('encryptedToken', encryptedToken);
};

export const removeToken = () => localStorage.removeItem('encryptedToken');

export const isTokenValid = () => {
  const tokenData = getToken();
  if (!tokenData) return false;
  return Date.now() < tokenData.expiresAt;
};

export const setupAxiosInterceptors = () => {
  axios.interceptors.request.use(
    (config) => {
      const tokenData = getToken();
      if (tokenData && isTokenValid()) {
        config.headers['Authorization'] = `Bearer ${tokenData.token}`;
      }
      return config;
    },
    (error) => Promise.reject(error)
  );

  axios.interceptors.response.use(
    (response) => response,
    (error) => {
      if (error.response && error.response.status === 401) {
        removeToken();
        // Redirect to login page or refresh token
      }
      return Promise.reject(error);
    }
  );
};
