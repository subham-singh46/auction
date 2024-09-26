import React, { useState, useEffect } from 'react';
import config from '../config'; // Import your config file for the host URL
import './Profile.css';
import dayjs from 'dayjs';

function Profile() {
    const [activeTab, setActiveTab] = useState('listings'); // "listings" or "biddings"
    const [tickets, setTickets] = useState([]); // Initialize tickets as an empty array
    const [loading, setLoading] = useState(false); // Track loading state
    const [error, setError] = useState(null); // Track any errors

    useEffect(() => {
        if (activeTab === 'listings') {
            fetchUserListings();
        }
    }, [activeTab]);

    const fetchUserListings = async () => {
        setLoading(true); // Start loading
        setError(null); // Reset errors
        const token = localStorage.getItem('jwtToken');

        if (!token) {
            console.error('No token found');
            return;
        }

        try {
            const response = await fetch(`${config.API_BASE_URL}/get-user-listing`, {
                method: 'GET',
                headers: {
                    Authorization: `Bearer ${token}`,
                },
            });

            if (!response.ok) {
                throw new Error(`Failed to fetch listings. Status: ${response.status}`);
            }

            const data = await response.json();
            if (data.Tickets) {
                setTickets(data.Tickets); // Set tickets array from the response
            } else {
                setTickets([]); // Ensure tickets is an empty array if data.Tickets is undefined
            }
        } catch (error) {
            console.error('Error fetching listings:', error);
            setError('Failed to fetch listings.');
        } finally {
            setLoading(false); // End loading
        }
    };

    return (
        <div className="profile-page">
            <h1>Profile</h1>

            <div className="tabs">
                <button
                    className={activeTab === 'listings' ? 'active' : ''}
                    onClick={() => setActiveTab('listings')}
                >
                    My Listings
                </button>
                <button
                    className={activeTab === 'biddings' ? 'active' : ''}
                    onClick={() => setActiveTab('biddings')}
                >
                    My Biddings
                </button>
            </div>

            <div className="tab-content">
                {activeTab === 'listings' && (
                    <div className="listings-section">
                        <h2>My Listings</h2>
                        {loading ? (
                            <p>Loading...</p>
                        ) : error ? (
                            <p className="error-message">{error}</p>
                        ) : tickets.length === 0 ? (
                            <p>You don't have any listings yet.</p>
                        ) : (
                            <ul>
                                {tickets.map((ticket, index) => (
                                    < li key={index} >
                                        <p><strong>Event Date:</strong> {new Date(ticket.eventDate).toLocaleDateString()}</p>
                                        <p><strong>Listed By:</strong> {ticket.listedBy}</p>
                                        <p><strong>Seats:</strong></p>
                                        <ul>
                                            {ticket.seatInfo.map((seat, index) => (
                                                <p key={index}>Seat {seat.seatNumber} (Block: {seat.block}, Level: {seat.level})</p>
                                            ))}
                                        </ul>
                                        <p><strong>Price:</strong> ₹{ticket.price}</p>
                                        <p><strong>Deadline:</strong> {dayjs(ticket.deadline).isValid() ? dayjs(ticket.deadline).format('DD/MM/YYYY') : new Date(ticket.eventDate).toLocaleDateString()}</p>
                                    </li>
                                ))}
                            </ul>
                        )}
                    </div>
                )}

                {activeTab === 'biddings' && (
                    <div className="biddings-section">
                        <h2>My Biddings</h2>
                        <p>Biddings functionality will be available soon.</p>
                    </div>
                )}
            </div>
        </div >
    );
}

export default Profile;
