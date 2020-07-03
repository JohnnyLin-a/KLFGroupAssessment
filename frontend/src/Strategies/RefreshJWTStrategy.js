import { Component } from 'react';
import { connect } from 'react-redux';
import { refreshJWTSuccess, logout } from '../Actions/AuthActions';
import { parseJWT } from '../Helpers/JWTHelper'

class RefreshJWTStrategy extends Component {
    componentDidMount = () => {
        console.log("componentDidMount this.props")
        // componentDidMount is called when page is refreshed, need to get actual expiry of jwt -45 seconds

        if (this.props.user.token !== "") {
            clearTimeout(this.timeout);
            this.startTimeout()
        }
    }



    startTimeout = () => {
        console.log("startTimeout start")
        const ms = this._getTimeoutMs();
        console.log("start timeout ms ", ms)
        if (ms > 0) {
            this.timeout = setTimeout(this.refreshJWT, ms);
        } else {
            this.props.logout()
        }
    }

    componentDidUpdate = prevProps => {
        if (prevProps.user.token !== this.props.user.token && this.props.user.token !== "") {
            console.log("componentDidUpdate true", prevProps.user, this.props.user);
            clearTimeout(this.timeout);
            this.startTimeout();
        } else {
            console.log("componentDidUpdate false")
            clearTimeout(this.timeout);
        }
    }

    refreshJWT = () => {
        console.log("refreshJWT start")
        fetch(`${process.env.REACT_APP_API_URL}/refresh`, {
            method: 'post',
            body: JSON.stringify({ token: this.props.user.token }),
        }).then(response => {
            console.log("fetch response ", response)
            switch (response.status) {
                case 200:
                    console.log("fetch success")
                    return response.json();
                case 401:
                    console.log("401 jwt expired")
                    break;
                case 500:
                    console.log("500 internal server error")
                    break;
                case 400:
                    console.log("400 bad request")
                    break;
                default:
                    console.log(`${response.status} other error`)
                    break;
            }
        }).then(data => {
            if (typeof data !== 'undefined' && this.props.user.token !== "") {
                console.log("refresh success ", data)
                this.props.refreshJWTSuccess(data);
            } else if (typeof data === 'undefined' && this.props.user.token !== "") {
                this.props.logout();
            }
        }).catch(error => {
            // network error, refresh failed
            console.error(error);
        });
    }

    _getTimeoutMs = () => {
        const jwtObj = parseJWT(this.props.user.token);
        console.log("jwt exp time", jwtObj.exp)
        const nowUnix = parseInt((new Date().getTime() / 1000).toFixed(0));
        console.log("times", jwtObj.exp, nowUnix);
        const diff = jwtObj.exp - nowUnix;
        console.log("diff", diff);

        if (diff > 0) {
            const ms = (diff - 45) * 1000
            console.log("ms", ms)
            return ms;
        }

        return -1
    }

    render() {
        return null
    }
}

const mapDispatchToProps = dispatch => {
    return {
        refreshJWTSuccess: payload => {
            dispatch(refreshJWTSuccess(payload));
        },
        logout: () => {
            dispatch(logout());
        },
    }
}

const mapStateToProps = state => {
    return {
        user: state.user
    }
}

export default connect(mapStateToProps, mapDispatchToProps)(RefreshJWTStrategy);