'use client';

import { Session } from 'next-auth';
import { signOut } from 'next-auth/react';
import React from 'react'
 interface ExtendedSession extends Session {
    error?: string;
  }
  export const  expiredRefreshToken = (session)=>{
  const extendedSession = session as ExtendedSession;

 if (extendedSession?.error === 'RefreshAccessTokenError') {
    signOut();
  }}
  function Error({ session }: { session: ExtendedSession }) {
    expiredRefreshToken(session)
    return (
      <div>
      </div>
    );
  }
  
  export default Error;
  
