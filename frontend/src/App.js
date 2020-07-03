import React from 'react';

import { BrowserRouter as Router, Switch, Route } from 'react-router-dom';
import Header from './Components/Header';
import Footer from './Components/Footer';
import Login from './Components/Auth/Login';
import Register from './Components/Auth/Register';
import Home from './Components/Home/Home';
import Profile from './Components/Profile/Profile';
import ContactUs from './Components/ContactUs/ContactUs';
import RefreshJWTStrategy from './Strategies/RefreshJWTStrategy';
import Container from 'react-bootstrap/Container';
import Demo from './Components/Demo/Demo';

function App() {
  return (
    <Router>
      <div>
        <RefreshJWTStrategy />
        <Header />
        <Container className="min-vh-80">
          <Switch>
            <Route path="/" exact component={Home} />
            <Route path="/login" component={Login} />
            <Route path="/demo" component={Demo} />
            <Route path="/register" component={Register} />
            <Route path="/profile" component={Profile} />
            <Route path="/contact-us" component={ContactUs} />
          </Switch>
        </Container>
        <Footer />
      </div>
    </Router>

  );
}

export default App;
