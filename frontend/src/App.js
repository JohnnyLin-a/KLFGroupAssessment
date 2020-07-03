import React, { Component } from 'react';

import { BrowserRouter as Router, Switch, Route } from 'react-router-dom';
import Header from './Components/Header';
import Footer from './Components/Footer';
import Login from './Components/Auth/Login';
import Register from './Components/Auth/Register';
import About from './Components/About/About';
import Home from './Components/Home/Home';
import Profile from './Components/Profile/Profile'
import RefreshJWTStrategy from './Strategies/RefreshJWTStrategy';

function App() {
  return (
    <Router>
      <div>
        <RefreshJWTStrategy />
        <Header />
        <Switch>
          <Route path="/" exact component={Home} />
          <Route path="/login" component={Login} />
          <Route path="/about" component={About} />
          <Route path="/register" component={Register} />
          <Route path="/profile" component={Profile} />
        </Switch>

        <Footer />
      </div>
    </Router>

  );
}

export default App;
