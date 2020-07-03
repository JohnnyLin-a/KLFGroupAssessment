import React, { Component } from 'react'
import Form from 'react-bootstrap/Form'
import Button from 'react-bootstrap/Button'
import Container from 'react-bootstrap/Container'
import Alert from 'react-bootstrap/Alert'
import { parseJWT } from '../../Helpers/JWTHelper'
import { connect } from 'react-redux'
import { loginSuccess } from '../../Actions/AuthActions'
import { Redirect } from 'react-router-dom/cjs/react-router-dom.min'
class Register extends Component {
    constructor(props) {
        super(props);
        this.state = {
            name: '',
            password: '',
            confirmPassword: '',
            errorMessage: '',
            submitDisabled: false,
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
        console.log("Register with", this.state)
        if (this.state.name !== '' && this.state.password !== '' && this.state.password === this.state.confirmPassword) {
            this.setState({ errorMessage: '', submitDisabled: true }, () => {
                fetch(`${process.env.REACT_APP_API_URL}/register`, {
                    method: 'post',
                    body: JSON.stringify({ name: this.state.name, password: this.state.password }),
                }).then(response => {
                    switch (response.status) {
                        case 200:
                            const jsonResponse = response.json();
                            return jsonResponse;
                        case 401:
                            this.setState({ errorMessage: "Name already exists", submitDisabled: false })
                            break;
                        case 500:
                            this.setState({ errorMessage: "Internal server error", submitDisabled: false })
                            break;
                        case 400:
                            this.setState({ errorMessage: "Invalid register details", submitDisabled: false })
                            break;
                        default:
                            this.setState({ errorMessage: "Network error. Please try again.", submitDisabled: false })
                            break;
                    }
                }).then(data => {
                    if (typeof data !== 'undefined') {
                        const jsonObj = parseJWT(data.token);
                        console.log(jsonObj);
                        this.props.loginSuccess({ token: data.token, name: jsonObj.Name });
                    }
                }).catch(error => {
                    console.error(error);
                });
            })
        } else {
            this.setState({ errorMessage: "Make sure your name and password is valid, and passwords must match", submitDisabled: false })
        }
    }

    render() {
        return (
            this.props.user.token ?
                <Redirect to="/" />
                :
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

                        {this.state.errorMessage && <Alert variant="danger">
                            {this.state.errorMessage}
                        </Alert>}

                        <Button variant="primary" type="submit" disabled={this.state.submitDisabled}>
                            Register
                    </Button>
                    </Form>
                </Container>

        )
    }
}

const mapDispatchToProps = dispatch => {
    return {
        loginSuccess: payload => {
            dispatch(loginSuccess(payload))
        },
    }
}

const mapStateToProps = state => {
    return {
        user: state.user
    }
}

export default connect(mapStateToProps, mapDispatchToProps)(Register);