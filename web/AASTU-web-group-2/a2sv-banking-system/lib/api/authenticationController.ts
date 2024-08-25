// authenticationController.ts
import { RegisterRequest, RegisterResponse, RefreshTokenResponse, LoginRequest, LoginResponse, ChangePasswordRequest, ChangePasswordResponse, RefreshTokenReturn } from '@/types/authenticationController.interface';

const BASE_URL = 'https://bank-dashboard-mih0.onrender.com'


const register = async (userDetails: RegisterRequest): Promise<RegisterResponse> => {
  try {
    const response = await fetch(`${BASE_URL}/auth/register`, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify(userDetails),
    });

    if (response.status === 200) {
      const data: RegisterResponse = await response.json();
      return data;
    } else {
      throw new Error(`Registration failed with status code: ${response.status}`);
    }
  } catch (error) {
    console.error('Error registering user:', error);
    throw error;
  }
};

const refreshToken = async (refresh_token:string): Promise<RefreshTokenReturn> => {
    try {
      const response = await fetch(`${BASE_URL}/auth/refresh_token`, {
        method: 'POST',
        headers: {
          'Authorization': `Bearer ${refresh_token}`, // Add the token to the headers
        },
      });
  
      if (response.status === 200) {
        const data: RefreshTokenResponse = await response.json();
        return data.data;
      } else {
        throw new Error(`Failed to refresh token. Status code: ${response.status}`);
      }
    } catch (error) {
      console.error('Error refreshing token:', error);
      throw error;
    }
  };

const login = async (credentials: LoginRequest): Promise<any> => {
    try {
      const response = await fetch(`${BASE_URL}/auth/login`, {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify(credentials),
      });
  
      if (response.status === 200) {
        const data: LoginResponse = await response.json();
        return data;
      } else {
        throw new Error(`Login failed with status code: ${response.status}`);
      }
    } catch (error) {
      console.error('Error logging in:', error);
      throw error;
    }
  };

  const changePassword = async (
    changePasswordDetails: ChangePasswordRequest,
    token: string
  ): Promise<ChangePasswordResponse> => {
    try {
      const response = await fetch(`${BASE_URL}/auth/change_password`, {
        method: 'POST',
        headers: {
          'Authorization': `Bearer ${token}`, // Add the token to the headers
          'Content-Type': 'application/json', // Missing Content-Type header
        },
        body: JSON.stringify(changePasswordDetails),
      });
  
      if (response.ok) {
        const data: ChangePasswordResponse = await response.json();
        return data;
      } else {
        const errorData = await response.json();
        throw new Error(
          `Change password failed with status code: ${response.status} - ${errorData.message || 'Unknown error'}`
        );
      }
    } catch (error) {
      console.error('Error changing password:', error);
      throw error;
    }
  };

// Named exports
export { register, refreshToken, login, changePassword };
