import React, { Component } from 'react'
import Form from 'react-bootstrap/Form'
import Button from 'react-bootstrap/Button'
import Container from 'react-bootstrap/Container'
import { connect } from 'react-redux';
import { loginSuccess } from '../../Actions/AuthActions'
import { Redirect } from 'react-router-dom'
import { parseJWT } from '../../Helpers/JWTHelper';
class Login extends Component {
    constructor(props) {
        super(props);
        this.state = {
            name: '',
            password: '',
            errorMessage: '',
        };

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

        fetch(`${process.env.REACT_APP_API_URL}/login`, {
            method: 'post',
            body: JSON.stringify({ name: this.state.name, password: this.state.password }),
        }).then(response => {
            switch (response.status) {
                case 200:
                    const jsonResponse = response.json();
                    return jsonResponse;
                case 401:
                    this.setState({ errorMessage: "Wrong Name/Password" })
                    break;
                case 500:
                    this.setState({ errorMessage: "Internal server error" })
                    break;
                case 400:
                    this.setState({ errorMessage: "Invalid login details" })
                    break;
                default:
                    this.setState({ errorMessage: "Network error. Please try again." })
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