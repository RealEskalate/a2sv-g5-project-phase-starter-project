import { put, takeEvery, call, all } from 'redux-saga/effects';
import { setUser } from './slice/userSlice';

// const session = getSession(); 
const authorization = session ? session.get('authorization') : "eyJhbGciOiJIUzM4NCJ9.eyJzdWIiOiJtZWxrZSIsImlhdCI6MTcyMzgxNDM4NCwiZXhwIjoxNzIzOTAwNzg0fQ.I9wHa2c4pCdEDU--_U-PHklh_mZztXOzYJ5dREjlVi_NmfnJEgomvyWm1c5chWTK";
// Mock API call
function fetchUserDataApi() {
    const token = 'your_access_token_here';  // Replace with the actual token
  
    return fetch('https://bank-dashboard-6acc.onrender.com/user/melke', {
      method: 'GET',
      headers: {
        'Authorization': `Bearer ${token}`,
        'Content-Type': 'application/json'
      }
    })
      .then(response => {
        if (!response.ok) {
          throw new Error('Network response was not ok');
        }
        return response.json();
      })
      .catch(error => { throw error });
  }
  
function* fetchUserData(): Generator<any, void, any> {
  try {
    const data = yield call(fetchUserDataApi);
    yield put(setUser(data.data.data));
  } catch (error) {
    console.error('Failed to fetch user data:', error);
  }
}

export function* watchFetchUserData() {
  yield takeEvery('user/fetchUserData', fetchUserData);
}

export default function* rootSaga() {
  yield all([
    watchFetchUserData(),
  ]);
} 

