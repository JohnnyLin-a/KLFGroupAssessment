import React, { Component } from 'react'
import Container from 'react-bootstrap/Container';
import Jumbotron from 'react-bootstrap/Jumbotron';
import Table from 'react-bootstrap/Table';
import Spinner from 'react-bootstrap/Spinner';
import Alert from 'react-bootstrap/Alert';

class Demo extends Component {
    constructor(props) {
        super(props);

        this.state = {
            tableData: null,
            errorMessage: '',
        }
    }

    componentDidMount = () => {
        fetch(`${process.env.REACT_APP_API_URL}/demo`, {
            method: 'get',
        }).then(response => {
            switch (response.status) {
                case 200:
                    const jsonResponse = response.json();
                    return jsonResponse;
                case 500:
                    this.setState({ errorMessage: "Internal server error" })
                    break;
                default:
                    this.setState({ errorMessage: "Network error. Please try again." })
                    break;
            }
        }).then(data => {
            if (typeof data !== 'undefined') {
                this.setState({ tableData: data }, () => {
                    console.log("data", data)
                })
            }
        }).catch(error => {
            console.error(error);
        });
    }

    render() {
        return (
            <Container className="p-5">
                {this.state.errorMessage &&
                    <Alert variant="danger">
                        {this.state.errorMessage}
                    </Alert>
                }

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
                                    {this.state.tableData.th.map(th =>
                                        <th>{th}</th>
                                    )}
                                </tr>
                            </thead>
                            <tbody>
                                {this.state.tableData.td.map(row =>
                                    <tr>
                                        {row.map(cellData =>
                                            <td>{cellData}</td>
                                        )}
                                    </tr>
                                )}
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