import {
    LOGIN
} from '../ActionTypes';

const initialState = {

};


const userReducer = (state = initialState, action) => {
    switch (action.type) {
        case LOGIN:
            return { ...state }
        default:
            return state;
    };
};

export default userReducer;