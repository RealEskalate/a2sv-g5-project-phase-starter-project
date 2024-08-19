'use client';
import { useSession } from 'next-auth/react';
import { useDispatch,useSelector } from 'react-redux';
import React, { useEffect } from 'react';

export default function Login() {
  const { data: session, status } = useSession();
  const x = useSelector((state) => state.user);

  
  const dispatch = useDispatch();

  useEffect(() => {
    if (status === 'authenticated' && session?.user?.accessToken) {
      dispatch({
        type: 'USER_FETCH_REQUESTED',
        payload: {
          userName: session.user.name,
          accessToken: session.user.accessToken,
        },
      });
    }
  }, [status, session, dispatch]);

  if (status === 'loading') {
    return <div>Loading...</div>;
  }

  if (status === 'unauthenticated') {
    return <div>Please log in</div>;
  }
  console.log(x,'x')
  return (
    <div>
      Welcome, {x.username}
    </div>
  );
}
