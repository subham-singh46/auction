import React from 'react';
import './Contact.css';

function Contact() {
    return (
        <div className="contact-page">
            <h1>Contact Us</h1>
            <form>
                <label>
                    Name:
                    <input type="text" name="name" />
                </label>
                <label>
                    Email:
                    <input type="email" name="email" />
                </label>
                <label>
                    Message:
                    <textarea name="message" />
                </label>
                <button type="submit">Submit</button>
            </form>
        </div>
    );
}

export default Contact;
