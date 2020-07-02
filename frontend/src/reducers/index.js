import { combineReducers } from 'redux';
import userReducer from './user';
/**
 * Root reducer, combined reducers
 */
const rootReducer = combineReducers({
    user: userReducer,
});

export default rootReducer;