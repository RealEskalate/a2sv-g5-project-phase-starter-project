import { fetchBaseQuery } from '@reduxjs/toolkit/query';
import { jwtDecode } from 'jwt-decode'; // Corrected import for jwt-decode
import { getSession, signOut } from 'next-auth/react';

export const baseQuery = (baseUrl = '/') => {
  return fetchBaseQuery({
    baseUrl: 'https://bank-dashboard-6acc.onrender.com',
    prepareHeaders: async (headers) => {
      const baseURL = 'https://bank-dashboard-6acc.onrender.com';

      const session = await getSession();
      // console.log('session from baseQuery', session);

      if (session && session.accessToken) {
        try {
          // Decode access token
          const decode: Record<string, any> = jwtDecode(session.accessToken);
          const currentTimestamp = Math.floor(Date.now() / 1000);
          // console.log('exp', decode, currentTimestamp);

          // Check if access token has expired
          if (decode.exp < currentTimestamp) {
            // console.log('Token has expired, refreshing token');

            // Attempt to refresh the token
            const refreshedToken = await fetch(`${baseURL}/auth/refresh_token`, {
              method: 'POST',
              headers: {
                'Content-Type': 'application/json',
                Authorization: `Bearer ${session.refreshToken}`,
              },
            }).then(async (res) => {
              // console.log('new token');
              if (res.ok) {
                const data = await res.json();
                // console.log('is here ', data);
                return data;
              }
              return null;
            });

            // console.log('refreshed token is', refreshedToken);
            if (refreshedToken) {
              headers.set('Authorization', `Bearer ${refreshedToken.data}`);
              // Update session with new tokens
              session.accessToken = refreshedToken.data;
            } else {
              // If refresh token is also expired or invalid, log out
              // console.log('Refresh token has expired, logging out');
              await signOut({ redirect: true, callbackUrl: '/login' }); // Redirect to login or another appropriate page
              return headers;
            }
          } else {
            // console.log('token is not expired');
            headers.set('Authorization', `Bearer ${session.accessToken}`);
          }
        } catch (error) {
          console.error('Error during token processing:', error);
          // Handle token decoding or API request errors
        }
      } else {
        // If no session or token, redirect to login
        // console.log('No session found, redirecting to login');
        await signOut({ redirect: true, callbackUrl: '/login' }); // Redirect to login or another appropriate page
      }

      headers.set('Content-Type', 'application/json');
      return headers;
    },
  });
};
