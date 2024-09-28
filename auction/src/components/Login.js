import React, { useState } from 'react';
import { useNavigate, Link } from 'react-router-dom';
import config from '../config';
import './Auth.css'

function Login({ setIsLoggedIn }) {
    const [email, setEmail] = useState('');
    const [password, setPassword] = useState('');
    const [error, setError] = useState('');
    const navigate = useNavigate();

    const handleSubmit = async (e) => {
        e.preventDefault();
        setError('');

        const response = await fetch(`https://Auction-postship-env.eba-pzad7jme.us-east-1.elasticbeanstalk.com/api/login`, {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify({ email, password }),
        });

        if (response.ok) {
            const data = await response.json();
            localStorage.setItem('jwtToken', data.token); // Assuming the response contains the token
            localStorage.setItem('userId', data.userId);
            setIsLoggedIn(true);
            navigate('/events'); // Redirect after successful login
        } else {
            setError('Login failed. Please check your credentials.');
        }
    };

    return (
        <div className="auth-container">
            <h2>Login</h2>
            <form onSubmit={handleSubmit}>
                <label>Email:</label>
                <input
                    type="email"
                    value={email}
                    onChange={(e) => setEmail(e.target.value)}
                    required
                />
                <label>Password:</label>
                <input
                    type="password"
                    value={password}
                    onChange={(e) => setPassword(e.target.value)}
                    required
                />
                {error && <p className="error">{error}</p>}
                <button type="submit">Login</button>
            </form>
            <p>Don't have an account? <Link to="/signup">Sign Up</Link></p>
        </div>
    );
}

export default Login;
