import React from 'react';
import './Modal.css';

function Modal({ children, closeModal }) {
    return (
        <div className="modal-overlay">
            <div className="modal-content">
                <button className="modal-close-btn" onClick={closeModal}>
                    &times; {/* This is the close button in the top-right corner */}
                </button>
                {children}
            </div>
        </div>
    );
}

export default Modal;
