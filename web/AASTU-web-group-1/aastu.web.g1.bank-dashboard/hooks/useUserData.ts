import { useQuery, useQueryClient } from 'react-query';
import ky from 'ky';
import { getSession } from 'next-auth/react';
import { refreshAccessToken } from '../lib/auth';

interface UserData {
  id: string;
  name: string;
  email: string;
}

// async function fetchUserData() {
//   const session = await getSession();
//   if (!session?.user?.accessToken) throw new Error('No access token found');

//   const response = await ky.get('/user/data', {
//     headers: {
//       Authorization: `Bearer ${session.user.accessToken}`,
//     },
//   }).json<UserData>();

//   return response;
// }

// export function useUserData() {
//   const queryClient = useQueryClient();

//   return useQuery<UserData, Error>('userData', fetchUserData, {
//     onError: async (error) => {
//       if (error.message === 'No access token found') {
//         const session = await getSession();
//         if (session?.user?.refreshToken) {
//           const newAccessToken = await refreshAccessToken(session.user.refreshToken);
//           if (newAccessToken) {
//             session.user.accessToken = newAccessToken;

//             queryClient.invalidateQueries('userData'); 
//           } else {
//             throw new Error('Failed to refresh access token');
//           }
//         }
//       }
//     },
//   });
// }
