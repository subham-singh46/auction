import React, { useState } from 'react';
import { useNavigate, Link } from 'react-router-dom';
import config from '../config';
import './Auth.css'

function SignUp({ setIsLoggedIn }) {
    const [name, setName] = useState('');
    const [email, setEmail] = useState('');
    const [mobile, setmobile] = useState('');
    const [password, setPassword] = useState('');
    const [confirmPassword, setConfirmPassword] = useState('');
    const [error, setError] = useState('');
    const navigate = useNavigate();

    const handleSubmit = async (e) => {
        e.preventDefault();
        setError('');

        if (password !== confirmPassword) {
            setError('Passwords do not match');
            return;
        }

        const response = await fetch(`${config.API_BASE_URL}/sign-up`, {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify({ name, email, mobile, password }),
        });

        if (response.ok) {
            const data = await response.json();
            localStorage.setItem('jwtToken', data.token); // Assuming the response contains the token
            localStorage.setItem('userId', data.userId);
            setIsLoggedIn(true);
            navigate('/events'); // Redirect after successful sign-up
        } else {
            setError('Sign-up failed. Please try again.');
        }
    };

    return (
        <div className="auth-container">
            <h2>Sign Up</h2>
            <form onSubmit={handleSubmit}>
                <label>Name:</label>
                <input
                    type="text"
                    value={name}
                    onChange={(e) => setName(e.target.value)}
                    required
                />
                <label>Email:</label>
                <input
                    type="email"
                    value={email}
                    onChange={(e) => setEmail(e.target.value)}
                    required
                />
                <label>Mobile:</label>
                <input
                    type="text"
                    value={mobile}
                    onChange={(e) => setmobile(e.target.value)}
                    required
                />
                <label>Password:</label>
                <input
                    type="password"
                    value={password}
                    onChange={(e) => setPassword(e.target.value)}
                    required
                />
                <label>Confirm Password:</label>
                <input
                    type="password"
                    value={confirmPassword}
                    onChange={(e) => setConfirmPassword(e.target.value)}
                    required
                />
                {error && <p className="error">{error}</p>}
                <button type="submit">Sign Up</button>
            </form>
            <p>Already have an account? <Link to="/login">Login Here</Link></p>
        </div>
    );
}

export default SignUp;
