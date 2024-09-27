import React from 'react';
import { Link } from 'react-router-dom';
import './Navbar.css';

function Navbar({ isLoggedIn, handleLogout }) {
    return (
        <nav className="navbar">
            <div className="logo">
                <Link to="/" className="nav-link">Find My Tickets</Link>
            </div>
            <ul className="nav-links">
                <li><Link to="/events" className="nav-link">Events</Link></li>

                {isLoggedIn && (
                    <>
                        <li><Link to="/list-ticket" className="nav-link">List Ticket</Link></li> {/* Only visible if logged in */}
                        <li><Link to="/profile" className="nav-link">Profile</Link></li>
                        <li>
                            <button onClick={handleLogout} className="nav-link logout-btn">
                                Logout
                            </button>
                        </li>
                    </>
                )}

                {!isLoggedIn && (
                    <li><Link to="/login" className="nav-link">Login</Link></li>
                )}
            </ul>
        </nav>
    );
}

export default Navbar;
