import React, { useState, useEffect } from 'react';
import '../App.css';
import api from './Api';
import {  toast } from 'react-toastify';
import 'react-toastify/dist/ReactToastify.css';


const ItemList = () => {
  const [items, setItems] = useState([]);
  const [quantities, setQuantities] = useState({}); 
  const token = localStorage.getItem('token');  

  
  useEffect(() => {
    const getItems = async () => {
      try {
        const data = await api.fetchItems();
        console.log(data); 
        setItems(data); 

        
        const initialQuantities = data.reduce((acc, item) => {
          acc[item.ID] = 1; 
          return acc;
        }, {});
        setQuantities(initialQuantities); 
      } catch (error) {
        console.error("Error fetching items:", error);
      }
    };
    getItems();
  }, []); 

 
  const handleQuantityChange = (e, itemId) => {
    const value = Math.max(1, e.target.value); 
    setQuantities((prevQuantities) => ({
      ...prevQuantities,
      [itemId]: value, 
    }));
  };


  const handleAddToCart = async (itemId) => {
    const token = localStorage.getItem('token');  
    if (!token) {
      alert('User not authenticated. Please log in.');
      return;
    }
  
    const quantity = quantities[itemId]; 
    try {
      await api.addToCart(token, itemId, quantity); 
      alert('Item added to cart');
    } catch (error) {
      alert('Error adding item to cart');
    }
  };
  

  const handleCheckout = async () => {
    try {
      await api.checkoutCart(token);
      toast.success('order placed successfully');
    } catch (error) {
      toast.error('Error during checkout');
    }
  };

  
  const handleViewCart = async () => {
    try {
      const cart = await api.fetchCart(token);
      alert(`Cart: ${JSON.stringify(cart)}`);
    } catch (error) {
      alert('Error fetching cart');
    }
  };

 
  // const handleViewOrders = async () => {
  //   try {
  //     const orders = await api.fetchOrders(token);
  //     alert(`Orders: ${JSON.stringify(orders)}`);
  //   } catch (error) {
  //     alert('Error fetching orders');
  //   }
  // };

  return (
    <div className="container">
      <h2 className="header">Items</h2>
      <div className="item-grid">
        {items && items.length > 0 ? (
          items.map((item) => (
            <div className="item-card" key={item.ID}>  
              <p className="item-name">{item.Name}</p>
              <p className="item-price">${item.Price}</p>
              <input
                type="number"
                value={quantities[item.ID] || 1} 
                onChange={(e) => handleQuantityChange(e, item.ID)} 
                min="1"
                className="quantity-input"
              />
              <button className="button" onClick={() => handleAddToCart(item.ID)}>
                Add to Cart
              </button>
            </div>
          ))
        ) : (
          <p>No items available.</p>
        )}
      </div>
      <div className="button-container">
        <button className="button button-checkout" onClick={handleCheckout}>
          Checkout
        </button>
        <button className="button" onClick={handleViewCart}>
          Cart
        </button>
      </div>
    </div>
  );
};

export default ItemList;
