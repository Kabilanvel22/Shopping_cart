import { Route, Routes } from 'react-router-dom';
import './App.css';
import Signup from './component/Signup';
import Signin from './component/Signin';
import ItemList from './component/ItemList';
import { ToastContainer } from 'react-toastify';



function App() {
  
  return (
    <>
      <Routes>
        <Route path="/" element={<Signin />} />
        <Route path="/register" element={<Signup />} />
        <Route path='/list' element={<ItemList/>} />
      </Routes>
      <ToastContainer />
    </>
  );
}

export default App;
