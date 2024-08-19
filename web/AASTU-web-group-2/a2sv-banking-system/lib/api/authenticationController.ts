// authenticationController.ts
import { RegisterRequest, RegisterResponse, RefreshTokenResponse, LoginRequest, LoginResponse, ChangePasswordRequest, ChangePasswordResponse } from '@/types/authenticationController.interface';
import { getServerSession } from 'next-auth';

const BASE_URL = 'https://bank-dashboard-6acc.onrender.com'

interface ExtendedUser {
  refresh_token: string;
  data: any; // Assuming `data` contains user information or other details
  accessToken?: string;
}

interface ExtendedSession {
  user?: ExtendedUser;
}

const fetchSession = async (): Promise<ExtendedSession> => {
  const session = await getServerSession();
  return session as ExtendedSession;
};

const getAccessToken = async (): Promise<string | undefined> => {
  const session = await fetchSession();
  return session?.user?.accessToken;
};

const getRefreshToken = async (): Promise<string | undefined> => {
  const session = await fetchSession();
  return session?.user?.refresh_token;
};

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
      const refresh_token = await getRefreshToken();
      const response = await fetch(`${BASE_URL}/auth/refresh_token`, {
        method: 'POST',
        headers: {
          'Authorization': `Bearer ${refresh_token}`, // Add the token to the headers
        },
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
    const token = await getAccessToken();
    const response = await fetch(`${BASE_URL}/auth/change_password`, {
      method: 'POST',
      headers: {
        'Authorization': `Bearer ${token}`, // Add the token to the headers

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
