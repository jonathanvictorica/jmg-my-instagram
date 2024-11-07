import React from 'react';
import { Routes, Route } from 'react-router-dom';
import Home from './pages/Home';
import About from './pages/About';
import NotFound from './pages/NotFound';

const App: React.FC = () => {
    return (
        <div>
            <h1>Welcome to My React App!</h1>
            <Routes>
                <Route path="/home" element={<Home />} />      {/* Página principal */}
                <Route path="/about" element={<About />} /> {/* Página de About */}
                <Route path="*" element={<NotFound />} />   {/* Página de error 404 */}
            </Routes>
        </div>
    );
};

export default App;
