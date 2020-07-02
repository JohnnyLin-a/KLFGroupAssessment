import {
    LOGIN_SUCCESS
} from '../ActionTypes';

export const loginSuccess = payload => {
    return { type: LOGIN_SUCCESS, payload: payload }
}