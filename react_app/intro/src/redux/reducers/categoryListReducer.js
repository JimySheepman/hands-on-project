import * as actionsTypes from "../actions/actionTypes"
import initialState from "./initialState";

export default function categoryListReducer(state=initialState.categories,action){
    switch (action.type) {
        case actionsTypes.GET_CATEGORIES_SUCCESS:
            return action.payload
    
        default:
            return state;
    }
}