import React, { useState, useEffect } from 'react';
import { BrowserRouter as Router, Route, Routes } from 'react-router-dom';
import Navbar from './components/Navbar';
import Footer from './components/Footer';
import Home from './components/Home';
import ListTicket from './components/ListTicket';
import Login from './components/Login';
import Profile from './components/Profile';
import EventsSection from './components/EventsSection';
import SignUp from './components/SignUp';

function App() {
    const [isLoggedIn, setIsLoggedIn] = useState(false);

    // Simulate login check by checking localStorage or any other logic
    useEffect(() => {
        const token = localStorage.getItem('jwtToken'); // Assuming you store the token here
        if (token) {
            setIsLoggedIn(true);
        }
    }, []);

    const handleLogout = () => {
        localStorage.removeItem('jwtToken'); // Remove token on logout
        localStorage.removeItem('userId'); // Remove user ID as well
        setIsLoggedIn(false);
    };

    return (
        <Router>
            <div className="App">
                <Navbar isLoggedIn={isLoggedIn} handleLogout={handleLogout} /> {/* Pass the login state and logout function */}
                <main>
                    <Routes>
                        <Route path="/" element={<Home />} />
                        <Route path="/list-ticket" element={<ListTicket />} />
                        <Route path="/events" element={<EventsSection setIsLoggedIn={setIsLoggedIn} />} />
                        <Route path="/login" element={<Login setIsLoggedIn={setIsLoggedIn} />} />
                        <Route path="/signup" element={<SignUp setIsLoggedIn={setIsLoggedIn} />} />
                        <Route path="/profile" element={isLoggedIn ? <Profile /> : <Login setIsLoggedIn={setIsLoggedIn} />} />
                    </Routes>
                </main>
                <Footer />
            </div>
        </Router>
    );
}

export default App;
