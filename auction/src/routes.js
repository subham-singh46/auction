import { BrowserRouter as Router, Route, Routes } from 'react-router-dom';
import Home from './Home';
import Events from './Events';
import Contact from './Contact';

export default function AppRoutes() {
    return (
        <Router>
            <Routes>
                <Route path="/" element={<Home />} />
                <Route path="/events" element={<Events />} />
                <Route path="/contact" element={<Contact />} />
            </Routes>
        </Router>
    );
}
