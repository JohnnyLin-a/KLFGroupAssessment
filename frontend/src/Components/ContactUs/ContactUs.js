import React, { Component } from 'react'
import Container from 'react-bootstrap/Container';
import Jumbotron from 'react-bootstrap/Jumbotron';
import Form from 'react-bootstrap/Form';
import Button from 'react-bootstrap/Button';
import Alert from 'react-bootstrap/Alert';

class ContactUs extends Component {
    constructor(props) {
        super(props);
        this.state = {
            name: '',
            message: '',
            errorMessage: '',
            successMessage: '',
        };

        this.handleEmailChange = this.handleEmailChange.bind(this);
        this.handleMessageChange = this.handleMessageChange.bind(this);
        this.handleSubmit = this.handleSubmit.bind(this);
    }

    handleEmailChange = event => {
        this.setState({ name: event.target.value });
    }

    handleMessageChange = event => {
        this.setState({ message: event.target.value });
    }

    handleSubmit = event => {
        event.preventDefault();

        this.setState({ successMessage: 'Successfully sent your message!' })
    }

    render() {
        return (
            <Container className="mt-3">
                <Jumbotron>
                    <h1 className="text-center p-3">Contact us</h1>
                    <h4 className="text-center">Email us directly at <a href="mailto: contact@accesslink.com">contact@accesslink.com</a></h4>
                    <h6 className="text-center">or use our message form below!</h6>
                    <Form onSubmit={this.handleSubmit}>
                        <Form.Group controlId="email">
                            <Form.Label>Your email</Form.Label>
                            <Form.Control type="text" placeholder="abc@example.com" value={this.state.email} onChange={this.handleEmailChange} />
                        </Form.Group>

                        <Form.Group controlId="message">
                            <Form.Label>Your message</Form.Label>
                            <Form.Control type="text" value={this.state.message} onChange={this.handleMessageChange} />
                        </Form.Group>

                        {this.state.successMessage && <Alert variant="success">
                            {this.state.successMessage}
                        </Alert>}
                        <Button variant="primary" type="submit">
                            Send
                        </Button>
                    </Form>
                </Jumbotron>
            </Container>
        )
    }
}



export default ContactUs;