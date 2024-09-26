import React, { useEffect, useState } from 'react';
import config from '../config';
import './EventsSection.css';

function EventsSection({ isLoggedIn }) {
    const [tickets, setTickets] = useState([]);
    const [view, setView] = useState('block');
    const [blockFilter, setBlockFilter] = useState('');
    const [levelFilter, setLevelFilter] = useState('');
    const [sortOption, setSortOption] = useState('deadline');
    const [loading, setLoading] = useState(true);
    const [error, setError] = useState(null);

    useEffect(() => {
        const token = localStorage.getItem('jwtToken');
        if (token) {
            fetchTickets();
        }
    }, []);

    const fetchTickets = async () => {
        const token = localStorage.getItem('jwtToken');
        if (!token) {
            setError('No token found');
            setLoading(false);
            return;
        }

        try {
            const response = await fetch(`${config.API_BASE_URL}/get-all-tickets`, {
                method: 'POST',
                headers: {
                    Authorization: `Bearer ${token}`,
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify({ limit: 100, offset: 0 }),
            });

            if (!response.ok) {
                throw new Error(`Failed to fetch tickets. Status: ${response.status}`);
            }

            const data = await response.json();
            setTickets(data.Tickets);
        } catch (error) {
            setError('Failed to fetch tickets.');
        } finally {
            setLoading(false);
        }
    };

    const filterTickets = (tickets) => {
        return tickets.filter(ticket => {
            const matchesBlock = blockFilter ? ticket.seatInfo.some(seat => seat.block === blockFilter) : true;
            const matchesLevel = levelFilter ? ticket.seatInfo.some(seat => seat.level === parseInt(levelFilter)) : true;
            return matchesBlock && matchesLevel;
        });
    };

    const sortTickets = (tickets) => {
        switch (sortOption) {
            case 'deadline':
                return tickets.sort((a, b) => new Date(a.deadline) - new Date(b.deadline));
            case 'price-asc':
                return tickets.sort((a, b) => a.price - b.price);
            case 'price-desc':
                return tickets.sort((a, b) => b.price - a.price);
            default:
                return tickets;
        }
    };

    const calculateTimeLeft = (deadline) => {
        const difference = new Date(deadline) - new Date();
        const days = Math.floor(difference / (1000 * 60 * 60 * 24));
        const hours = Math.floor((difference / (1000 * 60 * 60)) % 24);
        const minutes = Math.floor((difference / 1000 / 60) % 60);
        const seconds = Math.floor((difference / 1000) % 60);
        return `${days}d ${hours}h ${minutes}m ${seconds}s`;
    };

    const renderTicket = (ticket, index) => {
        const commonContent = (
            <>
                <h3>Coldplay Music of The Spheres World Tour</h3>
                <p><strong>Date:</strong> {new Date(ticket.eventDate).toLocaleDateString()}</p>
                <p><strong>Price:</strong> ₹{ticket.price}</p>
                <p><strong>Seats:</strong> {ticket.seatInfo.map(seat => `#${seat.seatNumber} (Block: ${seat.block}, Level: ${seat.level})`).join(', ')}</p>
                <p><strong>Time Left:</strong> {calculateTimeLeft(ticket.deadline)}</p>
                <button className="view-details-btn">View Details</button>
            </>
        );

        return view === 'block' ? (
            <div key={index} className="ticket-block">
                {commonContent}
            </div>
        ) : (
            <div key={index} className="ticket-list-item">
                {commonContent}
            </div>
        );
    };

    const renderTickets = () => {
        if (loading) return <p className="status-message">Loading tickets...</p>;
        if (error) return <p className="status-message error">Error: {error}</p>;

        let filteredTickets = filterTickets(tickets);
        let sortedTickets = sortTickets(filteredTickets);

        if (sortedTickets.length === 0) {
            return <p className="status-message">No tickets available.</p>;
        }

        return sortedTickets.map(renderTicket);
    };

    return (
        <div className="events-section">
            <h2>Coldplay Music of The Spheres World Tour</h2>

            <div className="filter-sort">
                <select onChange={(e) => setSortOption(e.target.value)} value={sortOption}>
                    <option value="deadline">Ending Soon</option>
                    <option value="price-asc">Price Low to High</option>
                    <option value="price-desc">Price High to Low</option>
                </select>

                <select onChange={(e) => setBlockFilter(e.target.value)} value={blockFilter}>
                    <option value="">All Blocks</option>
                    <option value="A">Block A</option>
                    <option value="B">Block B</option>
                    <option value="C">Block C</option>
                </select>

                <select onChange={(e) => setLevelFilter(e.target.value)} value={levelFilter}>
                    <option value="">All Levels</option>
                    <option value="1">Level 1</option>
                    <option value="2">Level 2</option>
                    <option value="3">Level 3</option>
                </select>

                <button
                    onClick={() => setView('block')}
                    className={view === 'block' ? 'active' : ''}
                >
                    Block View
                </button>
                <button
                    onClick={() => setView('list')}
                    className={view === 'list' ? 'active' : ''}
                >
                    List View
                </button>
            </div>

            <div className={`ticket-container ${view}-view`}>
                {renderTickets()}
            </div>
        </div>
    );
}

export default EventsSection;
