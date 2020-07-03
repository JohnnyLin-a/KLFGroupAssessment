import {
    LOGIN_SUCCESS,
    REFRESH_JWT_SUCCESS,
    LOGOUT,
} from '../ActionTypes';

export const loginSuccess = payload => {
    return { type: LOGIN_SUCCESS, payload: payload }
}

export const refreshJWTSuccess = payload => {
    return { type: REFRESH_JWT_SUCCESS, payload: payload }
}

export const logout = () => {
    return { type: LOGOUT }
}