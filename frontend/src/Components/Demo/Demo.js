import React, { Component } from 'react'
import Container from 'react-bootstrap/Container';
import Jumbotron from 'react-bootstrap/Jumbotron';
import Table from 'react-bootstrap/Table';
import Spinner from 'react-bootstrap/Spinner';

class Demo extends Component {
    constructor(props) {
        super(props);

        this.state = {
            tableData: true,
        }
    }

    render() {
        return (
            <Container className="p-5">
                <Jumbotron>
                    <h1 className="text-center">AccessLink demo</h1>
                    <h6 className="text-center p-3">
                        Displaying data from October 2019
                    </h6>
                    <hr />
                    {this.state.tableData ?
                        <Table responsive>
                            <thead>
                                <tr>
                                    <th>Table heading</th>
                                    <th>Table heading</th>
                                    <th>Table heading</th>
                                    <th>Table heading</th>
                                    <th>Table heading</th>
                                </tr>
                            </thead>
                            <tbody>
                                <tr>
                                    <td>Table cell</td>
                                    <td>Table cell</td>
                                    <td>Table cell</td>
                                    <td>Table cell</td>
                                    <td>Table cell</td>
                                </tr>
                                <tr>
                                    <td>Table cell</td>
                                    <td>Table cell</td>
                                    <td>Table cell</td>
                                    <td>Table cell</td>
                                    <td>Table cell</td>
                                </tr>
                                <tr>
                                    <td>Table cell</td>
                                    <td>Table cell</td>
                                    <td>Table cell</td>
                                    <td>Table cell</td>
                                    <td>Table cell</td>
                                </tr>
                            </tbody>
                        </Table>
                        :
                        <div className="text-center">
                            <h5 className="p-3">Loading</h5>
                            <Spinner animation="border" role="status">
                                <span className="sr-only m-3" >Loading...</span>
                            </Spinner>
                        </div>


                    }
                </Jumbotron>
            </Container>
        )
    }
}

export default Demo;