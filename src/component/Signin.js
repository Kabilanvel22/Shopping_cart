import React, { useState } from 'react';
import { toast } from 'react-toastify';
import './Signin.css';
import { Link, useNavigate } from 'react-router-dom';
import api from './Api';  

function Signin() {
  const [username, setUsername] = useState('');
  const [password, setPassword] = useState('');
  const [errors, setErrors] = useState({});
  const nav = useNavigate();

  const handleUsernameChange = (e) => {
    setUsername(e.target.value);
  };

  const handlePasswordChange = (e) => {
    setPassword(e.target.value);
  };

  const validateForm = () => {
    let formErrors = {};
    let isValid = true;

    if (username.trim() === '') {
      formErrors.username = 'Username is required';
      isValid = false;
    }

    if (password.trim() === '') {
      formErrors.password = 'Password is required';
      isValid = false;
    }

    setErrors(formErrors);
    return isValid;
  };

  const handleSubmit = async (e) => {
    e.preventDefault();
    if (validateForm()) {
      const loginData = { username, password };

      try {
        const response = await api.login(loginData.username, loginData.password);  
        if (response.token) {
          localStorage.setItem('token', response.token); 
          toast.success('Sign in successful');
          nav('/list'); 
        } else {
          toast.error(response.error || 'Sign in failed');
        }
      } catch (error) {
        toast.error('Error during sign in');
      }
    }
  };

  return (
    <div className="form-container">
      <h2 className="form-title">Sign In</h2>
      <form className="form" onSubmit={handleSubmit}>
        <div>
          <label className="form-label" label="username">Username:</label>
          <input
            type="text"
            id="username"
            value={username}
            onChange={handleUsernameChange}
            className={errors.username ? 'error form-input' : 'form-input'}
          />
          {errors.username && <p className="error-message">{errors.username}</p>}
        </div>
        <div>
          <label className="form-label" label="password">Password:</label>
          <input
            type="password"
            id="password"
            value={password}
            onChange={handlePasswordChange}
            className={errors.password ? 'error form-input' : 'form-input'}
          />
          {errors.password && <p className="error-message">{errors.password}</p>}
        </div>
        <button className="form-button" type="submit">Sign In</button>
        <Link to="/register" className="form-link">Don't have an account? Sign up</Link>
      </form>
    </div>
  );
}

export default Signin;
