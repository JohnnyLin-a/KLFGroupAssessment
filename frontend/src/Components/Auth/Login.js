import React, { Component } from 'react'
import Form from 'react-bootstrap/Form'
import Button from 'react-bootstrap/Button'
import Container from 'react-bootstrap/Container'
class Login extends Component {
    constructor(props) {
        super(props);
        this.state = { name: '', password: '' };

        this.handleNameChange = this.handleNameChange.bind(this);
        this.handlePasswordChange = this.handlePasswordChange.bind(this);
        this.handleSubmit = this.handleSubmit.bind(this);
    }

    handleNameChange(event) {
        this.setState({ name: event.target.value });
    }

    handlePasswordChange(event) {
        this.setState({ password: event.target.value });
    }

    handleSubmit(event) {
        event.preventDefault();
        console.log("Login with", this.state)
    }

    render() {
        return (
            <Container>
                <Form onSubmit={this.handleSubmit}>
                    <Form.Group controlId="login_name">
                        <Form.Label>Name</Form.Label>
                        <Form.Control type="text" placeholder="Name" value={this.state.value} onChange={this.handleNameChange} />
                    </Form.Group>

                    <Form.Group controlId="login_password">
                        <Form.Label>Password</Form.Label>
                        <Form.Control type="password" placeholder="Password" value={this.state.value} onChange={this.handlePasswordChange} />
                    </Form.Group>
                    <Button variant="primary" type="submit">
                        Login
                    </Button>
                </Form>
            </Container>

        )
    }
}

export default Login;