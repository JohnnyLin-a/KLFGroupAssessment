import React, { Component } from 'react'
import Container from 'react-bootstrap/Container';
import '../../css/Home/Home.css'
import Row from 'react-bootstrap/Row';
import Col from 'react-bootstrap/Col';
import { Link } from 'react-router-dom';
class Home extends Component {


    render() {
        return (
            <Container >
                <div className="row">
                    <div className="thumbnail text-center">
                        <img src="/img/access.jpg" alt="access" className="img-responsive img-fluid" />
                        <div className="caption caption-text">
                            <h1 className="text-white caption-title">Access Link</h1>
                            <h6 className="text-white caption-content">Your analytics when you need them</h6>
                        </div>
                    </div>
                </div>
                <Container fluid className="p-5">
                    <Row>
                        <Col sm className="align-self-center">
                            <h3 className="pb-5 text-center body-heading">Save time and resources by letting your analytics do the job!</h3>
                            <Container>
                                <h5 className="pb-5 text-center body-content">
                                    AccessLink™ keeps track of your users' activities in your app to quickly track their most frequent actions.
                                </h5>
                                <h5 className="pb-5 text-center body-content">
                                    Check out the demo <Link to="/demo">here</Link>!
                                </h5>
                            </Container>
                        </Col>
                        <Col sm className="align-self-center">
                            <Container>
                                <h5 className="text-center">
                                    Partnered with:
                                </h5>
                                <Row>
                                    <Col md className="align-self-center">
                                        {/* Shopify logo image from google images */}
                                        <img className="img-fluid" src="https://res.cloudinary.com/wagon/image/upload/c_fill,h_220,w_375/v1584632136/f2q94cvyc2kvuapkrtu6.png" alt="Shopify" />
                                    </Col>
                                    <Col md className="align-self-center">
                                        {/* Wordpress logo image from google images */}
                                        <img className="img-fluid" src="https://logos-download.com/wp-content/uploads/2016/03/WordPress_logo.png" alt="Wordpress" />
                                    </Col>
                                </Row>
                            </Container>
                        </Col>
                    </Row>

                </Container>
            </Container>
        )
    }
}

export default Home;