const API_BASE_URL = 'http://localhost:8080'; 

const api = {
 
  async login(username, password) {
    const response = await fetch(`${API_BASE_URL}/users/login`, {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({ username, password }),
    });
    return response.json();
  },

 
  async fetchItems() {
    const response = await fetch(`${API_BASE_URL}/items`);
    return response.json();
  },

 
  async addToCart(token, item_id, quantity) {
    const response = await fetch(`${API_BASE_URL}/carts`, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
        Authorization: `Bearer ${token}`,
      },
      body: JSON.stringify({ item_id, quantity }), 
    });
  
    return response.json();
  },
  


  async checkoutCart(token) {
    const response = await fetch(`${API_BASE_URL}/orders`, {
      method: 'POST',
      headers: { Authorization: `Bearer ${token}` },
    });
    return response.json();
  },

  
  async fetchCart(token) {
    const response = await fetch(`${API_BASE_URL}/carts`, {
      headers: { Authorization: `Bearer ${token}` },
    });
    return response.json();
  },

 
  async fetchOrders(token) {
    const response = await fetch(`${API_BASE_URL}/orders`, {
      headers: { Authorization: `Bearer ${token}` },
    });
    return response.json();
  },


  async addItem(token, itemData) {
    const response = await fetch(`${API_BASE_URL}/items`, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
        Authorization: `Bearer ${token}`,
      },
      body: JSON.stringify(itemData),
    });
    return response.json();
  },

 
  async registerUser(userData) {
    const response = await fetch(`${API_BASE_URL}/users`, {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(userData),
    });
    return response.json();
  },
};

export default api;
