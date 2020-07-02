import React, { Component } from 'react'
import Navbar from 'react-bootstrap/Navbar'
import Nav from 'react-bootstrap/Nav'
import Button from 'react-bootstrap/Button'
import NavDropdown from 'react-bootstrap/NavDropdown'
import { Link, withRouter } from 'react-router-dom'
class Header extends Component {
    state = {
        navExpanded: false
    }

    componentDidMount = () => {
        // console.log(this.props)
    }

    setNavExpanded = expanded => {
        this.setState({ navExpanded: expanded });
    }

    closeNav = () => {
        this.setState({ navExpanded: false });
    }

    render() {
        const name = 'name_name';
        const loggedIn = false;

        return (
            <Navbar bg="light" expand="lg" onToggle={this.setNavExpanded} expanded={this.state.navExpanded}>
                <Navbar.Brand href="/">AccessLink</Navbar.Brand>
                <Navbar.Toggle aria-controls="basic-navbar-nav" />
                <Navbar.Collapse id="basic-navbar-nav">
                    <Nav className="mr-auto">
                        <Link to="/" className="text-dark" style={{ textDecoration: 'none' }}>Home</Link>
                        <Link to="/about" className="ml-3 text-dark" style={{ textDecoration: 'none' }}>About</Link>
                    </Nav>
                    {loggedIn ?
                        <NavDropdown alignRight title={`Welcome, ${name}`} id="basic-nav-dropdown">
                            <NavDropdown.Item>
                                <Link className="nav-link" to="/profile">Profile</Link>
                            </NavDropdown.Item>
                            <NavDropdown.Divider />
                            <NavDropdown.Item onClick={() => { console.log("Logout!") }}>Logout</NavDropdown.Item>
                        </NavDropdown>
                        :
                        <>
                            <Button className="mr-3" onClick={() => { this.closeNav(); this.props.history.push("/login") }}>
                                Login
                            </Button>
                            <Button onClick={() => { this.closeNav(); this.props.history.push("/register") }}>
                                Register
                            </Button>
                        </>
                    }
                </Navbar.Collapse>
            </Navbar>
        )
    }
}

export default withRouter(Header);