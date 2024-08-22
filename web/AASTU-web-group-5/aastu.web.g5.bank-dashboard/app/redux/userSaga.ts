// userSaga.js
import { call, put, takeLatest, select } from 'redux-saga/effects';
import axios from 'axios';
import { setUser } from './slice/userSlice';


function fetchUserData(username: string, token: string) {
  return axios.get(`https://bank-dashboard-1tst.onrender.com/user/${username}`, {
    headers: {
      Authorization: `Bearer ${token}`,
    },
  });
}

  function* fetchUser(action: { payload: { token: String; username: String; }; }) {
    try { 
      const response = yield call(fetchUserData, action.payload.userName, action.payload.accessToken);
      const { data } = response.data;
      yield put(setUser(data));
    } catch (e) {
      console.error("Failed to fetch user data", e);
    }
  }

function* userSaga() {
  yield takeLatest('USER_FETCH_REQUESTED', fetchUser);
}

export default userSaga;
