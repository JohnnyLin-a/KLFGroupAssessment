import React, { Component } from 'react'
import Form from 'react-bootstrap/Form'
import Button from 'react-bootstrap/Button'
import Container from 'react-bootstrap/Container'
class Register extends Component {
    constructor(props) {
        super(props);
        this.state = {
            name: '',
            password: '',
            confirmPassword: '',
        };

        this.handleNameChange = this.handleNameChange.bind(this);
        this.handlePasswordChange = this.handlePasswordChange.bind(this);
        this.handleConfirmPasswordChange = this.handleConfirmPasswordChange.bind(this);
        this.handleSubmit = this.handleSubmit.bind(this);
    }

    handleNameChange(event) {
        this.setState({ name: event.target.value });
    }

    handlePasswordChange(event) {
        this.setState({ password: event.target.value });
    }

    handleConfirmPasswordChange(event) {
        this.setState({ confirmPassword: event.target.value });
    }

    handleSubmit(event) {
        event.preventDefault();
        console.log("Login with", this.state)
    }

    render() {
        return (
            <Container>
                <Form onSubmit={this.handleSubmit}>
                    <Form.Group controlId="register_name">
                        <Form.Label>Name</Form.Label>
                        <Form.Control type="text" placeholder="Name" value={this.state.name} onChange={this.handleNameChange} />
                    </Form.Group>

                    <Form.Group controlId="register_password">
                        <Form.Label>Password</Form.Label>
                        <Form.Control type="password" placeholder="Password" value={this.state.password} onChange={this.handlePasswordChange} />
                    </Form.Group>

                    <Form.Group controlId="register_confirm_password">
                        <Form.Label>Confirm Password</Form.Label>
                        <Form.Control type="password" placeholder="Confirm Password" value={this.state.confirmPassword} onChange={this.handleConfirmPasswordChange} />
                    </Form.Group>

                    <Button variant="primary" type="submit">
                        Register
                    </Button>
                </Form>
            </Container>

        )
    }
}

export default Register;