import React from 'react';
import './Home.css';

function Home() {
    return (
        <div className="home-page">
            <h1>Welcome to Find My Tickets</h1>
            <p>
                Wanted to buy the tickets to your favourite concert, but just missed it?
                Have some tickets for a show that you can't attend? You've come to the right place.
            </p>

            {/* Add the icons section here */}
            <div className="icon-section">
                <div className="icon-item">
                    <i className="fas fa-shield-alt"></i>
                    <p>Secure Transactions</p>
                </div>
                <div className="icon-item">
                    <i className="fas fa-certificate"></i>
                    <p>Authenticated Tickets</p>
                </div>
                <div className="icon-item">
                    <i className="fas fa-thumbs-up"></i>
                    <p>100% Satisfaction</p>
                </div>
            </div>
        </div>
    );
}

export default Home;
