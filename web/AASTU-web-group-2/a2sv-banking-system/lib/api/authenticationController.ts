// authenticationController.ts
import { RegisterRequest, RegisterResponse, RefreshTokenResponse, LoginRequest, LoginResponse, ChangePasswordRequest, ChangePasswordResponse } from '@/types/authenticationController.interface';

const BASE_URL = 'https://bank-dashboard-6acc.onrender.com'

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

const refreshToken = async (): Promise<RefreshTokenResponse> => {
    try {
      const response = await fetch(`${BASE_URL}/auth/refresh_token`, {
        method: 'POST',
      });
  
      if (response.status === 200) {
        const data: RefreshTokenResponse = await response.json();
        return data;
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

const changePassword = async (changePasswordDetails: ChangePasswordRequest): Promise<ChangePasswordResponse> => {
  try {
    const response = await fetch(`${BASE_URL}/auth/change_password`, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify(changePasswordDetails),
    });

    if (response.status === 200) {
      const data: ChangePasswordResponse = await response.json();
      return data;
    } else {
      throw new Error(`Change password failed with status code: ${response.status}`);
    }
  } catch (error) {
    console.error('Error changing password:', error);
    throw error;
  }
};

// Named exports
export { register, refreshToken, login, changePassword };
