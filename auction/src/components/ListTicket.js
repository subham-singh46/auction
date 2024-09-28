import React, { useState } from 'react';
import config from '../config'; // Import the API configuration file
import './ListTicket.css';

function ListTicket() {
    const [numberOfTickets, setNumberOfTickets] = useState(1);
    const [seatInfo, setSeatInfo] = useState([{ seatNumber: '', block: '', level: '' }]);
    const [eventDate, setEventDate] = useState('');
    const [price, setPrice] = useState('');
    const [deadline, setDeadline] = useState('');
    const [errorMessage, setErrorMessage] = useState('');
    const userId = localStorage.getItem('userId');

    const handleNumberOfTicketsChange = (e) => {
        const newNumberOfTickets = parseInt(e.target.value);
        setNumberOfTickets(newNumberOfTickets);

        const newSeatInfo = [];
        for (let i = 0; i < newNumberOfTickets; i++) {
            newSeatInfo.push({
                seatNumber: seatInfo[i]?.seatNumber || '',
                block: seatInfo[i]?.block || '',
                level: seatInfo[i]?.level || ''
            });
        }
        setSeatInfo(newSeatInfo);
    };

    const handleSeatChange = (index, field, value) => {
        const updatedSeatInfo = [...seatInfo];
        updatedSeatInfo[index][field] = value;
        setSeatInfo(updatedSeatInfo);
    };

    const handleSubmit = async (e) => {
        e.preventDefault();

        // Validate auction end date format (dd/mm/yyyy)
        const deadlineRegex = /^\d{2}\/\d{2}\/\d{4}$/;
        if (!deadline.match(deadlineRegex)) {
            setErrorMessage('Auction End Date must be in dd/mm/yyyy format.');
            return;
        }

        if (!eventDate || !price || !deadline) {
            setErrorMessage('Please fill in all required fields.');
            return;
        }

        // Parse the event and deadline dates manually (convert dd/mm/yyyy to ISO format)
        const eventDateParts = eventDate.split('/');
        const eventDateISO = new Date(`${eventDateParts[2]}-${eventDateParts[1]}-${eventDateParts[0]}T00:00:00.000Z`);

        const deadlineParts = deadline.split('/');
        const deadlineISO = new Date(`${deadlineParts[2]}-${deadlineParts[1]}-${deadlineParts[0]}T00:00:00.000Z`);

        const ticketData = {
            userId,
            eventDate: eventDateISO.toISOString(),
            venue: 'DY Patil Stadium, Nerul, Navi Mumbai',
            numberOfTickets,
            seatInfo: seatInfo.map((seat, index) => ({
                seatNumber: index + 1,
                block: seat.block,
                level: parseInt(seat.level)
            })),
            price: parseInt(price),
            bestOffer: parseInt(price),
            deadline: deadlineISO.toISOString(),
        };

        try {

            const token = localStorage.getItem('jwtToken');
            const response = await fetch(`https://Auction-postship-env.eba-pzad7jme.us-east-1.elasticbeanstalk.com/api/add-ticket`, {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                    'Authorization': `Bearer ${token}`
                },
                body: JSON.stringify(ticketData),
            });

            if (response.ok) {
                setEventDate('');
                setPrice('');
                setDeadline('');
                setSeatInfo([{ seatNumber: '', block: '', level: '' }]);
                setNumberOfTickets(1);
                setErrorMessage('');
                alert('Ticket successfully listed!');
            } else {
                setErrorMessage('Failed to list the ticket. Please try again.');
            }
        } catch (error) {
            setErrorMessage('An error occurred while listing the ticket.');
        }
    };

    return (
        <div className="list-ticket-container">
            <h2>List Your Ticket</h2>
            <form onSubmit={handleSubmit} className="list-ticket-form">
                <div className="form-group">
                    <label>Event Date </label>
                    <input
                        type="text"
                        value={eventDate}
                        onChange={(e) => setEventDate(e.target.value)}
                        placeholder="Enter event date in dd/mm/yyyy format"
                        required
                    />
                </div>

                <div className="form-group">
                    <label>Number of Tickets</label>
                    <select value={numberOfTickets} onChange={handleNumberOfTicketsChange} required>
                        {[1, 2, 3, 4].map((num) => (
                            <option key={num} value={num}>
                                {num}
                            </option>
                        ))}
                    </select>
                </div>

                {seatInfo.map((seat, index) => (
                    <div key={index} className="seat-info-group">
                        <div className="form-group">
                            <label>Seat Number {index + 1}</label>
                            <input
                                type="text"
                                value={seat.seatNumber}
                                onChange={(e) => handleSeatChange(index, 'seatNumber', e.target.value)}
                                required
                            />
                        </div>
                        <div className="form-group">
                            <label>Block</label>
                            <input
                                type="text"
                                value={seat.block}
                                onChange={(e) => handleSeatChange(index, 'block', e.target.value)}
                                required
                            />
                        </div>
                        <div className="form-group">
                            <label>Level</label>
                            <input
                                type="number"
                                value={seat.level}
                                onChange={(e) => handleSeatChange(index, 'level', e.target.value)}
                                required
                            />
                        </div>
                    </div>
                ))}

                <div className="form-group">
                    <label>Looking to Sell for (₹)</label>
                    <input
                        type="number"
                        value={price}
                        onChange={(e) => setPrice(e.target.value)}
                        required
                    />
                </div>

                <div className="form-group">
                    <label>Auction End Date</label>
                    <input
                        type="text"
                        value={deadline}
                        onChange={(e) => setDeadline(e.target.value)}
                        placeholder="Enter auction end date in dd/mm/yyyy format"
                        required
                    />
                </div>

                {errorMessage && <p className="error-message">{errorMessage}</p>}

                <button type="submit" className="submit-button">
                    List Ticket
                </button>
            </form>
        </div>
    );
}

export default ListTicket;
