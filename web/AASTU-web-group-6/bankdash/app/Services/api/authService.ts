import LoginValue from "@/types/LoginValue";
import UserValue from "@/types/UserValue";
import axios from "axios";

const API_URL = "https://bank-dashboard-6acc.onrender.com/auth";

// Request & Response Interfaces


interface RegisterResponse {
  success: boolean;
  message: string;
  data: {
    access_token: string;
    refresh_token: string;
    data: UserValue;
  };
}



interface LoginResponse {
  success: boolean;
  message: string;
  data: {
    access_token: string;
    refresh_token: string;
    data: any;
  };
}

interface RefreshTokenResponse {
  success: boolean;
  message: string;
  data: string;
}

interface ChangePasswordRequest {
  password: string;
  newPassword: string;
}

interface ChangePasswordResponse {
  success: boolean;
  message: string;
  data: any;
}

// Utility function to handle requests
const handleRequest = async <T>(
  method: string,
  endpoint: string,
  data?: any,
  accessToken?: string
): Promise<T> => {
  try {
    const response = await axios({
      method,
      url: endpoint,
      data,
      headers: {
        Authorization: accessToken ? `Bearer ${accessToken}` : undefined,
        "Content-Type": "application/json",
      },
    });
    return response.data as T;
  } catch (error) {
    if (axios.isAxiosError(error)) {
      console.error("Axios error:", error.message);
    } else {
      console.error("Unexpected error:", error);
    }
    throw error;
  }
};

// AuthService class
class AuthService {
  public static register(
    formData: UserValue
  ): Promise<RegisterResponse> {
    return handleRequest<RegisterResponse>(
      "POST",
      `${API_URL}/register`,
      formData
    );
  }

  public static login(
    loginData: LoginValue
  ): Promise<LoginResponse> {
    return handleRequest<LoginResponse>(
      "POST",
      `${API_URL}/login`,
      loginData
    );
  }

  public static refreshToken(
    accessToken: string
  ): Promise<RefreshTokenResponse> {
    return handleRequest<RefreshTokenResponse>(
      "POST",
      `${API_URL}/refresh_token`,
      undefined,
      accessToken
    );
  }

  public static changePassword(
    passwordData: ChangePasswordRequest,
    accessToken: string
  ): Promise<ChangePasswordResponse> {
    return handleRequest<ChangePasswordResponse>(
      "POST",
      `${API_URL}/change_password?password=${passwordData.password}&newPassword=${passwordData.newPassword}`,
      undefined,
      accessToken
    );
  }
}

export default AuthService;
