import axios from 'axios';

export const login = async (username, password) => {
  const response = await axios.post('/api/login', { username, password });
  return response.data;
};

export const getLaunchData = async () => {
  const response = await axios.get('/api/launch');
  return response.data;
};
