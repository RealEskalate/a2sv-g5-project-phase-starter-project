import ky from "ky";

interface Credentials {
  userName: string;
  password: string;
}

interface SignInResponse {
  success: boolean;
  message: string;
  data: {
    id: string;
    name: string;
    email: string;
    access_token: string;
    refresh_token: string;
  };
}

export async function signInWithCredentials(credentials: Credentials): Promise<SignInResponse> {
  return await ky.post("https://bank-dashboard-6acc.onrender.com/auth/login", {
    json: credentials,
  }).json<SignInResponse>();
}

export async function refreshAccessToken(refreshToken: string): Promise<string | null> {
    try {
      const response = await ky.post('https://bank-dashboard-6acc.onrender.com/refresh-token', {
        json: { refreshToken },
      }).json<{ access_token: string }>();
  
      return response.access_token;
    } catch (error) {
      console.error('Failed to refresh access token:', error);
      return null;
    }
  }
