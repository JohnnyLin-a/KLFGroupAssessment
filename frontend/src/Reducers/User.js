import {
    LOGIN_SUCCESS
} from '../ActionTypes';

const initialState = {
    token: ""
};


const userReducer = (state = initialState, action) => {
    switch (action.type) {
        case LOGIN_SUCCESS:
            console.log("Reducer action.payload.token", action.payload.token)
            return { ...state, token: action.payload.token }
        default:
            return state;
    };
};

export default userReducer;