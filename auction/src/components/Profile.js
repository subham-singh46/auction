import React, { useEffect, useState } from 'react';
import config from '../config';
import './Profile.css';

function Profile() {
    const [activeTab, setActiveTab] = useState('listings'); // 'listings' or 'biddings'
    const [listings, setListings] = useState([]); // Ensure it's initialized as an empty array
    const [biddings, setBiddings] = useState([]); // Ensure it's initialized as an empty array
    const [loading, setLoading] = useState(true);
    const [error, setError] = useState(null);

    useEffect(() => {
        if (activeTab === 'listings') {
            fetchUserListings();
        } else if (activeTab === 'biddings') {
            fetchUserBids();
        }
    }, [activeTab]);

    const fetchUserListings = async () => {
        const token = localStorage.getItem('jwtToken');
        if (!token) {
            setError('No token found');
            setLoading(false);
            return;
        }

        try {
            const response = await fetch(`http://Auction-postship-env.eba-pzad7jme.us-east-1.elasticbeanstalk.com/api/get-user-listing`, {
                method: 'GET',
                headers: {
                    Authorization: `Bearer ${token}`,
                },
            });

            if (!response.ok) {
                throw new Error(`Failed to fetch listings. Status: ${response.status}`);
            }

            const data = await response.json();
            setListings(data.Tickets || []); // Default to an empty array if Tickets is undefined
        } catch (error) {
            setError('Failed to fetch listings.');
        } finally {
            setLoading(false);
        }
    };

    const fetchUserBids = async () => {
        const token = localStorage.getItem('jwtToken');
        if (!token) {
            setError('No token found');
            setLoading(false);
            return;
        }

        try {
            const response = await fetch(`http://Auction-postship-env.eba-pzad7jme.us-east-1.elasticbeanstalk.com/api/get-user-bids`, {
                method: 'GET',
                headers: {
                    Authorization: `Bearer ${token}`,
                },
            });

            if (!response.ok) {
                throw new Error(`Failed to fetch biddings. Status: ${response.status}`);
            }

            const data = await response.json();
            setBiddings(data.bids || []); // Default to an empty array if bids is undefined
        } catch (error) {
            setError('Failed to fetch biddings.');
        } finally {
            setLoading(false);
        }
    };

    const renderListings = () => {
        if (loading) return <p>Loading...</p>;
        if (error) return <p>{error}</p>;

        if (!listings || listings.length === 0) {
            return <p>No listings available.</p>;
        }

        return listings.map((listing) => (
            <div key={listing.ticketId} className="listing-item">
                <h3>{listing.eventName}</h3>
                <p><strong>Date:</strong> {new Date(listing.eventDate).toLocaleDateString()}</p>
                <p><strong>Price:</strong> ₹{listing.price}</p>
                <p><strong>Seats:</strong> {listing.seatInfo.map(seat => `#${seat.seatNumber} (Block: ${seat.block}, Level: ${seat.level})`).join(', ')}</p>
                <p><strong>Time Left:</strong> {calculateTimeLeft(listing.deadline)}</p>
            </div>
        ));
    };

    const renderBiddings = () => {
        if (loading) return <p>Loading...</p>;
        if (error) return <p>{error}</p>;

        if (!biddings || biddings.length === 0) {
            return <p>No biddings available.</p>;
        }

        return biddings.map((bid) => (
            <div key={bid.BidId} className="bidding-item">
                <h3>{bid.venue}</h3>
                <p><strong>Original Price:</strong> ₹{bid.originalPrice}</p>
                <p><strong>Your Bid:</strong> ₹{bid.bidPrice}</p>
                <p><strong>Bid Date:</strong> {new Date(bid.createdAt).toLocaleDateString()}</p>
            </div>
        ));
    };

    const calculateTimeLeft = (deadline) => {
        const difference = new Date(deadline) - new Date();
        const days = Math.floor(difference / (1000 * 60 * 60 * 24));
        const hours = Math.floor((difference / (1000 * 60 * 60)) % 24);
        const minutes = Math.floor((difference / 1000 / 60) % 60);
        const seconds = Math.floor((difference / 1000) % 60);
        return `${days}d ${hours}h ${minutes}m ${seconds}s`;
    };

    return (
        <div className="profile-section">
            <div className="profile-tabs">
                <button
                    onClick={() => setActiveTab('listings')}
                    className={activeTab === 'listings' ? 'active' : ''}
                >
                    My Listings
                </button>
                <button
                    onClick={() => setActiveTab('biddings')}
                    className={activeTab === 'biddings' ? 'active' : ''}
                >
                    My Biddings
                </button>
            </div>

            <div className="profile-content">
                {activeTab === 'listings' ? renderListings() : renderBiddings()}
            </div>
        </div>
    );
}

export default Profile;
