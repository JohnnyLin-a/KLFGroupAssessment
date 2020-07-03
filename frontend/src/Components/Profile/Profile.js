import React, { Component } from 'react'
import Container from 'react-bootstrap/Container'
import Jumbotron from 'react-bootstrap/Jumbotron'
import Form from 'react-bootstrap/Form'
import Row from 'react-bootstrap/Row'
import Col from 'react-bootstrap/Col'
import { connect } from 'react-redux'
import { Redirect } from 'react-router-dom/cjs/react-router-dom.min'
import Button from 'react-bootstrap/Button'
import Alert from 'react-bootstrap/Alert'
import { parseJWT } from '../../Helpers/JWTHelper'
import { loginSuccess } from '../../Actions/AuthActions'


class Profile extends Component {
    constructor(props) {
        super(props);
        this.state = {
            submitDisabled: false,
            name: '',
            password: '',
            confirmPassword: '',
            successMessageName: '',
            successMessagePassword: '',
            errorMessageName: '',
            errorMessagePassword: '',
        }

        this.handleNameChange = this.handleNameChange.bind(this);
        this.handlePasswordChange = this.handlePasswordChange.bind(this);
        this.handleConfirmPasswordChange = this.handleConfirmPasswordChange.bind(this);
        this.handleSubmitChangeName = this.handleSubmitChangeName.bind(this);
        this.handleSubmitChangePassword = this.handleSubmitChangePassword.bind(this);
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

    handleSubmitChangeName(event) {
        event.preventDefault();

        this.setState({ submitDisabled: true, errorMessage: '' }, () => {
            if (this.state.name === "") {
                this.setState({ errorMessageName: "Name cannot be empty!", submitDisabled: false });
                return
            }
            if (this.state.name === this.props.user.name) {
                this.setState({ errorMessageName: "New name cannot be the same as your current name!", submitDisabled: false });
                return
            }

            fetch(`${process.env.REACT_APP_API_URL}/updatename`, {
                method: 'post',
                body: JSON.stringify({ name: this.state.name, token: this.props.user.token }),
            }).then(response => {
                switch (response.status) {
                    case 200:
                        const jsonResponse = response.json();
                        return jsonResponse;
                    case 401:
                        this.setState({ errorMessageName: "Login expired, please refresh.", submitDisabled: false })
                        break;
                    case 500:
                        this.setState({ errorMessageName: "Internal server error", submitDisabled: false })
                        break;
                    case 400:
                        this.setState({ errorMessageName: "Invalid request (app error)", submitDisabled: false })
                        break;
                    default:
                        this.setState({ errorMessageName: "Network error. Please try again.", submitDisabled: false })
                        break;
                }
            }).then(data => {
                if (typeof data !== 'undefined') {
                    const jsonObj = parseJWT(data.token);
                    this.setState({ successMessageName: "Changed name!", submitDisabled: false, name: '' }, () => {
                        this.props.loginSuccess({ token: data.token, name: jsonObj.Name });
                    })
                }
            }).catch(error => {
                console.error(error);
            });
        })
    }

    handleSubmitChangePassword(event) {
        event.preventDefault();

        this.setState({ submitDisabled: true, errorMessage: '' }, () => {
            if (this.state.password !== this.state.confirmPassword) {
                this.setState({ errorMessagePassword: "Passwords do not match!", submitDisabled: false });
                return
            }
            if (this.state.password === "") {
                this.setState({ errorMessagePassword: "Password cannot be empty!", submitDisabled: false });
                return
            }


        })
    }

    render() {
        return (
            this.props.user.token ?
                <Container className="pt-3">
                    <Jumbotron>
                        <h1 className="text-center">Profile</h1>
                        <Form>
                            <Row>
                                <Col>
                                    <h4 className="text-center m-3">Update account details</h4>
                                </Col>
                            </Row>
                            <Row className="m-3">
                                <Col>
                                    <Form.Group controlId="name">
                                        <Form.Control type="text" placeholder={this.props.user.name} value={this.state.name} onChange={this.handleNameChange} />
                                    </Form.Group>
                                </Col>
                            </Row>
                            {this.state.successMessageName && <Alert variant="success">
                                {this.state.successMessageName}
                            </Alert>}
                            {this.state.errorMessageName && <Alert variant="danger">
                                {this.state.errorMessageName}
                            </Alert>}
                            <Row>
                                <Col className="text-center">
                                    <Button variant="primary" disabled={this.state.submitDisabled} onClick={this.handleSubmitChangeName}>
                                        Update name
                                    </Button>
                                </Col>
                            </Row>

                            <hr />
                            <Row className="m-3">
                                <Col>
                                    <Form.Group controlId="name">
                                        <Form.Control type="password" placeholder="Password" value={this.state.password} onChange={this.handlePasswordChange} />
                                    </Form.Group>
                                </Col>
                                <Col>
                                    <Form.Group controlId="name">
                                        <Form.Control type="password" placeholder="Confirm password" value={this.state.confirmPassword} onChange={this.handleConfirmPasswordChange} />
                                    </Form.Group>
                                </Col>
                            </Row>

                            {this.state.successMessagePassword && <Alert variant="success">
                                {this.state.successMessagePassword}
                            </Alert>}
                            {this.state.errorMessagePassword && <Alert variant="danger">
                                {this.state.errorMessagePassword}
                            </Alert>}

                            <Row>
                                <Col className="text-center">
                                    <Button variant="primary" disabled={this.state.submitDisabled} onClick={this.handleSubmitChangePassword}>
                                        Update password
                                    </Button>
                                </Col>
                            </Row>
                        </Form>
                    </Jumbotron>
                </Container >
                :
                <Redirect to="/" />
        )
    }
}

const mapStateToProps = state => {
    return {
        user: state.user
    }
}

const mapDispatchToProps = dispatch => {
    return {
        loginSuccess: payload => {
            dispatch(loginSuccess(payload))
        },
    }
}

export default connect(mapStateToProps, mapDispatchToProps)(Profile);