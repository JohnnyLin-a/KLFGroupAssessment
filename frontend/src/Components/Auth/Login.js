import React, { Component } from 'react'
import Form from 'react-bootstrap/Form'
import Button from 'react-bootstrap/Button'
import Container from 'react-bootstrap/Container'
import { connect } from 'react-redux';
import { loginSuccess } from '../../Actions/AuthActions'
import { Redirect } from 'react-router-dom'
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
        console.log("Login with", this.state);


        fetch(`${process.env.REACT_APP_API_URL}/login`, {
            method: 'post',
            body: JSON.stringify({ name: this.state.name, password: this.state.password }),
        }).then(response => {
            console.log("Login response code " + response.status, response)
            switch (response.status) {
                case 200:
                    console.log("login success response body", response.body);
                    this.props.loginSuccess({ token: response.body });
                    break;
                case 401:
                    console.log("login unauthorized");
                    break;
                case 500:
                    console.log("login server error");
                    break;
                case 400:
                    console.log("login sent a bad request to api");
                    break;
                default:
                    console.log("login received something completely off", response.status);
                    break;
            }
        }).then(data => {
            console.log(data);
        }).catch(error => {
            console.error(error);
        });
    }

    render() {

        return (
            this.props.user.token ?
                <Redirect to="/" />
                :
                <Container>
                    <Form onSubmit={this.handleSubmit}>
                        <Form.Group controlId="login_name">
                            <Form.Label>Name</Form.Label>
                            <Form.Control type="text" placeholder="Name" value={this.state.name} onChange={this.handleNameChange} />
                        </Form.Group>

                        <Form.Group controlId="login_password">
                            <Form.Label>Password</Form.Label>
                            <Form.Control type="password" placeholder="Password" value={this.state.password} onChange={this.handlePasswordChange} />
                        </Form.Group>
                        <Button variant="primary" type="submit">
                            Login
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

export default connect(mapStateToProps, mapDispatchToProps)(Login);