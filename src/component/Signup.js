import React, { useState } from 'react';
import { toast } from 'react-toastify';
import './Signup.css';
import { useNavigate } from 'react-router-dom';
import api from './Api';  

function Signup() {
  const [username, setUsername] = useState('');
  const [password, setPassword] = useState('');
  const [confirmPassword, setConfirmPassword] = useState('');
  const [errors, setErrors] = useState({});
  const nav = useNavigate();

  const validateForm = () => {
    let formErrors = {};
    let isValid = true;

    if (!username.trim()) {
      formErrors.username = 'Username is required';
      isValid = false;
    }

    
    if (!password.trim()) {
      formErrors.password = 'Password is required';
      isValid = false;
    } else if (password.length < 8) {
      formErrors.password = 'Password must be at least 8 characters long';
      isValid = false;
    } else if (!/[A-Z]/.test(password)) {
      formErrors.password = 'Password must contain at least one uppercase letter';
      isValid = false;
    } else if (!/[a-z]/.test(password)) {
      formErrors.password = 'Password must contain at least one lowercase letter';
      isValid = false;
    } else if (!/\d/.test(password)) {
      formErrors.password = 'Password must contain at least one number';
      isValid = false;
    } else if (!/[!@#$%^&*]/.test(password)) {
      formErrors.password = 'Password must contain at least one special character';
      isValid = false;
    }

    if (!confirmPassword.trim()) {
      formErrors.confirmPassword = 'Confirm Password is required';
      isValid = false;
    } else if (password !== confirmPassword) {
      formErrors.confirmPassword = 'Passwords do not match';
      isValid = false;
    }

    setErrors(formErrors);
    return isValid;
  };

  const handleSubmit = async (e) => {
    e.preventDefault();
    if (validateForm()) {
      const userData = { username, password };

      try {
        const response = await api.registerUser(userData);  

        if (response.message) {
          toast.success('Sign up successful');
          nav('/'); 
        } else {
          toast.error(response.error || 'Sign up failed');
        }
      } catch (error) {
        toast.error('Error during signup');
      }
    }
  };

  return (
    <div className="signup-container">
      <h2 className="signup-title">Sign Up</h2>
      <form className="signup-form" onSubmit={handleSubmit}>
        <div>
          <label className="signup-label">Username</label>
          <input
            type="text"
            value={username}
            onChange={(e) => setUsername(e.target.value)}
            className={errors.username ? 'error signup-input' : 'signup-input'}
          />
          {errors.username && <p className="error-message">{errors.username}</p>}
        </div>
        <div>
          <label className="signup-label">Password</label>
          <input
            type="password"
            value={password}
            onChange={(e) => setPassword(e.target.value)}
            className={errors.password ? 'error signup-input' : 'signup-input'}
          />
          {errors.password && <p className="error-message">{errors.password}</p>}
        </div>
        <div>
          <label className="signup-label">Confirm Password</label>
          <input
            type="password"
            value={confirmPassword}
            onChange={(e) => setConfirmPassword(e.target.value)}
            className={errors.confirmPassword ? 'error signup-input' : 'signup-input'}
          />
          {errors.confirmPassword && <p className="error-message">{errors.confirmPassword}</p>}
        </div>
        <button className="signup-button" type="submit">Sign Up</button>
      </form>
    </div>
  );
}

export default Signup;
