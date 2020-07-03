import React, { Component } from 'react'
import Navbar from 'react-bootstrap/Navbar'
import Nav from 'react-bootstrap/Nav'
import Button from 'react-bootstrap/Button'
import NavDropdown from 'react-bootstrap/NavDropdown'
import { Link, withRouter } from 'react-router-dom'
import { connect } from 'react-redux';
import { logout } from '../Actions/AuthActions'
import Container from 'react-bootstrap/Container'
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
        return (
            <Navbar bg="light" expand="lg" onToggle={this.setNavExpanded} expanded={this.state.navExpanded}>
                <Container>
                    <Navbar.Brand><Link to="/" className="m-3 text-dark" style={{ textDecoration: 'none' }}>AccessLink</Link></Navbar.Brand>
                    <Navbar.Toggle aria-controls="basic-navbar-nav" />
                    <Navbar.Collapse id="basic-navbar-nav">
                        <Nav className="mr-auto">
                            <Link to="/" className="m-3 text-dark" style={{ textDecoration: 'none' }}>Home</Link>
                            <Link to="/about" className="m-3 text-dark" style={{ textDecoration: 'none' }}>About</Link>
                            <Link to="/contact-us" className="m-3 text-dark" style={{ textDecoration: 'none' }}>Contact us</Link>
                        </Nav>
                        {this.props.user.token ?
                            <NavDropdown alignRight title={`Welcome, ${this.props.user.name}`} id="basic-nav-dropdown">
                                <NavDropdown.Item>
                                    <Link className="nav-link" to="/profile">Profile</Link>
                                </NavDropdown.Item>
                                <NavDropdown.Divider />
                                <NavDropdown.Item onClick={() => { this.props.dispatch(logout()) }}>Logout</NavDropdown.Item>
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
                </Container>
            </Navbar>
        )
    }
}

const mapStateToProps = state => {
    return {
        user: state.user
    }
}
export default connect(mapStateToProps)(withRouter(Header));