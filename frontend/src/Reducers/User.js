import {
    LOGIN_SUCCESS,
    REFRESH_JWT_SUCCESS,
    LOGOUT
} from '../ActionTypes';

const initialState = {
    token: "",
    name: "",
};


const userReducer = (state = initialState, action) => {
    switch (action.type) {
        case LOGIN_SUCCESS:
            return { ...state, token: action.payload.token, name: action.payload.name };
        case REFRESH_JWT_SUCCESS:
            return { ...state, token: action.payload.token };
        case LOGOUT:
            return initialState;
        default:
            return state;
    };
};

export default userReducer;